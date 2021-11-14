package retry

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
)

const (
	defaultMaxRetry       = uint64(30)
	defaultInterval       = 500 * time.Millisecond
	defaultMaxElapsedTime = 30 * time.Minute
)

// backoffRetryer 指数的バックオフ(Exponential Backoff)が可能な Retryer
type backoffRetryer struct {
	maxRetry   uint64
	interval   time.Duration
	notify     Notify
	newBackoff func() backoff.BackOff
	timer      backoff.Timer // テスト時ダミーtimer設定
}

// BackoffOption Retryerのオプション指定func
type BackoffOption func(r *backoffRetryer)

// WithMaxRetryCount 最大リトライ回数のオプション
func WithMaxRetryCount(cnt uint64) BackoffOption {
	return func(r *backoffRetryer) {
		r.maxRetry = cnt
	}
}

// WithInterval リトライ間隔のオプション
//	NewExponentialBackoff の場合、初回のリトライ間隔となる。
func WithInterval(interval time.Duration) BackoffOption {
	return func(r *backoffRetryer) {
		r.interval = interval
	}
}

// Notify リトライ通知用func
//	リトライ(sleep)前に呼び出される。
type Notify func(ctx context.Context, retryCount uint64, err error, interval time.Duration)

// WithNotify リトライ通知のオプション
func WithNotify(notify Notify) BackoffOption {
	return func(r *backoffRetryer) {
		r.notify = notify
	}
}

// NewExponentialBackoff 指数的バックオフ(Exponential Backoff)による Retryer
func NewExponentialBackoff(opts ...BackoffOption) Retryer {

	r := &backoffRetryer{
		maxRetry: defaultMaxRetry,
		interval: defaultInterval,
	}
	for _, o := range opts {
		o(r)
	}

	r.newBackoff = func() backoff.BackOff {
		bo := backoff.NewExponentialBackOff()
		bo.InitialInterval = r.interval
		bo.MaxElapsedTime = defaultMaxElapsedTime
		return bo
	}
	return r
}

// NewFixed 固定インターバルによる Retryer
func NewFixed(opts ...BackoffOption) Retryer {

	r := &backoffRetryer{
		maxRetry: defaultMaxRetry,
		interval: defaultInterval,
	}
	for _, o := range opts {
		o(r)
	}

	r.newBackoff = func() backoff.BackOff {
		return backoff.NewConstantBackOff(r.interval)
	}
	return r
}

// Execute 指定された処理を実行する。
//	処理エラーの場合、必要に応じてリトライされる。
func (r *backoffRetryer) Execute(ctx context.Context, op Operation) error {

	// リトライの準備
	bo := r.newBackoff()
	bo = backoff.WithMaxRetries(bo, r.maxRetry)
	bo = backoff.WithContext(bo, ctx)
	var notify backoff.Notify
	if r.notify != nil {
		retryCnt := uint64(0)
		notify = func(err error, interval time.Duration) {
			retryCnt++
			r.notify(ctx, retryCnt, err, interval)
		}
	}

	// 処理実行
	return backoff.RetryNotifyWithTimer(func() error {
		stopRetry, err := op(ctx)
		if stopRetry {
			// リトライしない
			return backoff.Permanent(err)
		}
		return err
	}, bo, notify, r.timer)
}
