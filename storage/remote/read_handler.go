// shouldUseClient returns true if the client should be used for the given query based on its matchers.
func shouldUseClient(client *RemoteReadClient, req *prompb.ReadRequest) bool {
	// If no matchers configured, always use the client
	if len(client.matchers) == 0 {
		return true
	}

	// Check each query in the request
	for _, query := range req.Queries {
		for _, configMatchers := range client.matchers {
			if matchesQuery(configMatchers, query.Matchers) {
				return true
			}
		}
	}
	return false
}

// matchesQuery checks if a query matches the required matchers.
func matchesQuery(requiredMatchers []*labels.Matcher, queryMatchers []*prompb.LabelMatcher) bool {
	// Convert prompb matchers to labels.Matcher for comparison
	var queryLabelMatchers []*labels.Matcher
	for _, m := range queryMatchers {
		var matchType labels.MatchType
		switch m.Type {
		case prompb.LabelMatcher_EQ:
			matchType = labels.MatchEqual
		case prompb.LabelMatcher_NEQ:
			matchType = labels.MatchNotEqual
		case prompb.LabelMatcher_RE:
			matchType = labels.MatchRegexp
		case prompb.LabelMatcher_NRE:
			matchType = labels.MatchNotRegexp
		default:
			// Unknown matcher type - skip this matcher
			continue
		}
		
		queryLabelMatchers = append(queryLabelMatchers, &labels.Matcher{
			Type:  matchType,
			Name:  m.Name,
			Value: m.Value,
		})
	}

	// For each required matcher, check if it's satisfied by the query matchers
	for _, required := range requiredMatchers {
		matched := false
		for _, query := range queryLabelMatchers {
			if required.Name == query.Name {
				// If names match, check if the values match according to the matching rules
				if required.Matches(query.Value) {
					matched = true
					break
				}
			}
		}
		if !matched {
			return false
		}
	}
	return true
}

// RemoteRead handles HTTP requests querying series from a remote server.
func (api *API) RemoteRead(r *http.Request) (apiFuncResult, error) {
	ctx := r.Context()
	if err := api.remoteReadEnabled(); err != nil {
		return apiFuncResult{}, err
	}

	req, err := decodeReadRequest(r)
	if err != nil {
		return apiFuncResult{}, err
	}

	externalLabels := api.config.ExternalLabels
	sortedExternalLabels := api.config.ExternalLabels.Map()

	// Filter clients based on matchers
	var matchingClients []*RemoteReadClient
	for _, client := range api.remoteReadClients {
		if shouldUseClient(client, req) {
			matchingClients = append(matchingClients, client)
		}
	}

	// If no clients match, use all clients as a fallback
	if len(matchingClients) == 0 {
		matchingClients = api.remoteReadClients
	}

	responseStreams := make(chan readResponseStream, len(matchingClients))
	defer close(responseStreams)

	g, ctx := errgroup.WithContext(ctx)

	for _, client := range matchingClients {
		// Execute remote read requests in parallel.
		client := client
		g.Go(func() error {
			resp, err := client.Read(ctx, req, externalLabels)
			if err != nil {
				return err
			}

			responseStreams <- readResponseStream{
				stream: resp,
				// We'll only get the remoteReadClientCustomName in tests.
				client: client.remoteReadClientCustomName,
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return apiFuncResult{}, err
	}

	hasData := false
	responses := []*prompb.ReadResponse{}

	for i := 0; i < len(matchingClients); i++ {
		stream := <-responseStreams
		clientResponses, warning, err := stream.stream.Responses()

		if warning != "" {
			level.Warn(api.logger).Log("msg", "Warning on remote read", "warning", warning)
		}

		if err != nil {
			return apiFuncResult{
				statusCode: http.StatusInternalServerError,
				err:        errors.Wrapf(err, "error receiving read response from remote read endpoint %s", stream.client),
			}, nil
		}

		for _, resp := range clientResponses {
			// Exit early if no timeseries.
			if len(resp.Results) == 0 {
				continue
			}

			hasData = true
			responses = append(responses, resp)
		}
	}

	if !hasData {
		return apiFuncResult{
			statusCode: http.StatusNoContent,
		}, nil
	}

	resp := prompb.ReadResponse{
		Results: make([]*prompb.QueryResult, 0, len(responses)),
	}
	for _, r := range responses {
		resp.Results = append(resp.Results, r.Results...)
	}

	data, err := resp.Marshal()
	if err != nil {
		return apiFuncResult{
			statusCode: http.StatusInternalServerError,
			err:        err,
		}, nil
	}

	return apiFuncResult{
		statusCode: http.StatusOK,
		response: HTTPResponse{
			data:              bytes.NewBuffer(data),
			contentTypeProto:  true,
			contentEncoding:   "snappy",
			externalLabels:    sortedExternalLabels,
			disableCompressed: api.enableAdmin,
		},
	}, nil
}
