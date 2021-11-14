package retry

import (
	"context"
	"testing"
	"time"

	"github.com/nhannt315/real_estate_api/pkg/errors"
)

func TestExponentialBackoff(t *testing.T) {

	tt := newTestTimer()
	tn := &testNotify{}
	target := NewExponentialBackoff(
		WithMaxRetryCount(5),
		WithInterval(100*time.Millisecond),
		WithNotify(tn.Notify),
	).(*backoffRetryer)
	target.timer = tt

	tests := []struct {
		name     string
		newOp    func() Operation
		retryCnt int
		errStr   string
	}{
		{"no error", func() Operation {
			return func(ctx context.Context) (stopRetry bool, err error) {
				return false, nil
			}
		}, 0, ""},
		{"retry", func() Operation {
			cnt := 0
			return func(ctx context.Context) (stopRetry bool, err error) {
				cnt++
				if cnt <= 3 {
					return false, errors.New("err.hoge.1")
				}
				return false, nil
			}
		}, 3, ""},
		{"retry over", func() Operation {
			return func(ctx context.Context) (stopRetry bool, err error) {
				return false, errors.New("err.hoge.2")
			}
		}, 5, "err.hoge.2"},
		{"stop retry", func() Operation {
			cnt := 0
			return func(ctx context.Context) (stopRetry bool, err error) {
				cnt++
				if cnt == 4 {
					return true, errors.New("err.hoge.3")
				}
				return false, errors.New("err.hoge.3")
			}
		}, 3, "err.hoge.3"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tn.t = t
			tt.Clear()

			ctx := context.Background()
			err := target.Execute(ctx, test.newOp())
			if test.errStr == "" {
				if err != nil {
					t.Errorf("error. %+v", err)
				}
			} else if test.errStr != err.Error() {
				t.Errorf("invalid error. %+v", err)
			}

			if len(tt.intervals) != test.retryCnt {
				t.Errorf("invalid retry count. %d != %d",
					len(tt.intervals), test.retryCnt)
			}
		})
	}
}

func TestFixed(t *testing.T) {

	const initInterval = 100 * time.Millisecond
	tt := newTestTimer()
	tn := &testNotify{}
	target := NewFixed(
		WithMaxRetryCount(5),
		WithInterval(initInterval),
		WithNotify(tn.Notify),
	).(*backoffRetryer)
	target.timer = tt

	tests := []struct {
		name     string
		newOp    func() Operation
		retryCnt int
		errStr   string
	}{
		{"no error", func() Operation {
			return func(ctx context.Context) (stopRetry bool, err error) {
				return false, nil
			}
		}, 0, ""},
		{"retry", func() Operation {
			cnt := 0
			return func(ctx context.Context) (stopRetry bool, err error) {
				cnt++
				if cnt <= 3 {
					return false, errors.New("err.hoge.1")
				}
				return false, nil
			}
		}, 3, ""},
		{"retry over", func() Operation {
			return func(ctx context.Context) (stopRetry bool, err error) {
				return false, errors.New("err.hoge.2")
			}
		}, 5, "err.hoge.2"},
		{"stop retry", func() Operation {
			cnt := 0
			return func(ctx context.Context) (stopRetry bool, err error) {
				cnt++
				if cnt == 4 {
					return true, errors.New("err.hoge.3")
				}
				return false, errors.New("err.hoge.3")
			}
		}, 3, "err.hoge.3"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tn.t = t
			tt.Clear()

			ctx := context.Background()
			err := target.Execute(ctx, test.newOp())
			if test.errStr == "" {
				if err != nil {
					t.Errorf("error. %+v", err)
				}
			} else if test.errStr != err.Error() {
				t.Errorf("invalid error. %+v", err)
			}

			if len(tt.intervals) != test.retryCnt {
				t.Errorf("invalid retry count. %d != %d",
					len(tt.intervals), test.retryCnt)
			}
			// fixedのため interval が固定か確認
			for i, d := range tt.intervals {
				if d != initInterval {
					t.Errorf("invalid interval.[%d] %v",
						i, tt.intervals)
					break
				}
			}
		})
	}
}

type testTimer struct {
	c         chan time.Time
	intervals []time.Duration
}

func newTestTimer() *testTimer {
	return &testTimer{
		c: make(chan time.Time, 5),
	}
}

func (tt *testTimer) Clear() {
	tt.intervals = make([]time.Duration, 0)
}

func (tt *testTimer) Start(interval time.Duration) {
	tt.intervals = append(tt.intervals, interval)
	// テスト用タイマーのため即タイマー発生させる
	tt.c <- time.Now()
}

func (tt *testTimer) Stop() {
}

func (tt *testTimer) C() <-chan time.Time {
	return tt.c
}

type testNotify struct {
	t *testing.T
}

func (tn *testNotify) Notify(ctx context.Context, retryCount uint64, err error, interval time.Duration) {
	tn.t.Logf("[retry:%d][interval:%s]%v", retryCount, interval, err)
}
