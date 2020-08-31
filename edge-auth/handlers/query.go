package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// MakeQueryHandler returns whether a client can access a resource
func MakeQueryHandler(config *Config, permittedPrefix []string, restrictedPrefix []string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		resource := query.Get("r")

		status := http.StatusOK
		if len(resource) == 0 {
			status = http.StatusBadRequest
		} else if isProtected(resource, restrictedPrefix) {
			status = http.StatusForbidden
		} else if !isProtected(resource, permittedPrefix) {
			// TODO: handle dashes in user names
			parts := strings.SplitN(strings.TrimPrefix(resource, "/function/"), "-", 2)
			customer := parts[0]
			function := parts[1]
			if function != "favicon.ico" {
				url := fmt.Sprintf("%s/balances/%s:spend", config.ReceiptVerifierURI, customer)
				resp, err := http.Post(url, "text/plain", bytes.NewBuffer([]byte(config.RequestPrice)))
				if err != nil || resp.StatusCode != 200 {
					status = http.StatusPaymentRequired
				}
			}
		}

		log.Printf("Validate %s => %d\n", resource, status)

		w.WriteHeader(status)

	}
}

func isProtected(resource string, protected []string) bool {
	for _, prefix := range protected {
		if strings.HasPrefix(resource, prefix) {
			return true
		}
	}
	return false
}
