package handlers

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

// MakeQueryHandler returns whether a client can access a resource
func MakeQueryHandler(config *Config, restrictedPrefix []string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		resource := query.Get("r")

		status := http.StatusOK
		if len(resource) == 0 {
			status = http.StatusBadRequest
		} else if hasPrefix(resource, restrictedPrefix) {
			status = http.StatusUnauthorized
		}

		log.Printf("Validate %s => %d\n", resource, status)

		if status == http.StatusTemporaryRedirect {
			var redirect *url.URL

			switch config.OAuthProvider {
			case gitlabName:
				redirect = buildGitLabURL(config)

				break
			case githubName:
				redirect = buildGitHubURL(config, "", config.Scope)

				break
			}

			http.Redirect(w, r, redirect.String(), http.StatusTemporaryRedirect)
			return
		}
		w.WriteHeader(status)

	}
}

func hasPrefix(resource string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(resource, prefix) {
			return true
		}
	}
	return false
}
