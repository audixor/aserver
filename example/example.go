//
// Copyright (c) 2024 Tenebris Technologies Inc.
// See LICENSE for further information.
//

// This is a simple example of how to implement an HTTP server using the easysrv package

package main

import (
	"github.com/tenebris-tech/easysrv"
	"net/http"
)

func main() {

	// Create a new EasySrv instance
	e := easysrv.New()

	// Set a log file
	e.LogFile = "example.log"

	// Enable the test handler
	e.TestHandler = true

	// Enable debug logging
	e.Debug = true

	// Add a handler for /help
	e.AddRoute(easysrv.Route{
		Name:    "help",
		Method:  "GET",
		Pattern: "/help",
		Handler: getHelp})

	// Start the server
	err := e.Start()
	if err != nil {
		panic(err)
	}
}

// getHelp returns a help message
// Handlers must accept *http.Request and return an easysrv.Response structure
// See easysrv/handlers.go for more examples including additional parameters
func getHelp(*http.Request) easysrv.Response {
	return easysrv.Response{Details: "This is a help message", Status: "ok", Code: http.StatusOK}
}
