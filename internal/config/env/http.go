package env

import (
	"errors"
	"os"

	"github.com/sarastee/gomobile-test-assignment/internal/config"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

// HTTPCfgSearcher searcher for HTTP config.
type HTTPCfgSearcher struct{}

// NewHTTPCfgSearcher get instance for HTTP config searcher.
func NewHTTPCfgSearcher() *HTTPCfgSearcher {
	return &HTTPCfgSearcher{}
}

// Get searcher for HTTP config.
func (s *HTTPCfgSearcher) Get() (*config.HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	return &config.HTTPConfig{
		Host: host,
		Port: port,
	}, nil
}
