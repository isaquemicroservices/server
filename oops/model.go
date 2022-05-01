package oops

const (
	grpcCode    = 1000
	defaultCode = 2000
)

// Error a type of error handling
type Error struct {
	Code       int      `json:"code"`
	Message    string   `json:"msg"`
	StatusCode int      `json:"-"`
	Trace      []string `json:"-"`
	Err        error    `json:"-"`
}

// Error return error message
func (e *Error) Error() string {
	return e.Message
}

// Unwrap return the specific error cause for this error
func (e *Error) Unwrap() error {
	return e.Err
}
