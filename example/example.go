//
// Copyright (c) 2024 Tenebris Technologies Inc.
// See LICENSE for further information.
//

// This is a simple example of how to implement an HTTP server using the easysrv package

package main

import (
	"net/http"

	"github.com/audixor/aserver"
)

func main() {

	// Create a new EasySrv instance
	server, err := easysrv.New(
		easysrv.WithLogFile("example.log"),
		easysrv.WithListen(":8080"),
		easysrv.WithTestHandler(true),
		easysrv.WithDebug(true),
	)
	if err != nil {
		panic(err)
	}

	// Add a handler for /help
	server.AddRoute(easysrv.Route{
		Name:    "help",
		Method:  "GET",
		Pattern: "/help",
		Handler: getHelp})

	// Start the server
	err = server.Start()
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
