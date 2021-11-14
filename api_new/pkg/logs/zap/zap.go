package zap

import (
	"context"
	"fmt"
	"io"
	"math"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/nhannt315/real_estate_api/pkg/logs"
	"github.com/nhannt315/real_estate_api/pkg/logs/fields"
	"github.com/nhannt315/real_estate_api/pkg/rollbar"
)

var lockedStdOutSyncer = zapcore.Lock(os.Stdout)

// SetLockedStdOutSyncer 外部パッケージのtest用のlockedStdOutSyncerに対するSetter
func SetLockedStdOutSyncer(w io.Writer) {
	lockedStdOutSyncer = zapcore.AddSync(w)
}

// Logger is wrapper of zap.Logger,
// that logs everything to zap.
type Logger struct {
	level   logs.Level
	logger  *zap.Logger
	rollbar *rollbar.Client
	fields  []logs.Field
}

// A WriteSyncer is an io.Writer that can also flush any buffered data. Note
// that *os.File (and thus, os.Stderr and os.Stdout) implement WriteSyncer.
type WriteSyncer interface {
	io.Writer
	Sync() error
}

// NewLogger returns new logger instance
// Stacktrace is shown above warn level
func NewLogger(conf *logs.Config, rollbar *rollbar.Client, fields ...logs.Field) logs.Logger {
	return newDefault(lockedStdOutSyncer, conf, rollbar, fields...)
}

func newDefault(out WriteSyncer, conf *logs.Config, rollbar *rollbar.Client, fields ...logs.Field) logs.Logger {
	var opts []zap.Option
	if len(fields) > 0 {
		opts = append(opts, zap.Fields(toZapFields(fields)...))
	}
	opts = append(opts, zap.AddCaller())
	opts = append(opts, zap.AddCallerSkip(1))

	return newLogger(out, conf, rollbar, opts...)
}

func newLogger(out WriteSyncer, conf *logs.Config, rollbar *rollbar.Client, options ...zap.Option) logs.Logger {

	encodeConf := zap.NewProductionEncoderConfig()
	encodeConf.EncodeTime = jstTimeEncoder
	encodeConf.EncodeDuration = durationEncoder
	encodeConf.LevelKey = logs.LevelFieldKey
	encodeConf.TimeKey = logs.TimeFieldKey
	encodeConf.StacktraceKey = logs.StackFieldKey

	var encoder zapcore.Encoder
	switch conf.Format {
	case logs.ConsoleFormat:
		encoder = zapcore.NewConsoleEncoder(encodeConf)
	case logs.JSONFormat:
		encoder = zapcore.NewJSONEncoder(encodeConf)
	case logs.UnknownFormat: // ここに来る前にConfigのrequiredチェックでerrorになるはず
		panic("invalid log format")
	}

	atomicLevel := zap.NewAtomicLevel()
	level := conf.Level
	switch level {
	case logs.ErrorLevel:
		atomicLevel.SetLevel(zap.ErrorLevel)
	case logs.WarnLevel:
		atomicLevel.SetLevel(zap.WarnLevel)
	case logs.InfoLevel:
		atomicLevel.SetLevel(zap.InfoLevel)
	case logs.DebugLevel, logs.UnknownLevel:
		atomicLevel.SetLevel(zap.DebugLevel)
	}
	zapLv := atomicLevel.Level()

	core := zapcore.NewCore(encoder, out,
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapLv
		}))

	zl := zap.New(core, options...)

	return &Logger{level: level, logger: zl, rollbar: rollbar}
}

// Error prints error log
// nolint: gocritic
func (sl *Logger) Error(ctx context.Context, err error) {
	flds := append(sl.fields, fields.Stack(err))
	sl.logger.Error(err.Error(), toZapFieldsAndDefaultValues(ctx, flds)...)
	if sl.rollbar != nil {
		sl.rollbar.ErrorWithExtras("Error", err, loggerFieldsToMap(flds))
	}
}

// Info prints info log
func (sl *Logger) Info(ctx context.Context, msg string) {
	sl.logger.Info(msg, toZapFieldsAndDefaultValues(ctx, sl.fields)...)
}

// Infof prints info log with the specified msg format
func (sl *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	sl.logger.Info(msg, toZapFieldsAndDefaultValues(ctx, sl.fields)...)
}

// Warn prints warn log
func (sl *Logger) Warn(ctx context.Context, msg string) {
	sl.logger.Warn(msg, toZapFieldsAndDefaultValues(ctx, sl.fields)...)
	if sl.rollbar != nil {
		sl.rollbar.MessageWithExtras("Warning", msg, loggerFieldsToMap(sl.fields))
	}
}

// Warnf prints warn log with the specified msg format
func (sl *Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	sl.logger.Warn(msg, toZapFieldsAndDefaultValues(ctx, sl.fields)...)
}

// IsDebugEnabled checks debug is possible
func (sl *Logger) IsDebugEnabled(ctx context.Context) bool {
	return sl.level.IsDebugEnabled()
}

// Debug prints debug log
func (sl *Logger) Debug(ctx context.Context, msg string) {
	sl.logger.Debug(msg, toZapFieldsAndDefaultValues(ctx, sl.fields)...)
}

// Debugf prints debug log with the specified msg format
func (sl *Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	sl.logger.Debug(msg, toZapFieldsAndDefaultValues(ctx, sl.fields)...)
}

// AddFields add filed to logger
func (sl *Logger) AddFields(fields ...logs.Field) logs.Logger {
	newLogger := *sl
	newLogger.fields = append(newLogger.fields, fields...)
	return &newLogger
}

func (sl *Logger) AddStack(err error) logs.Logger {
	return sl.AddFields(fields.Error(err))
}

// Printf prints log with the specified msg format
func (sl *Logger) Printf(format string, args ...interface{}) {
	sl.logger.Info(fmt.Sprintf(format, args...))
}

// Wait waits until flush all log
func (sl *Logger) Wait() {
	if sl.rollbar != nil {
		sl.rollbar.Wait()
	}
}

// NewChildLogger apply input options in logger, and return that
func (sl *Logger) NewChildLogger(opt *logs.Options) logs.Logger {
	childZapLogger := sl.logger
	if opt.CallerSkip != nil {
		childZapLogger = childZapLogger.WithOptions(zap.AddCallerSkip(*opt.CallerSkip))
	}
	if len(opt.Fields) != 0 {
		childZapLogger = childZapLogger.With(toZapFields(opt.Fields)...)
	}

	return &Logger{level: sl.level, logger: childZapLogger}
}

func toZapFieldsAndDefaultValues(ctx context.Context, flds []logs.Field) []zap.Field {
	zapFields := toZapFields(flds)

	if fld := fields.GoroutineID(ctx); fld != nil {
		zapFields = append(zapFields, loggerFieldToZapField(fld))
	}

	return zapFields
}

func toZapFields(flds []logs.Field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(flds))
	for _, f := range flds {
		zapFields = append(zapFields, loggerFieldToZapField(f))
	}

	return zapFields
}

func loggerFieldToZapField(fld logs.Field) zap.Field {
	switch v := fld.Value().(type) {
	case int64:
		return zap.Int64(fld.Key(), v)
	case string:
		return zap.String(fld.Key(), v)
	case uint64:
		return zap.Uint64(fld.Key(), v)
	case bool:
		return zap.Bool(fld.Key(), v)
	case time.Duration:
		return zap.Duration(fld.Key(), v)
	default:
		return zap.Field{Key: fld.Key(), Type: zapcore.UnknownType, Interface: fld.Value()}
	}
}

func loggerFieldsToMap(flds []logs.Field) map[string]interface{} {
	m := make(map[string]interface{})
	for _, f := range flds {
		m[f.Key()] = f.Value()
	}
	return m
}

// JSTTimeEncoder makes zap logger's time format
func jstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "2006-01-02T15:04:05.000+09:00"
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}

func durationEncoder(value time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	f := float64(value) / float64(time.Millisecond)

	switch {
	case f < 1: // 1msec より小さい場合は、micro secまで小数点表示
		enc.AppendFloat64(math.Round(f*1000) / 1000)
	default: // 1msec 以上は小数点以下切り捨て
		enc.AppendUint64(uint64(f))
	}
}
