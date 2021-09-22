package errors

type Error struct {
	Msg string
}

// New  create new error
func New(msg string) *Error {
	return &Error{Msg: msg}
}

// Error
func (e *Error) Error() string {
	return e.Msg
}
