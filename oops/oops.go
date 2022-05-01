package oops

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

// fromError wraps errors to provide user readable messages
func fromError(e error) error {
	message, code, responseStatus := "Unknown error", 0, 400

	switch err := e.(type) {
	case error:
		if status, ok := status.FromError(err); ok {
			message, code, e = grpc(status)
		}
	case nil:
		return nil
	}

	return &Error{
		Code:       code,
		Message:    message,
		StatusCode: responseStatus,
		Err:        e,
	}
}

// Err create annotated error instance from any error value
func Err(err error) error {
	var e *Error
	if !errors.As(err, &e) {
		err = fromError(err)
	} else if err == e {
		err = fromError(err)
	}
	return errors.WithStack(err)
}

// Wrap an error adding an information message
func Wrap(err error, message string) error {
	return errors.Wrap(Err(err), message)
}

// NewErr creates an annotated errors instance with default values
func NewErr(message string) error {
	return Err(&Error{
		Code:       defaultCode,
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Err:        errors.Errorf("Inline error message: '%s'", message),
	})
}

// Handling handles an error by setting a message and a response status code
func Handling(err error, c *gin.Context) {
	var e *Error
	if !errors.As(err, &e) {
		Handling(Err(err), c)
		return
	}

	e.Message = err.Error()
	e.Trace, _ = RebuildStackTrace(err, e)

	c.JSON(e.StatusCode, e)
	c.Set("error", err)
	c.Abort()
}

// RebuildStackTrace reconstructs the stack trace
func RebuildStackTrace(err error, bound error) (output []string, traced bool) {
	var (
		wrapped wrappedError
		tracer  stackTracer
	)

	if errors.As(err, &wrapped) {
		var internal error = wrapped.Unwrap()
		if wrapped.Unwrap() != bound {
			output, traced = RebuildStackTrace(internal, bound)
		}
		if !traced && errors.As(err, &tracer) {
			stack := tracer.StackTrace()
			for _, frame := range stack {
				output = append(output, fmt.Sprintf("%+v", frame))
			}
			traced = true
		}
	}
	return
}
