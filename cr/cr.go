package cr

import (
	"log"
	"os"

	"fmt"
	"net/http"

	"cloud.google.com/go/compute/metadata"
)

var (
	logger = log.New(os.Stdout, "", 0)
)

// AuthorizeRequest appends header authorization using the default service account identity
func AuthorizeRequest(req *http.Request, audience string) error {

	tokenURL := fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", audience)
	idToken, err := metadata.Get(tokenURL)
	if err != nil {
		return fmt.Errorf("Error getting metadata for %s: %v", tokenURL, err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", idToken))

	return nil

}
