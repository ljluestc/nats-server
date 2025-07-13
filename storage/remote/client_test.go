func TestNewReadClient(t *testing.T) {
	cfg := &config.RemoteReadConfig{
		Name: "test-client",
		URL:  &config_util.URL{URL: &url.URL{Host: "localhost"}},
	}
	c, err := NewReadClient("test", cfg, nil)
	require.NoError(t, err)
	require.Equal(t, c.remoteName, "test")
	require.Equal(t, c.url.String(), "//localhost")
}

func TestNewReadClientWithMatchers(t *testing.T) {
	cfg := &config.RemoteReadConfig{
		Name: "test-client",
		URL:  &config_util.URL{URL: &url.URL{Host: "localhost"}},
		RequiredMatchers: []string{`{cluster="A"}`, `{job="prometheus"}`},
	}
	c, err := NewReadClient("test", cfg, nil)
	require.NoError(t, err)
	require.Equal(t, c.remoteName, "test")
	require.Equal(t, c.url.String(), "//localhost")
	require.Equal(t, len(c.matchers), 2)
	
	// Check the first matcher set - {cluster="A"}
	require.Equal(t, len(c.matchers[0]), 1)
	require.Equal(t, c.matchers[0][0].Name, "cluster")
	require.Equal(t, c.matchers[0][0].Value, "A")
	require.Equal(t, c.matchers[0][0].Type, labels.MatchEqual)
	
	// Check the second matcher set - {job="prometheus"}
	require.Equal(t, len(c.matchers[1]), 1)
	require.Equal(t, c.matchers[1][0].Name, "job")
	require.Equal(t, c.matchers[1][0].Value, "prometheus")
	require.Equal(t, c.matchers[1][0].Type, labels.MatchEqual)
}

func TestNewReadClientWithInvalidMatchers(t *testing.T) {
	cfg := &config.RemoteReadConfig{
		Name: "test-client",
		URL:  &config_util.URL{URL: &url.URL{Host: "localhost"}},
		RequiredMatchers: []string{`{invalid matcher}`},
	}
	_, err := NewReadClient("test", cfg, nil)
	require.Error(t, err)
}

func TestShouldUseClient(t *testing.T) {
	// Client with no matchers
	clientNoMatchers := &ReadClient{
		matchers: nil,
	}
	
	// Client with cluster="A" matcher
	clientWithMatcher, err := NewReadClient("test", &config.RemoteReadConfig{
		Name:             "test-client",
		URL:              &config_util.URL{URL: &url.URL{Host: "localhost"}},
		RequiredMatchers: []string{`{cluster="A"}`},
	}, nil)
	require.NoError(t, err)
	
	// Create a request with cluster="A" matcher
	reqWithMatchingLabel := &prompb.ReadRequest{
		Queries: []*prompb.Query{
			{
				Matchers: []*prompb.LabelMatcher{
					{
						Type:  prompb.LabelMatcher_EQ,
						Name:  "cluster",
						Value: "A",
					},
				},
			},
		},
	}
	
	// Create a request with cluster="B" matcher
	reqWithNonMatchingLabel := &prompb.ReadRequest{
		Queries: []*prompb.Query{
			{
				Matchers: []*prompb.LabelMatcher{
					{
						Type:  prompb.LabelMatcher_EQ,
						Name:  "cluster",
						Value: "B",
					},
				},
			},
		},
	}
	
	// Client with no matchers should accept all requests
	require.True(t, shouldUseClient(clientNoMatchers, reqWithMatchingLabel))
	require.True(t, shouldUseClient(clientNoMatchers, reqWithNonMatchingLabel))
	
	// Client with matcher should only accept matching requests
	require.True(t, shouldUseClient(clientWithMatcher, reqWithMatchingLabel))
	require.False(t, shouldUseClient(clientWithMatcher, reqWithNonMatchingLabel))
}
