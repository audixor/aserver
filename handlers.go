//
// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// See LICENSE for further information.
//

package easysrv

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

// HandlerHealth implements a health check for load balancers, etc.
//
//goland:noinspection GoUnusedParameter
func (e *EasySrv) HandlerHealth(r *http.Request) Response {
	var resp Response

	// Check for presence of the file that indicates the server is down
	if _, err := os.Stat(e.DownFile); err == nil {
		// file exists, send status down and 503
		resp.Status = "down"
		resp.Code = http.StatusServiceUnavailable
		resp.Details = "server is shutting down"
	} else {
		// does not exist - send ok and 200
		resp.Status = "ok"
		resp.Code = http.StatusOK
		resp.Details = "health check ok"
	}
	return resp
}

//goland:noinspection GoUnusedParameter
func (e *EasySrv) Handler401(r *http.Request) Response {
	return e.status4xx(http.StatusUnauthorized, "not authorized")
}

//goland:noinspection GoUnusedParameter
func (e *EasySrv) Handler404(r *http.Request) Response {
	return e.status4xx(http.StatusNotFound, "object does not exist")
}

//goland:noinspection GoUnusedParameter
func (e *EasySrv) Handler405(r *http.Request) Response {
	return e.status4xx(http.StatusMethodNotAllowed, "method not allowed")
}

// status4xx returns a 4xx error
func (e *EasySrv) status4xx(code int, message string) Response {
	return Response{Details: message, Status: "error", Code: code}
}

// HandlerTest accepts an optional 'id' variable and echos it back
// This is an example of a handler that can receive a variable in the URL or not
// Note that two routes are defined in routes.go, one with the variable and one without
func (e *EasySrv) HandlerTest(r *http.Request) Response {
	var resp Response

	// Get parameter
	vars := mux.Vars(r)
	id := vars["id"]

	// Create example response
	resp.Status = "ok"
	resp.Code = http.StatusOK

	if id == "" {
		resp.Details = "no ID received"
	} else {
		resp.Details = fmt.Sprintf("received ID %s", id)
	}

	// Send Response
	return resp
}
