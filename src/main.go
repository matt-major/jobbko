package main

import (
	"net/http"
	"time"

	"github.com/matt-major/jobbko/src/context"
)

func main() {
	appContext := context.CreateApplicationContext()

	r := NewRouter(appContext)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	orchestrator := ProcessorOrchestrator{
		numProcessors:  2,
		numGroups:      10,
		maxConcurrency: 2,
	}
	orchestrator.StartProcessors(appContext)

	if err := srv.ListenAndServe(); err != nil {
		appContext.Logger.Error(err)
	}
}
