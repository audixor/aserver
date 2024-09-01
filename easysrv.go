//
// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// See LICENSE for further information.
//

// Package easysrv implements a wrapper around gorilla/mux router and net/http server
// to create a production grade server and simply the creation of routes and handlers.
package easysrv

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/netutil"
)

type EasySrv struct {
	Headers          Headers
	Routes           Routes
	Listen           string
	HTTPTimeout      int
	HTTPIdleTimeout  int
	MaxConcurrent    int
	LogFile          string // Optional, defaults to stdout
	DownFile         string
	HealthHandler    bool
	TestHandler      bool
	StrictSlash      bool
	DefaultHeaders   bool
	TLS              bool
	TLSCertFile      string
	TLSKeyFile       string
	TLSStrongCiphers bool
	Debug            bool
	server           *http.Server
	Logger           Logger // Our logger interface for compatibility - see logger.go
}

type Handler func(request *http.Request) Response

type Route struct {
	Name    string
	Method  string
	Pattern string
	Handler Handler
}

type Routes []Route

type Header struct {
	Key   string
	Value string
}

type Headers []Header

// Response provides a consistent format for API responses
// Data is used to hold an appropriate structure
type Response struct {
	Status  string `json:"status"`            // Text Status
	Code    int    `json:"code"`              // HTTP status code
	Details string `json:"details,omitempty"` // optional response details
	Data    any    `json:"data,omitempty"`    // any is the new interface{}
}

// New returns a EasySrv struct with default values
//
//goland:noinspection GoUnusedExportedFunction
func New() *EasySrv {

	// Set defaults
	e := EasySrv{
		Listen:           "127.0.0.1:8080",
		HTTPTimeout:      60,
		HTTPIdleTimeout:  60,
		MaxConcurrent:    100,
		LogFile:          "", // Default to stdout
		DownFile:         "",
		HealthHandler:    true,
		TestHandler:      false,
		StrictSlash:      false,
		DefaultHeaders:   true,
		TLS:              false,
		TLSCertFile:      "",
		TLSKeyFile:       "",
		TLSStrongCiphers: true,
		Debug:            false,
	}
	return &e
}

// Start starts the API
func (e *EasySrv) Start() error {
	var err error

	// Set the logger to stdout if not already set
	if e.Logger == nil {
		e.Logger, err = NewSimpleLogger(e.LogFile)
		if err != nil {
			return err
		}
	}

	e.Logger.Info("Starting server", Fields{"listen": e.Listen}) // TODO

	// Add default headers if requested
	if e.DefaultHeaders {
		e.AddHeader("Cache-Control", "no-cache, no-store, must-revalidate")
		e.AddHeader("Pragma", "no-cache")
		e.AddHeader("Expires", "0")
	}

	// Add the health handler if requested
	if e.HealthHandler {
		e.AddRoute(Route{
			Name:    "health",
			Method:  "GET",
			Pattern: "/health",
			Handler: e.HandlerHealth,
		})
	}

	// Add the test handler if requested
	if e.TestHandler {
		e.AddRoutes(Routes{
			Route{
				Name:    "test",
				Method:  "GET",
				Pattern: "/test",
				Handler: e.HandlerTest,
			},
			Route{
				Name:    "test",
				Method:  "GET",
				Pattern: "/test/{id}",
				Handler: e.HandlerTest,
			},
		})
	}

	// Create a new gorilla/mux router
	router := mux.NewRouter().StrictSlash(e.StrictSlash)

	// Iterate through routes
	for _, route := range e.Routes {

		// Register the route using our wrapper
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(e.Wrapper(route.Name, route.Handler))
	}

	// Add catch all and not found handler
	router.PathPrefix("/").Handler(e.Wrapper("Handler404", e.Handler404))
	router.NotFoundHandler = e.Wrapper("Handler404", e.Handler404)
	router.MethodNotAllowedHandler = e.Wrapper("Handler404", e.Handler405)

	// Create server
	s := &http.Server{
		Addr:              e.Listen,
		Handler:           router,
		ReadHeaderTimeout: time.Duration(e.HTTPTimeout) * time.Second,
		ReadTimeout:       time.Duration(e.HTTPTimeout) * time.Second,
		WriteTimeout:      time.Duration(e.HTTPTimeout) * time.Second,
		IdleTimeout:       time.Duration(e.HTTPIdleTimeout) * time.Second,
	}

	// Add TLS configuration if option is enabled
	if e.TLS {
		if e.TLSCertFile == "" || e.TLSKeyFile == "" {
			return errors.New("TLS cert or key file not specified")
		}

		// Load the cert and key
		cert, err := tls.LoadX509KeyPair(e.TLSCertFile, e.TLSKeyFile)
		if err != nil {
			return err
		}

		// Create the TLS configuration
		tlsConfig := tls.Config{Certificates: []tls.Certificate{cert}}
		tlsConfig.MinVersion = tls.VersionTLS12

		if e.TLSStrongCiphers {
			tlsConfig.CipherSuites = []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
			}
		}

		// Add to the HTTP server config
		s.TLSConfig = &tlsConfig
	}

	// Start our customized server
	return e.listen(s)
}

func (e *EasySrv) Stop() error {

	// Tell the server it has 10 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Protect against nil server
	if e.server == nil {
		return errors.New("server is not running")
	}

	// Shutdown the server
	if err := e.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %s", err.Error())
	}

	// Shutdown was successful
	return nil
}

// AddRoutes adds routes to the router
func (e *EasySrv) AddRoutes(routes Routes) {
	// Iterate over routes and add to the router
	for _, route := range routes {
		e.AddRoute(route)
	}
}

// AddRoute adds a route to the router
func (e *EasySrv) AddRoute(route Route) {
	e.Routes = append(e.Routes, route)
}

// AddHeader adds a header to the list
func (e *EasySrv) AddHeader(key, value string) {
	e.Headers = append(e.Headers, Header{key, value})
}

// listen is a replacement for ListenAndServe that implements a concurrent session limit
// using netutil.LimitListener. If maxConcurrent is 0, no limit is imposed.
func (e *EasySrv) listen(srv *http.Server) error {

	// Store the server to allow for a graceful shutdown
	e.server = srv

	// Get listen address, default to ":http"
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}

	// Create listener
	rawListener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// If maxConcurrent > 0 wrap the listener with a limited listener
	var listener net.Listener
	if e.MaxConcurrent > 0 {
		listener = netutil.LimitListener(rawListener, e.MaxConcurrent)
	} else {
		listener = rawListener
	}

	// Start TLS or non-TLS listener
	if e.TLS {
		// This will use the previously configured TLS information
		return srv.ServeTLS(listener, "", "")
	} else {
		return srv.Serve(listener)
	}
}
