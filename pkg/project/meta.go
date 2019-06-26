package project

import (
	"log"
	"net/http"
	"os"
	"strings"

	meta "cloud.google.com/go/compute/metadata"
)

var (
	logger      = log.New(os.Stdout, "[ka] ", 0)
	agentName   = "gcputil"
	projectKeys = []string{
		"GCP_PROJECT",
		"GOOGLE_CLOUD_PROJECT",
		"GCLOUD_PROJECT",
	}
)

// GetProjectIDOrFail checks common projectID env vars first,
// then tries to retreave the project ID from meta data server
// Fails if neither are not successful
func GetProjectIDOrFail() string {
	p, err := deriveProjectID(agentName)
	if err != nil {
		logger.Fatalf("Error while getting project ID: %v", err)
	}
	return p
}

// GetProjectID checks common projectID env vars first,
// then tries to retreave the project ID from meta data server
func GetProjectID() (project string, err error) {
	return deriveProjectID(agentName)
}

func deriveProjectID(agent string) (p string, err error) {

	for _, key := range projectKeys {
		if val, ok := os.LookupEnv(key); ok {
			logger.Printf("Found %s: %s", key, val)
			return strings.TrimSpace(val), nil
		}
	}

	mc := meta.NewClient(&http.Client{
		Transport: userAgentTransport{
			userAgent: agent,
			base:      http.DefaultTransport,
		},
	})

	return mc.ProjectID()

}

type userAgentTransport struct {
	userAgent string
	base      http.RoundTripper
}

// RoundTrip implements the transport interface
// // https://godoc.org/cloud.google.com/go/compute/metadata#example-NewClient
func (t userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent)
	return t.base.RoundTrip(req)
}
