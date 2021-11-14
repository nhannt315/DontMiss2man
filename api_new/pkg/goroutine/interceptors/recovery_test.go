package interceptors

import (
	"context"

	"github.com/nhannt315/real_estate_api/pkg/logs"
	"github.com/nhannt315/real_estate_api/pkg/logs/zap"

	"github.com/nhannt315/real_estate_api/pkg/goroutine"

	"testing"
)

func TestRecoveryInterceptor(t *testing.T) {
	tests := []struct {
		name       string
		newHandler func(log logs.Logger) goroutine.Handler
		outputLogs []string
	}{
		{"no panic",
			func(log logs.Logger) goroutine.Handler {
				return func(ctx context.Context, arg interface{}) {
					log.Infof(ctx, "no panic: %s", arg)
				}
			},
			[]string{
				`"msg":"no panic: Argument_sample"`,
			}},
		{"panic handler",
			func(log logs.Logger) goroutine.Handler {
				return func(ctx context.Context, arg interface{}) {
					panic("panic test")
				}
			},
			[]string{
				`.+"severity":"error".+,"msg":"panic. panic test".+`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := zap.NewTest(t)
			interceptor := RecoveryInterceptor(log)
			ctx := context.Background()
			arg := "Argument_sample"
			handler := tt.newHandler(log)
			interceptor(ctx, handler, arg)
			log.AssertMessagePatterns(tt.outputLogs)
			log.ClearOutputs()
		})
	}
}
