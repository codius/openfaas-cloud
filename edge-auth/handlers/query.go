package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// TODO: import me
type FunctionBalance struct {
	Balance     uint64 `json:"balance,string"`
	Invocations uint64 `json:"remainingInvocations,string"`
}

// MakeQueryHandler returns whether a client can access a resource
func MakeQueryHandler(config *Config, permitted []string, restrictedPrefix []string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		resource := query.Get("r")

		status := http.StatusOK
		if len(resource) == 0 {
			status = http.StatusBadRequest
		} else if hasPrefix(resource, restrictedPrefix) {
			status = http.StatusUnauthorized
		} else if !hasPrefix(resource, permitted) {
			gatewayURL := os.Getenv("gateway_url")
			function := strings.SplitN(strings.TrimPrefix(resource, "/function/"), "/", 2)[0]
			invocations, err := getRemainingInvocations(function, gatewayURL)
			if err != nil || invocations == 0 {
				if err != nil {
					log.Printf("Error getting remaining invocations for %s => %s", function, err)
				} else {
					log.Printf("No remaining invocations for %s", function)
				}
				status = http.StatusPaymentRequired
			}
		}

		log.Printf("Validate %s => %d\n", resource, status)

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

func getRemainingInvocations(function string, gatewayURL string) (uint64, error) {
	if resp, err := http.Get(gatewayURL + "/function/billing?function=" + function); err != nil {
		return 0, err
	} else if resp.Body == nil {
		return 0, fmt.Errorf("no invocations for function: %s", function)
	} else {
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}
		// fmt.Printf("%s", string(bodyBytes))
		balance := FunctionBalance{}
		if err := json.Unmarshal(bodyBytes, &balance); err != nil {
			return 0, err
		}
		return balance.Invocations, nil
	}
}
