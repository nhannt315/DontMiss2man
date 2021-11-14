package interceptors

import (
	"context"

	"github.com/nhannt315/real_estate_api/pkg/goroutine"
)

// Chain new interceptor of multiple interceptors arranger.
//	this interceptor call nested multiple interceptors.
func Chain(interceptors ...goroutine.Interceptor) goroutine.Interceptor {
	n := len(interceptors)

	return func(ctx context.Context, handler goroutine.Handler, arg interface{}) {
		chainer := func(interceptor goroutine.Interceptor, handler goroutine.Handler) goroutine.Handler {
			return func(ctx context.Context, arg interface{}) {
				interceptor(ctx, handler, arg)
			}
		}

		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = chainer(interceptors[i], chainedHandler)
		}
		chainedHandler(ctx, arg)
	}
}
