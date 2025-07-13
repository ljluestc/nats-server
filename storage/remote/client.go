// ReadClient reads from a remote endpoint.
type ReadClient struct {
	remoteName     string
	url            *url.URL
	client         *http.Client
	timeout        time.Duration
	retryOnRateLimit bool
	readRecent     bool
	parserFunc     prompb.Parser
	headers        map[string]string
	matchers       [][]*labels.Matcher
}
