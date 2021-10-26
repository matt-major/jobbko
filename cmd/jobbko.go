package main

import (
	"net/http"
	"time"

	"github.com/matt-major/jobbko/pkg/context"
	"github.com/matt-major/jobbko/pkg/processor"
	"github.com/matt-major/jobbko/pkg/router"
)

func main() {
	appContext := context.CreateApplicationContext()

	r := router.NewRouter(appContext)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	orchestrator := processor.ProcessorOrchestrator{
		NumProcessors:  2,
		NumGroups:      10,
		MaxConcurrency: 2,
	}
	orchestrator.StartProcessors(appContext)

	if err := srv.ListenAndServe(); err != nil {
		appContext.Logger.Error(err)
	}
}
