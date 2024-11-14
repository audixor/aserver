//
// Copyright (c) 2024 Tenebris Technologies Inc.
// See LICENSE for further information.
//

// This is a simple example of how to implement an HTTP server using the aserver package

package main

import (
	"net/http"

	"github.com/audixor/aserver"
)

func main() {

	// Create a new AServer instance
	server, err := aserver.New(
		aserver.WithLogFile("example.log"),
		aserver.WithListen(":8080"),
		aserver.WithTestHandler(true),
		aserver.WithDebug(true),
	)
	if err != nil {
		panic(err)
	}

	// Add a handler for /help
	server.AddRoute(aserver.Route{
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
// Handlers must accept *http.Request and return an aserver.Response structure
// See aserver/handlers.go for more examples including additional parameters
func getHelp(*http.Request) aserver.Response {
	return aserver.Response{Details: "This is a help message", Status: "ok", Code: http.StatusOK}
}
