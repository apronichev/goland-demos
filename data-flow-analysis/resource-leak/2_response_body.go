package resource_leak

import (
	"log"
	"net/http"
)

func ping(client *http.Client, url string) error {
	resp, err := client.Get(url) // resource leak!
	if err != nil {
		return err
	}

	log.Println("response status:", resp.Status)
	// process the response...

	// the response body is not closed
	return nil
}
