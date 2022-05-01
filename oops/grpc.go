package oops

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func grpc(status *status.Status) (message string, code int, err error) {
	code = grpcCode + int(status.Code())
	err = fmt.Errorf(fmt.Sprintf("%v, %v", status.Message(), status.Details()))

	switch status.Code() {
	case codes.Canceled:
		message = "Service operation has been canceled."
	case codes.Unknown:
		message = "Unknown service error."
	case codes.InvalidArgument:
		message = "Service received invalid parameters."
	case codes.DeadlineExceeded:
		message = "Service took longer than expected, please try again."
	case codes.NotFound:
		message = "The service found no data."
	case codes.AlreadyExists:
		message = "Entity already exists in the service."
	case codes.PermissionDenied:
		message = "Permission denied to service."
	case codes.ResourceExhausted:
		message = "Service resource has been exhausted."
	case codes.FailedPrecondition:
		message = "Preconditions do not meet service requirements."
	case codes.Aborted:
		if details := status.Details(); len(details) > 0 {
			message, err = status.Message(), fmt.Errorf(fmt.Sprintf("%v", details))
		} else {
			message = "Service operation was aborted."
		}
	case codes.OutOfRange:
		message = "Service operation is out of range."
	case codes.Unimplemented:
		message = "Method not implemented in service."
	case codes.Internal:
		message = "An internal error occurred in the service."
	case codes.Unavailable:
		message = "The service is unavailable."
	case codes.DataLoss:
		message = "Service data was lost."
	case codes.Unauthenticated:
		message = "No service authentication."
	default:
		message = "Unknown service error."
	}

	return
}
