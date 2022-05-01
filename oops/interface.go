package oops

import "github.com/pkg/errors"

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type wrappedError interface {
	Unwrap() error
}
