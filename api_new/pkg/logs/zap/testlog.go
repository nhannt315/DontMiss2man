package zap

import (
	"bytes"
	"context"
	"regexp"
	"strings"
	"sync"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/nhannt315/real_estate_api/pkg/logs"
)

// TestLogger go testで使えるテスト用logger
//	内部では、とりあえず zap logger使う
type TestLogger struct {
	t            *testing.T
	log          logs.Logger
	outputBuffer *outputBuffer
}

func (l *TestLogger) Errorf(ctx context.Context, err error, format string, args ...interface{}) {
	panic("implement me")
}

func (l *TestLogger) IsDebugEnabled(ctx context.Context) bool {
	panic("implement me")
}

type outputBuffer struct {
	mutex   sync.Mutex
	buffer  *bytes.Buffer
	outputs []string
}

// NewTest creates new test logger that be used in go test
func NewTest(t *testing.T) *TestLogger {
	buffer := &bytes.Buffer{}
	log := newLogger(
		zapcore.AddSync(buffer),
		&logs.Config{
			Level:  logs.DebugLevel,
			Format: logs.JSONFormat,
		},
		nil,
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)

	l := &TestLogger{
		t:   t,
		log: log,
		outputBuffer: &outputBuffer{
			mutex:   sync.Mutex{},
			buffer:  buffer,
			outputs: make([]string, 0, 100),
		},
	}
	var _ logs.Logger = l // check logger implements
	return l
}

// AssertLast checks the contents of Testing and buffer
func (l *TestLogger) AssertLast(
	a *Asserter) {
	buf := l.outputBuffer
	buf.mutex.Lock()
	defer buf.mutex.Unlock()

	if len(buf.outputs) == 0 {
		l.t.Error("no log outputs.")
		return
	}
	a.Assert(l.t, buf.outputs[len(buf.outputs)-1])
}

// AssertMessages 指定されたメッセージが指定された順で出力されているか検証する。
//	メッセージは完全一致でのチェックではなく、部分一致で含まれていれば良い。
func (l *TestLogger) AssertMessages(
	msgs []string) {
	buf := l.outputBuffer
	buf.mutex.Lock()
	defer buf.mutex.Unlock()

	if len(msgs) == 0 {
		return // 不要
	}

	msgIdx := 0
	for _, output := range buf.outputs {
		msg := msgs[msgIdx]
		if !strings.Contains(output, msg) {
			continue // 見つからない
		}
		// 見つかったら次へ
		msgIdx++
		if msgIdx == len(msgs) {
			break // 全て見つかったら終了
		}
	}
	if msgIdx < len(msgs) {
		// 見つからないメッセージがあるためエラー
		l.t.Errorf("no output log. %s",
			msgs[msgIdx])
	}
}

// AssertMessagesStrict 指定されたメッセージが指定された順で出力されているか検証する。
//	メッセージは完全一致でのチェックではなく、部分一致で含まれていれば良い。
//  AssertMessagesと違い指定されていないメッセージが指定されたメッセージ間に含まれていた場合はエラー。
func (l *TestLogger) AssertMessagesStrict(
	msgs []string) {
	buf := l.outputBuffer
	buf.mutex.Lock()
	defer buf.mutex.Unlock()

	if len(msgs) == 0 {
		return // 不要
	}
	msgIdx := 0
	for _, output := range buf.outputs {
		msg := msgs[msgIdx]
		if !strings.Contains(output, msg) {
			// 指定されていないメッセージが含まれていたためエラー
			l.t.Errorf("unexpected output log. %s",
				output)
			return
		}
		// 見つかったら次へ
		msgIdx++
		if msgIdx == len(msgs) {
			break // 全て見つかったら終了
		}
	}
	if msgIdx < len(msgs) {
		// 見つからないメッセージがあるためエラー
		l.t.Errorf("no output log. %s",
			msgs[msgIdx])
	}
}

// AssertMessagePatterns 指定されたメッセージが指定された順で出力されているか検証する。
//	メッセージは、正規表現でマッチングされる。
func (l *TestLogger) AssertMessagePatterns(
	patterns []string) {
	buf := l.outputBuffer
	buf.mutex.Lock()
	defer buf.mutex.Unlock()

	if len(patterns) == 0 {
		return // 不要
	}
	exps := make([]*regexp.Regexp, 0, len(patterns))
	for _, p := range patterns {
		exps = append(exps, regexp.MustCompile(p))
	}

	msgIdx := 0
	for _, output := range buf.outputs {
		exp := exps[msgIdx]
		if !exp.MatchString(output) {
			continue // 見つからない
		}
		// 見つかったら次へ
		msgIdx++
		if msgIdx == len(exps) {
			break // 全て見つかったら終了
		}
	}
	if msgIdx < len(exps) {
		// 見つからないメッセージがあるためエラー
		l.t.Errorf("no output log. %s",
			patterns[msgIdx])
	}
}

// CountMessagePatterns 指定されたメッセージがそれぞれ何回出力されているか取得する。
//	メッセージは、正規表現でマッチングされる。
func (l *TestLogger) CountMessagePatterns(
	patterns []string) []int {
	buf := l.outputBuffer
	buf.mutex.Lock()
	defer buf.mutex.Unlock()

	if len(patterns) == 0 {
		return []int{} // 不要
	}
	exps := make([]*regexp.Regexp, 0, len(patterns))
	for _, p := range patterns {
		exps = append(exps, regexp.MustCompile(p))
	}

	counts := make([]int, len(patterns))
	for expIdx, exp := range exps {
		for _, output := range buf.outputs {
			if exp.MatchString(output) {
				counts[expIdx]++
			}
		}
	}
	return counts
}

// Asserter is struct that contains fields to verify the contents of the log in go test
type Asserter struct {
	Level        logs.Level
	Caller       string
	Message      string
	StackPattern string // regex pattern
	Fields       map[string]string
}

// Assert checks all fields with the exception of caller, msg, stack fields
// nolint: gocognit
func (a *Asserter) Assert(t *testing.T, out string) {

	kv := make(map[string]string)
	for k, v := range a.Fields {
		kv[k] = v
	}
	for _, k := range []string{"severity", "caller", "msg"} {
		kv[k] = ""
	}
	if len(a.StackPattern) > 0 {
		kv["stack"] = ""
	}
	for key, val := range kv {

		ptnStr := strings.ReplaceAll(`"KEYWORD":"([^"]*)"`, "KEYWORD", key)
		reg := regexp.MustCompile(ptnStr)
		matchs := reg.FindStringSubmatch(out)
		switch key {
		case "severity":
			if len(matchs) > 1 {
				lv := strings.ReplaceAll(strings.ToLower(a.Level.String()), "level", "")
				if lv != matchs[1] {
					t.Errorf("invalid log Level. %s != %s",
						lv, matchs[1])
				}
			} else {
				t.Errorf("no match. output log key. %s", key)
			}
		case "caller":
			if len(matchs) > 1 {
				if !strings.HasPrefix(matchs[1], a.Caller) {
					t.Errorf("invalid Caller. <%s> not in <%s>",
						a.Caller, matchs[1])
				}
			} else {
				t.Errorf("no match. output log key. %s", key)
			}
		case "msg":
			if len(matchs) > 1 {
				if a.Message != matchs[1] {
					t.Errorf("invalid message. %s != %s",
						a.Message, matchs[1])
				}
			} else {
				t.Errorf("no match. output log key. %s", key)
			}
		case "stack":
			if len(matchs) > 1 {
				r := regexp.MustCompile(a.StackPattern)
				if !r.MatchString(matchs[1]) {
					t.Errorf("invalid stack. <%s> no match <%s>",
						matchs[1], a.StackPattern)
				}
			} else {
				t.Errorf("no match. output log key. %s", key)
			}
		default:
			if len(matchs) > 1 {
				if val != matchs[1] {
					t.Errorf("invalid %s. %s != %s", key,
						val, matchs[1])
				}
			} else {
				t.Errorf("no match. output log key. %s", key)
			}
		}
	}
}

// ClearOutputs clears buffer
func (l *TestLogger) ClearOutputs() {
	buf := l.outputBuffer
	buf.mutex.Lock()
	defer buf.mutex.Unlock()

	buf.outputs = buf.outputs[0:0]
}

func (l *TestLogger) executeLogger(logFunc func()) {
	buf := l.outputBuffer
	buf.mutex.Lock()
	defer buf.mutex.Unlock()
	buf.buffer.Reset()

	logFunc()

	out := buf.buffer.String()
	out = strings.TrimRight(out, "\r\n") // 末尾の改行とる
	if len(out) > 0 {
		l.t.Log(out)
		buf.outputs = append(buf.outputs, out)
	}
}

// ErrorfWithStack prints error log with the specified msg format +stack field
func (l *TestLogger) Error(ctx context.Context, err error) {
	l.executeLogger(func() {
		l.log.Error(ctx, err)
	})
}

// Info prints info log
func (l *TestLogger) Info(ctx context.Context, msg string) {
	l.executeLogger(func() {
		l.log.Info(ctx, msg)
	})
}

// Infof prints info log with the specified msg format
func (l *TestLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	l.executeLogger(func() {
		l.log.Infof(ctx, format, args...)
	})
}

// Warn prints warn log
func (l *TestLogger) Warn(ctx context.Context, msg string) {
	l.executeLogger(func() {
		l.log.Warn(ctx, msg)
	})
}

// Warnf prints warn log with the specified msg format
func (l *TestLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.executeLogger(func() {
		l.log.Warnf(ctx, format, args...)
	})
}

// Debug prints debug log
func (l *TestLogger) Debug(ctx context.Context, msg string) {
	l.executeLogger(func() {
		l.log.Debug(ctx, msg)
	})
}

// Debugf prints debug log with the specified msg format
func (l *TestLogger) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.executeLogger(func() {
		l.log.Debugf(ctx, format, args...)
	})
}

func (l *TestLogger) AddFields(fields ...logs.Field) logs.Logger {
	return &TestLogger{
		t:            l.t,
		log:          l.log.AddFields(fields...),
		outputBuffer: l.outputBuffer,
	}
}

func (l *TestLogger) AddStack(err error) logs.Logger {
	return &TestLogger{
		t:            l.t,
		log:          l.log.AddStack(err),
		outputBuffer: l.outputBuffer,
	}
}

func (l *TestLogger) Printf(format string, args ...interface{}) {
	l.executeLogger(func() {
		l.log.Printf(format, args...)
	})
}

func (l *TestLogger) Wait() {
}

// NewChildLogger apply input options in logger, and return that
func (l *TestLogger) NewChildLogger(opts *logs.Options) logs.Logger {
	return &TestLogger{
		t:            l.t,
		log:          l.log.NewChildLogger(opts),
		outputBuffer: l.outputBuffer,
	}
}
