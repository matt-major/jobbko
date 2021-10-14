package main

import (
	"log"
	"net/http"
	"time"

	aws "github.com/matt-major/jobbko/app/aws"
	"github.com/matt-major/jobbko/app/router"
)

func main() {
	r := router.New()

	aws.InitAwsSession()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
