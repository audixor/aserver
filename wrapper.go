//
// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// See LICENSE for further information.
//

package easysrv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/tenebris-tech/easysrv/SimpleLogger"
)

// Wrapper returns a standard http.HandlerFunc
// This wrapper provides consistent logging and HTTP headers
func (e *EasySrv) Wrapper(handlerName string, hFunc Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		// Get the start time and source IP
		startTime := time.Now()
		src := e.getIP(req)

		// Set headers
		for _, header := range e.Headers {
			w.Header().Set(header.Key, header.Value)
		}

		// Call the actual handler to service the request
		resp := hFunc(req)

		// Set reply headers
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "-1")

		// Send the response
		w.WriteHeader(resp.Code)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			e.Logger.Error(e.SEid+11, fmt.Sprintf("JSON encode error in %s handler: %s", handlerName, err.Error()),
				SimpleLogger.Fields{
					"src":     src,
					"method":  req.Method,
					"uri":     req.RequestURI,
					"handler": handlerName,
					"error":   err.Error(),
				})
		}

		// Get duration of request
		duration := time.Since(startTime)

		// Remove parameters from URI to avoid logging confidential information
		uri := strings.Split(req.RequestURI, "?")[0]

		// Log the event
		e.Logger.Info(e.SEid+10, fmt.Sprintf("%s %s %d", req.Method, uri, resp.Code),
			LoggerFields{
				"src":      src,
				"method":   req.Method,
				"uri":      uri,
				"code":     resp.Code,
				"handler":  handlerName,
				"duration": fmt.Sprintf("%.4f", duration.Seconds()),
			})
	})
}

// getIP returns an IP address by reading the forwarded-for
// header (for proxies or load balancers) and falls back to use the remote address.
func (e *EasySrv) getIP(r *http.Request) string {
	var s = ""
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		s = forwarded
	} else {
		s = r.RemoteAddr
	}

	// Clean up, remove port number
	if len(s) > 0 {
		if strings.HasPrefix(s, "[") {
			// IPv6 address
			t := strings.Split(s, "]")
			s = t[0][1:]
		} else {
			// IPv4 - hack off port number
			t := strings.Split(s, ":")
			s = t[0]
		}
	}
	return s
}
