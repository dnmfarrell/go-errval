// Package errval provides a generic error handler which reduces boilerplate
// and fragile error checking code.
package errval

// IErrVal is type which defines the interface of any ErrVal. The Catch method
// must be called in order to obtain the return value.
type IErrVal[A any] interface {
	Catch(func(error)) (bool, A)
}

// ErrVal is a struct type which can hold an error and a return value.
type ErrVal[A any] struct {
	err error
	val A
}

// Catch takes an error handler. If the ErrVal was built with an error, it
// passes it to the error handler and returns false and a zero value.
// Otherwise Catch returns true and the value the ErrVal was built with.
func (e *ErrVal[A]) Catch(f func(error)) (bool, A) {
	if e.err != nil {
		f(e.err)
		return false, e.val
	}
	return true, e.val
}

// Err is the error constructor for ErrVal. This should be used to return an
// error.
func Err[A any](e error) *ErrVal[A] {
	return &ErrVal[A]{err: e}
}

// Val is the value constructor for ErrVal. This should be used to return a
// value (when there is no error).
func Val[A any](v A) *ErrVal[A] {
	return &ErrVal[A]{val: v}
}
