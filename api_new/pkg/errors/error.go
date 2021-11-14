package errors

import (
	goerrors "errors"
	"fmt"
	"io"

	pkgerrors "github.com/pkg/errors"
)

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(message string) error {
	return &errWithStack{
		message: message,
		goErr:   goerrors.New(message),
		pkgErr:  pkgerrors.New(message),
	}
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	return &errWithStack{
		message: message,
		goErr:   goerrors.New(message),
		pkgErr:  pkgerrors.New(message),
	}
}

// Is reports whether any error in err's chain matches target.
//
// An error is considered to match a target if it is equal to that target or if
// it implements a method Is(error) bool such that Is(target) returns true.
func Is(err error, target error) bool {
	if e, ok := err.(*errWithStack); ok {
		if goerrors.Is(e.goErr, target) {
			return true
		}
	}
	return goerrors.Is(err, target)
}

// As finds the first error in err's chain that matches the type to which target
// points, and if so, sets the target to its value and returns true. An error
// matches a type if it is assignable to the target type, or if it has a method
// As(interface{}) bool such that As(target) returns true. As will panic if target
// is not a non-nil pointer to a type which implements error or is of interface type.
//
// The As method should set the target to its value and return true if err
// matches the type to which target points.
func As(err error, target interface{}) bool {
	if e, ok := err.(*errWithStack); ok {
		if goerrors.As(e.goErr, target) {
			return true
		}
	}
	return goerrors.As(err, target)
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	return &errWithStack{
		isWrapped: true,
		message:   message,
		goErr:     fmt.Errorf(message+": %w", err),
		pkgErr:    pkgerrors.Wrap(err, message),
	}
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is call, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	return &errWithStack{
		isWrapped: true,
		message:   message,
		goErr:     fmt.Errorf(message+": %w", err),
		pkgErr:    pkgerrors.Wrap(err, message),
	}
}

// Unwrap returns the result of calling the Unwrap method on err, if err implements
// Unwrap. Otherwise, Unwrap returns nil.
func Unwrap(err error) error {
	if e, ok := err.(*errWithStack); ok {
		return goerrors.Unwrap(e.goErr)
	}
	return goerrors.Unwrap(err)
}

type errWithStack struct {
	isWrapped     bool
	message       string // wrapしたerror内容含まない このerrorのメッセージ
	goErr, pkgErr error
}

func (es *errWithStack) Error() string {
	return es.goErr.Error()
}

func (es *errWithStack) Unwrap() error {
	return goerrors.Unwrap(es.goErr)
}

// nolint: errcheck
func (es *errWithStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			// stackの1行目には "xxx: yyy: zzz" のような wrapされたerror含むメッセージも出力
			if es.isWrapped {
				// このためwrapされている場合、自errorのメッセージを stackの前に出力
				fmt.Fprintf(s, "%s: %+v", es.message, es.pkgErr)
			} else {
				// wrapされていない最下層errorなら、そのままstack出力
				fmt.Fprintf(s, "%+v", es.pkgErr)
			}
			return
		}
		fmt.Fprintf(s, "%v", es.pkgErr)

	default:
		io.WriteString(s, es.goErr.Error())
	}
}
