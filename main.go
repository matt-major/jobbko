package main

import (
	"log"
	"net/http"
	"time"

	"github.com/matt-major/jobbko/awsc"
)

func main() {
	r := NewRouter()

	awsc.InitAwsSession()

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	orchestrator := ProcessorOrchestrator{
		numProcessors:  1,
		numGroups:      1,
		maxConcurrency: 10,
	}
	orchestrator.StartProcessors()

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
