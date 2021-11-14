package errors

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func Test_Print(t *testing.T) {
	tests := []struct {
		name string
		arg  error
		want string
	}{
		{
			name: "New",
			arg:  New("message only"),
			want: "message only",
		},

		{
			name: "Errorf",
			arg:  Errorf("message with %s, %d", "string arg", 1),
			want: "message with string arg, 1",
		},

		{
			name: "Wrap",
			arg:  Wrap(New("Error A"), "Error B"),
			want: "Error B: Error A",
		},

		{
			name: "Wrapf",
			arg:  Wrapf(New("Error A"), "Error B with %s, %d", "string arg", 1),
			want: "Error B with string arg, 1: Error A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := fmt.Sprintf("%v", tt.arg)
			//fmt.Printf("%+v \n", tt.arg)
			if s != tt.want {
				t.Errorf("Print = %v, want %v", s, tt.want)
			}
		})
	}
}

func Test_Wrap_And_Is(t *testing.T) {
	e := Errorf("error with args: %s", "str")
	tests := []struct {
		name string
		arg  error
		want error
	}{
		{
			name: "wrap",
			arg:  Wrap(e, "wrap"),
			want: e,
		},
		{
			name: "wrapf",
			arg:  Wrapf(e, "wrapf. args: %s, %d", "str", 1),
			want: e,
		},
		{
			name: "wrap and wrap",
			arg:  Wrap(Wrapf(e, "wrapf. args: %s, %d", "str", 1), "wrap"),
			want: e,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !Is(tt.arg, tt.want) {
				t.Errorf("not matched err: %v, target: %v", tt.arg, tt.want)
			}
		})
	}
}

func Test_Wrap_And_As(t *testing.T) {
	e := &testError{
		msg:  "test",
		code: 100,
	}

	tests := []struct {
		name string
		arg  error
		want error
	}{
		{
			name: "wrap",
			arg:  Wrap(e, "wrap"),
			want: e,
		},
		{
			name: "wrapf",
			arg:  Wrapf(e, "wrapf. args: %s, %d", "str", 1),
			want: e,
		},
		{
			name: "wrap and wrap",
			arg:  Wrap(Wrapf(e, "wrapf. args: %s, %d", "str", 1), "wrap"),
			want: e,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var unwrapped *testError

			ok := As(tt.arg, &unwrapped)
			if !ok {
				t.Errorf("not matched err: %v, target: %v", tt.arg, tt.want)
			}

			if !reflect.DeepEqual(tt.want, unwrapped) {
				t.Errorf("not matched err: %v, target: %v", unwrapped, tt.want)
			}
		})
	}

}

func Test_Wrap_And_Unwrap(t *testing.T) {
	e := &testError{
		msg:  "test",
		code: 100,
	}

	tests := []struct {
		name string
		arg  error
		want error
	}{
		{
			name: "wrap",
			arg:  Wrap(e, "wrap"),
			want: e,
		},
		{
			name: "wrapf",
			arg:  Wrapf(e, "wrapf. args: %s, %d", "str", 1),
			want: e,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			unwrapped := Unwrap(tt.arg)

			if !reflect.DeepEqual(tt.want, unwrapped) {
				t.Errorf("not matched err: %v, target: %v", unwrapped, tt.want)
			}
		})
	}

}

type testError struct {
	msg  string
	code int
}

func (e *testError) Error() string {
	return fmt.Sprintf("%s. code: %d", e.msg, e.code)
}

func Test_Stack(t *testing.T) {
	tests := []struct {
		name          string
		newError      func() error
		expectMessage string
		expectPattern string
	}{
		{"no wrap", func() error {
			return stackTestXxx("test hoge", false)
		}, "test hoge",
			`^test hoge.+stackTestZzz.+error_test.go.+stackTestYyy.+error_test.go.+stackTestXxx.+error_test.go.+Test_Stack.+error_test.go`},
		{"wrap", func() error {
			return stackTestXxx("test hoge", true)
		}, "xxx: yyy: test hoge",
			`^xxx: yyy: test hoge.+stackTestZzz.+error_test.go.+stackTestYyy.+error_test.go.+stackTestXxx.+error_test.go.+Test_Stack.+error_test.go` +
				`.+yyy.+stackTestYyy.+error_test.go.+stackTestXxx.+error_test.go.+Test_Stack.+error_test.go` +
				`.+xxx.+stackTestXxx.+error_test.go.+Test_Stack.+error_test.go`},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			err := test.newError()
			if test.expectMessage != err.Error() {
				t.Errorf("invalid error. <<%s>> != <<%s>>",
					test.expectMessage, err.Error())
			}
			if test.expectMessage != fmt.Sprintf("%s", err.Error()) {
				t.Errorf("invalid format v. <<%s>> != <<%s>>",
					test.expectMessage, fmt.Sprintf("%s", err.Error()))
			}
			if test.expectMessage != fmt.Sprintf("%v", err.Error()) {
				t.Errorf("invalid format v. <<%s>> != <<%s>>",
					test.expectMessage, fmt.Sprintf("%v", err.Error()))
			}

			s := fmt.Sprintf("%+v", err)
			t.Log(s)

			s = strings.ReplaceAll(s, "\n", "\\n")
			p := regexp.MustCompile(test.expectPattern)
			if !p.MatchString(s) {
				t.Errorf("unmatch stack -> %s", s)
			}
		})
	}
}

func stackTestXxx(msg string, wrap bool) error {
	if wrap {
		return Wrap(stackTestYyy(msg, wrap), "xxx")
	}
	return stackTestYyy(msg, wrap)
}

func stackTestYyy(msg string, wrap bool) error {
	if wrap {
		return Wrap(stackTestZzz(msg), "yyy")
	}
	return stackTestZzz(msg)
}

func stackTestZzz(msg string) error {
	return Errorf(msg)
}
