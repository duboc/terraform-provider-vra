package vra

import (
	"testing"

	"github.com/go-openapi/runtime/client"
)

func TestClient(t *testing.T) {
	var tests = []struct {
		url  string
		host string
		path string
	}{
		{"http://www.example.com", "www.example.com", "/"},
		{"http://www.example.com/", "www.example.com", "/"},
		{"http://www.example.com/foo/bar", "www.example.com", "/foo/bar"},
		{"http://www.example.com/foo/bar/", "www.example.com", "/foo/bar/"},
	}

	for _, tt := range tests {
		apiClient, err := getAPIClient(tt.url, "", true)
		if err != nil {
			t.Errorf("getAPIClient returned error %s", err)
		}
		transport := apiClient.Transport.(*client.Runtime)
		if tt.host != transport.Host || tt.path != transport.BasePath {
			t.Errorf("getAPIClient expected host %s path %s, actual host %s, path %s", tt.host, tt.path, transport.Host, transport.BasePath)
		}
	}
}
