//
// Copyright (c) 2021-2024 Tenebris Technologies Inc.
// See LICENSE for further information.
//

package easysrv

// Functional options

func WithLogger(logger Logger) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.Logger = logger
		return nil
	}
}

func WithListen(listen string) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.Listen = listen
		return nil
	}
}

func WithHTTPTimeout(t int) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.HTTPTimeout = t
		return nil
	}
}

func WithHTTPIdleTimeout(t int) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.HTTPIdleTimeout = t
		return nil
	}
}

func WithMaxConcurrent(m int) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.MaxConcurrent = m
		return nil
	}
}

func WithLogFile(logfile string) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.LogFile = logfile
		return nil
	}
}

func WithDownFile(down string) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.DownFile = down
		return nil
	}
}

func WithSEid(seid uint32) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.SEid = seid
		return nil
	}
}

func WithHealthHandler(h bool) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.HealthHandler = h
		return nil
	}
}

func WithTestHandler(t bool) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.TestHandler = t
		return nil
	}
}

func WithStrictSlash(s bool) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.StrictSlash = s
		return nil
	}
}

func WithDefaultHeaders(d bool) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.DefaultHeaders = d
		return nil
	}
}

func WithTLS(t bool) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.TLS = t
		return nil
	}
}

func WithTLSCertFile(certfile string) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.TLSCertFile = certfile
		return nil
	}
}

func WithTLSKeyFile(keyfile string) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.TLSKeyFile = keyfile
		return nil
	}
}

func WithTLSStrongCiphers(c bool) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.TLSStrongCiphers = c
		return nil
	}
}

func WithDebug(d bool) func(*EasySrv) error {
	return func(e *EasySrv) error {
		e.Debug = d
		return nil
	}
}
