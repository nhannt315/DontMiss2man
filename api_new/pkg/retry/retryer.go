package retry

import "context"

// Operation Retryer によってリトライ可能な処理(func)
type Operation func(ctx context.Context) (stopRetry bool, err error)

// Retryer リトライ考慮して処理実行できる。
//	リトライ方法は、実装に依存する。
type Retryer interface {
	// Execute 指定された処理を実行する。
	//	処理エラーの場合、必要に応じてリトライされる。
	Execute(ctx context.Context, op Operation) error
}
