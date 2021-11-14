package logs

import (
	"context"
)

// Logger is interface to log application info
type Logger interface {
	// Info prints info log
	Info(ctx context.Context, msg string)
	// Infof prints info log with the specified msg format
	Infof(ctx context.Context, format string, args ...interface{})

	// Error prints error log
	Error(ctx context.Context, err error)

	// Warn prints warn log
	Warn(ctx context.Context, msg string)
	// Warnf prints warn log with the specified msg format
	Warnf(ctx context.Context, format string, args ...interface{})

	// IsDebugEnabled checks debug is possible
	IsDebugEnabled(ctx context.Context) bool

	// Debug prints debug log
	Debug(ctx context.Context, msg string)
	// Debugf prints debug log with the specified msg format
	Debugf(ctx context.Context, format string, args ...interface{})

	// AddField add field for structure logging
	AddFields(fields ...Field) Logger
	// AAddStack equivalent to AddFields(fields.Stack(err))
	AddStack(err error) Logger

	// Printf prints log with the specified msg format
	Printf(format string, args ...interface{})

	// Wait waits until flush all log
	Wait()

	// NewChildLogger apply input options in logger, and return that
	NewChildLogger(opts *Options) Logger
}

// Options はNewChildLoggerでloggerを作成する時option指定するため使う。
type Options struct {
	// CallerSkip  作成するloggerへ追加する caller のスキップ数
	CallerSkip *int
	// Fields  作成するloggerへ追加するフィールド
	Fields []Field
}
