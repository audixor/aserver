# easysrv
Golang HTTPS server made easy

Copyright (c) 2024 Tenebris Technologies Inc.
Released under the MIT License. Please see the LICENSE file for additional information.

## Warning
This is alpha code and should not be used in production.

## Overview
This project exists because I hate writing and maintaining the same code more than once
and some quirks in the Go standard library don't meet realistic production requirements. 
Most notably, they don't provide a mechanism to limit the number of concurrent requests.
This could result in resource exhaustion and cause the application to crash. In many
applications it makes much more sense to limit the number of concurrent requests and

This package implements a production quality HTTP server with the following features:

- Limit the number of concurrent requests
- Log requests
- HTTPS (TLS) support
- Consistent HTTP headers added to every response
- Creating a file with a specific name will cause the health check to fail
- A simple logger that can be replaced with a custom logger if desired

Handlers are also simplified. Instead of using the `http.Handler` interface, handle
functions must accept *http.Request and return an easysrv.Response structure.

## Contributions
PRs are welcome provided that they are consistent with the MIT License.
