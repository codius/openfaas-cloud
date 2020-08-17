package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/openfaas/openfaas-cloud/edge-auth/handlers"
)

func main() {
	var receiptVerifierURI string
	var requestPrice string

	if val, exists := os.LookupEnv("receipt_verifier_uri"); exists {
		receiptVerifierURI = val
	}

	if val, exists := os.LookupEnv("request_price"); exists {
		requestPrice = val
	}

	config := &handlers.Config{
		ReceiptVerifierURI: receiptVerifierURI,
		RequestPrice:       requestPrice,
	}

	permittedPrefix := []string{
		"/function/system-dashboard",
		"/function/system-github-event",
	}

	// Functions which make up the pipeline and which should not
	// be exposed via public ingress.
	restrictedPrefix := []string{
		"/function/ofc-",
		"/function/github-push",
		"/function/git-tar",
		"/function/buildshiprun",
		"/function/garbage-collect",
		"/function/github-status",
		"/function/import-secrets",
		"/function/pipeline-log",
		"/function/list-functions",
		"/function/audit-event",
		"/function/echo",
		"/function/metrics",
		"/function/function-logs",

		//AWS
		"/function/register-image",

		// GitLab
		"/function/gitlab-status",
		"/function/gitlab-push",
	}

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK."))
	})

	router.HandleFunc("/q/", handlers.MakeQueryHandler(config, permittedPrefix, restrictedPrefix))
	// router.HandleFunc("/login/", handlers.MakeLoginHandler(config))
	// router.HandleFunc("/oauth2/", handlers.MakeOAuth2Handler(config))
	router.HandleFunc("/healthz/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK."))
	})

	timeout := time.Second * 10
	port := 8080
	if v, exists := os.LookupEnv("port"); exists {
		val, _ := strconv.Atoi(v)
		port = val
	}

	log.Printf("Using port: %d\n", port)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        router,
		ReadTimeout:    timeout,
		WriteTimeout:   timeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
