package interceptors

import (
	"context"

	"github.com/nhannt315/real_estate_api/pkg/goroutine"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

// RecoveryInterceptor new interceptor of panic recovery.
//	recover if the interceptor receive a panic, and logged.
func RecoveryInterceptor(log logs.Logger) goroutine.Interceptor {
	return func(ctx context.Context, handler goroutine.Handler, arg interface{}) {
		defer func() {
			if r := recover(); r != nil {
				logs.LogPanic(ctx, log, r)
			}
		}()
		handler(ctx, arg)
	}
}
