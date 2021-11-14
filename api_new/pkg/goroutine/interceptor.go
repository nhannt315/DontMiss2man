package goroutine

import (
	"context"
)

// Handler is invoked by Interceptor to complete the normal execution.
type Handler func(ctx context.Context, arg interface{})

// Interceptor provides a hook to intercept the execution on the goroutine.
// 	It is used to create goroutine, in case specified by RegisterInterceptor.
//	It is the responsibility of the interceptor to invoke handler to complete the goroutine.
type Interceptor func(ctx context.Context, handler Handler, arg interface{})

var interceptor Interceptor

func goWithInterceptor(ctx context.Context, arg interface{}, handler Handler) {
	if interceptor == nil {
		go handler(ctx, arg)
	} else {
		go interceptor(ctx, handler, arg)
	}
}

// RegisterInterceptor goroutine作成時使うinterceptorを指定。
func RegisterInterceptor(inputInterceptor Interceptor) {
	interceptor = inputInterceptor
}
