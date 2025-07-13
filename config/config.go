
// RemoteReadConfig is the configuration for reading from remote storage.
type RemoteReadConfig struct {
	URL              *config_util.URL   `yaml:"url"`
	Name             string             `yaml:"name,omitempty"`
	RemoteTimeout    model.Duration     `yaml:"remote_timeout,omitempty"`
	Headers          map[string]string  `yaml:"headers,omitempty"`
	ReadRecent       bool               `yaml:"read_recent,omitempty"`
	RequiredMatchers []string           `yaml:"required_matchers,omitempty"`
	OAuth2           *config_util.OAuth2 `yaml:"oauth2,omitempty"`

	// We cannot do proper Go type embedding below as the parser will then parse
	// values arbitrarily into the overflow maps of further-down types.
	HTTPClientConfig config_util.HTTPClientConfig `yaml:",inline"`
}
