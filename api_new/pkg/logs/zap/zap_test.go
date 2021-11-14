package zap

import (
	"bytes"
	"context"
	"regexp"
	"testing"

	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/logs"

	"go.uber.org/zap/zapcore"
)

func TestLogger_Error(t *testing.T) {
	outbuf := &bytes.Buffer{}
	lg := newDefault(
		zapcore.AddSync(outbuf), &logs.Config{
			Level:  logs.ErrorLevel,
			Format: logs.JSONFormat,
		},
		nil)
	type args struct {
		ctx context.Context
		msg string
	}
	tests := []struct {
		name   string
		args   args
		regexp string
	}{
		{
			name: "Error 1: Normal",
			args: args{
				ctx: context.Background(),
				msg: "Error message",
			},
			regexp: `{"severity":"error","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Error message","stack":".+TestLogger_Error.+zap_test.go.+"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lg.Error(tt.args.ctx, errors.New(tt.args.msg))
			checkRegexpMatching(outbuf, tt.regexp, t)
			outbuf.Reset()
		})
	}
}

func TestLogger_Warn(t *testing.T) {
	outbuf := &bytes.Buffer{}
	lg := newDefault(zapcore.AddSync(outbuf), &logs.Config{
		Level:  logs.WarnLevel,
		Format: logs.JSONFormat,
	},
		nil)
	type args struct {
		ctx context.Context
		msg string
	}
	tests := []struct {
		name   string
		args   args
		regexp string
	}{
		{
			name: "Warn 1: Normal",
			args: args{
				ctx: context.Background(),
				msg: "Warn message",
			},
			regexp: `{"severity":"warn","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Warn message"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lg.Warn(tt.args.ctx, tt.args.msg)
			checkRegexpMatching(outbuf, tt.regexp, t)
			outbuf.Reset()
		})
	}
}

func TestLogger_Warnf(t *testing.T) {
	outbuf := &bytes.Buffer{}
	lg := newDefault(
		zapcore.AddSync(outbuf), &logs.Config{
			Level:  logs.WarnLevel,
			Format: logs.JSONFormat,
		},
		nil)
	type args struct {
		ctx  context.Context
		fmt  string
		args []interface{}
	}
	tests := []struct {
		name   string
		args   args
		regexp string
	}{
		{
			name: "Warnf 1: 0 Arguments",
			args: args{
				ctx:  context.Background(),
				fmt:  "Warn format",
				args: []interface{}{},
			},
			regexp: `{"severity":"warn","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Warn format"}`,
		},
		{
			name: "Warnf 2: 1 Arguments",
			args: args{
				ctx:  context.Background(),
				fmt:  "Warn format %s",
				args: []interface{}{"args1"},
			},
			regexp: `{"severity":"warn","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Warn format args1"}`,
		},
		{
			name: "Warnf 3: 2 Arguments",
			args: args{
				ctx:  context.Background(),
				fmt:  "Warn format %s %s",
				args: []interface{}{"args1", "args2"},
			},
			regexp: `{"severity":"warn","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Warn format args1 args2"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lg.Warnf(tt.args.ctx, tt.args.fmt, tt.args.args...)
			checkRegexpMatching(outbuf, tt.regexp, t)
			outbuf.Reset()
		})
	}
}

func TestLogger_Info(t *testing.T) {
	outbuf := &bytes.Buffer{}
	lg := newDefault(zapcore.AddSync(outbuf), &logs.Config{
		Level:  logs.InfoLevel,
		Format: logs.JSONFormat,
	},
		nil)
	type args struct {
		ctx context.Context
		msg string
	}
	tests := []struct {
		name   string
		args   args
		regexp string
	}{
		{
			name: "Info 1: Normal",
			args: args{
				ctx: context.Background(),
				msg: "Info message",
			},
			regexp: `{"severity":"info","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Info message"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lg.Info(tt.args.ctx, tt.args.msg)
			checkRegexpMatching(outbuf, tt.regexp, t)
			outbuf.Reset()
		})
	}
}

func TestLogger_Infof(t *testing.T) {
	outbuf := &bytes.Buffer{}
	lg := newDefault(zapcore.AddSync(outbuf), &logs.Config{
		Level:  logs.InfoLevel,
		Format: logs.JSONFormat,
	},
		nil)
	type args struct {
		ctx  context.Context
		fmt  string
		args []interface{}
	}
	tests := []struct {
		name   string
		args   args
		regexp string
	}{
		{
			name: "Infof 1: 0 Arguments",
			args: args{
				ctx:  context.Background(),
				fmt:  "Info format",
				args: []interface{}{},
			},
			regexp: `{"severity":"info","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Info format"}`,
		},
		{
			name: "Infof 2: 1 Arguments",
			args: args{
				ctx:  context.Background(),
				fmt:  "Info format %s",
				args: []interface{}{"args1"},
			},
			regexp: `{"severity":"info","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Info format args1"}`,
		},
		{
			name: "Infof 3: 2 Arguments",
			args: args{
				ctx:  context.Background(),
				fmt:  "Info format %s %s",
				args: []interface{}{"args1", "args2"},
			},
			regexp: `{"severity":"info","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Info format args1 args2"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lg.Infof(tt.args.ctx, tt.args.fmt, tt.args.args...)
			checkRegexpMatching(outbuf, tt.regexp, t)
			outbuf.Reset()
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	outbuf := &bytes.Buffer{}
	lg := newDefault(zapcore.AddSync(outbuf), &logs.Config{
		Level:  logs.DebugLevel,
		Format: logs.JSONFormat,
	},
		nil)
	type args struct {
		ctx context.Context
		msg string
	}
	tests := []struct {
		name   string
		args   args
		regexp string
	}{
		{
			name: "Debug 1: Normal",
			args: args{
				ctx: context.Background(),
				msg: "Debug message",
			},
			regexp: `{"severity":"debug","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Debug message"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lg.Debug(tt.args.ctx, tt.args.msg)
			checkRegexpMatching(outbuf, tt.regexp, t)
			outbuf.Reset()
		})
	}
}

func TestLogger_Debugf(t *testing.T) {
	outbuf := &bytes.Buffer{}
	lg := newDefault(zapcore.AddSync(outbuf), &logs.Config{
		Level:  logs.DebugLevel,
		Format: logs.JSONFormat,
	},
		nil)
	type args struct {
		ctx  context.Context
		fmt  string
		args []interface{}
	}
	tests := []struct {
		name   string
		args   args
		regexp string
	}{
		{
			name: "Debugf 1: 0 Arguments",
			args: args{
				ctx:  context.Background(),
				fmt:  "Debug format",
				args: []interface{}{},
			},
			regexp: `{"severity":"debug","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Debug format"}`,
		},
		{
			name: "Debugf 2: 1 Arguments",
			args: args{
				ctx:  context.Background(),
				fmt:  "Debug format %s",
				args: []interface{}{"args1"},
			},
			regexp: `{"severity":"debug","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Debug format args1"}`,
		},
		{
			name: "Debugf 3: 2 Arguments",
			args: args{
				ctx:  context.Background(),
				fmt:  "Debug format %s %s",
				args: []interface{}{"args1", "args2"},
			},
			regexp: `{"severity":"debug","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Debug format args1 args2"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lg.Debugf(tt.args.ctx, tt.args.fmt, tt.args.args...)
			checkRegexpMatching(outbuf, tt.regexp, t)
			outbuf.Reset()
		})
	}
}

func checkRegexpMatching(logBuffer *bytes.Buffer, regeString string, t *testing.T) {
	re := regexp.MustCompile(regeString)
	if !(re.MatchString(logBuffer.String())) {
		t.Errorf("invalid log. : %s", logBuffer.String())
	}
}

func TestLogger_AddField(t *testing.T) {
	type args struct {
		fld logs.Field
		ctx context.Context
		msg string
	}
	tests := []struct {
		name   string
		args   args
		regexp string
	}{
		{
			name: "add string fields",
			args: args{
				fld: logs.NewStringField("key", "value"),
				ctx: context.Background(),
				msg: "Info message",
			},
			regexp: `{"severity":"info","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Info message","key":"value"}`,
		},
		{
			name: "add string fields",
			args: args{
				fld: logs.NewInt64Field("key", 123),
				ctx: context.Background(),
				msg: "Info message",
			},
			regexp: `{"severity":"info","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Info message","key":123}`,
		},
	}
	outbuf := &bytes.Buffer{}
	sl := newDefault(zapcore.AddSync(outbuf), &logs.Config{
		Level:  logs.InfoLevel,
		Format: logs.JSONFormat,
	},
		nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taggedLogger := sl.AddFields(tt.args.fld)
			taggedLogger.Info(tt.args.ctx, tt.args.msg)
			checkRegexpMatching(outbuf, tt.regexp, t)
			outbuf.Reset()
		})
	}
}

func TestLogger_Multiple_AddField(t *testing.T) {
	type args struct {
		flds []logs.Field
		ctx  context.Context
		msg  string
	}
	tests := []struct {
		name   string
		args   args
		regexp string
	}{
		{
			name: "add same fields with different value",
			args: args{
				flds: []logs.Field{
					logs.NewStringField("key", "value1"),
					logs.NewStringField("key", "value2"),
				},
				ctx: context.Background(),
				msg: "Info message",
			},
			regexp: `{"severity":"info","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Info message","key":"value1","key":"value2"}`,
		},
		{
			name: "add same fields with same value",
			args: args{
				flds: []logs.Field{
					logs.NewStringField("key", "value"),
					logs.NewStringField("key", "value"),
				},
				ctx: context.Background(),
				msg: "Info message",
			},
			regexp: `{"severity":"info","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Info message","key":"value","key":"value"}`,
		},
		{
			name: "add same fields with same value",
			args: args{
				flds: []logs.Field{
					logs.NewStringField("key1", "value"),
					logs.NewStringField("key2", "value"),
				},
				ctx: context.Background(),
				msg: "Info message",
			},
			regexp: `{"severity":"info","timestamp":.+","caller":"zap/zap_test.go:\d+","msg":"Info message","key1":"value","key2":"value"}`,
		},
	}
	outbuf := &bytes.Buffer{}
	for _, tt := range tests {
		sl := newDefault(zapcore.AddSync(outbuf), &logs.Config{
			Level:  logs.InfoLevel,
			Format: logs.JSONFormat,
		},
			nil)
		t.Run(tt.name, func(t *testing.T) {
			for _, f := range tt.args.flds {
				sl = sl.AddFields(f)
			}
			sl.Info(tt.args.ctx, tt.args.msg)
			checkRegexpMatching(outbuf, tt.regexp, t)
			outbuf.Reset()
		})
	}
}
