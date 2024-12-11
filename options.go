//
// Copyright (c) 2024 Tenebris Technologies Inc.
// Please see the LICENSE file for details
//

package aserver

// Functional options

func WithLogger(logger Logger) func(*AServer) error {
	return func(e *AServer) error {
		e.Logger = logger
		return nil
	}
}

func WithListen(listen string) func(*AServer) error {
	return func(e *AServer) error {
		e.Listen = listen
		return nil
	}
}

func WithHTTPTimeout(t int) func(*AServer) error {
	return func(e *AServer) error {
		e.HTTPTimeout = t
		return nil
	}
}

func WithHTTPIdleTimeout(t int) func(*AServer) error {
	return func(e *AServer) error {
		e.HTTPIdleTimeout = t
		return nil
	}
}

func WithMaxConcurrent(m int) func(*AServer) error {
	return func(e *AServer) error {
		e.MaxConcurrent = m
		return nil
	}
}

func WithLogFile(logfile string) func(*AServer) error {
	return func(e *AServer) error {
		e.LogFile = logfile
		return nil
	}
}

func WithDownFile(down string) func(*AServer) error {
	return func(e *AServer) error {
		e.DownFile = down
		return nil
	}
}

func WithSEid(seid uint32) func(*AServer) error {
	return func(e *AServer) error {
		e.SEid = seid
		return nil
	}
}

func WithHealthHandler(h bool) func(*AServer) error {
	return func(e *AServer) error {
		e.HealthHandler = h
		return nil
	}
}

func WithTestHandler(t bool) func(*AServer) error {
	return func(e *AServer) error {
		e.TestHandler = t
		return nil
	}
}

func WithStrictSlash(s bool) func(*AServer) error {
	return func(e *AServer) error {
		e.StrictSlash = s
		return nil
	}
}

func WithDefaultHeaders(d bool) func(*AServer) error {
	return func(e *AServer) error {
		e.DefaultHeaders = d
		return nil
	}
}

func WithTLS(t bool) func(*AServer) error {
	return func(e *AServer) error {
		e.TLS = t
		return nil
	}
}

func WithTLSCertFile(certfile string) func(*AServer) error {
	return func(e *AServer) error {
		e.TLSCertFile = certfile
		return nil
	}
}

func WithTLSKeyFile(keyfile string) func(*AServer) error {
	return func(e *AServer) error {
		e.TLSKeyFile = keyfile
		return nil
	}
}

func WithTLSStrongCiphers(c bool) func(*AServer) error {
	return func(e *AServer) error {
		e.TLSStrongCiphers = c
		return nil
	}
}

func WithDebug(d bool) func(*AServer) error {
	return func(e *AServer) error {
		e.Debug = d
		return nil
	}
}
