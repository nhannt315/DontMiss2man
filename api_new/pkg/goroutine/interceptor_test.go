package goroutine_test

import (
	"context"
	"testing"

	"github.com/nhannt315/real_estate_api/pkg/logs"
	"github.com/nhannt315/real_estate_api/pkg/logs/zap"

	"github.com/nhannt315/real_estate_api/pkg/goroutine"
)

func Test_goWithInterceptor(t *testing.T) {
	done := make(chan bool)
	tests := []struct {
		name string

		newInterceptor func(log logs.Logger) goroutine.Interceptor
		outputLogs     []string
	}{
		{
			"nil interceptor", func(log logs.Logger) goroutine.Interceptor {
				return nil
			}, []string{
				`.+"handler is executed: Argument_sample".+"goroutine_id":\d+.+`,
			},
		},
		{
			"not nil interceptor", func(log logs.Logger) goroutine.Interceptor {
				return func(ctx context.Context, handler goroutine.Handler, arg interface{}) {
					log.Info(ctx, "Interceptor is executed")
					handler(ctx, arg)
				}
			}, []string{
				`.+"msg":"Interceptor is executed".+"goroutine_id":\d+.+`,
				`.+"msg":"handler is executed: Argument_sample".+"goroutine_id":\d+.+`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := zap.NewTest(t)
			ctx := goroutine.NewContextWithGoroutineID()
			arg := "Argument_sample"
			interceptor := tt.newInterceptor(log)
			goroutine.RegisterInterceptor(interceptor)
			goroutine.GoWithContextAndArgument(ctx, arg, func(ctx context.Context, arg interface{}) {
				log.Infof(ctx, "handler is executed: %s", arg)
				done <- true
			})
			<-done
			log.AssertMessagePatterns(tt.outputLogs)
			log.ClearOutputs()
		})
	}
}
