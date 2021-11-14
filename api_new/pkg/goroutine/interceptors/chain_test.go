package interceptors

import (
	"context"
	"testing"

	"github.com/nhannt315/real_estate_api/pkg/logs"
	"github.com/nhannt315/real_estate_api/pkg/logs/zap"

	"github.com/nhannt315/real_estate_api/pkg/goroutine"
)

type testInterceptor struct {
	before, after string
	log           logs.Logger
}

func (i *testInterceptor) interceptor(ctx context.Context, handler goroutine.Handler, arg interface{}) {
	i.log.Infof(ctx, "before: %s", i.before)
	handler(ctx, arg)
	i.log.Infof(ctx, "after: %s", i.after)
}

func TestChain(t *testing.T) {
	tests := []struct {
		name         string
		interceptors []*testInterceptor
		outputLogs   []string
	}{
		{"no interceptor", nil,
			[]string{
				"handler: Argument_sample",
			}},
		{"1 interceptor", []*testInterceptor{
			{before: "i_1_before", after: "i_1_after"},
		},
			[]string{
				"before: i_1_before",
				"handler: Argument_sample",
				"after: i_1_after",
			}},
		{"3 interceptor", []*testInterceptor{
			{before: "i_1_before", after: "i_1_after"},
			{before: "i_2_before", after: "i_2_after"},
			{before: "i_3_before", after: "i_3_after"},
		}, []string{
			"before: i_1_before",
			"before: i_2_before",
			"before: i_3_before",
			"handler: Argument_sample",
			"after: i_3_after",
			"after: i_2_after",
			"after: i_1_after",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := zap.NewTest(t)
			ctx := context.Background()
			arg := "Argument_sample"
			handler := func(ctx context.Context, arg interface{}) {
				log.Infof(ctx, "handler: %s", arg)
			}

			is := make([]goroutine.Interceptor, 0)
			for _, i := range tt.interceptors {
				i.log = log
				is = append(is, i.interceptor)
			}

			chain := Chain(is...)
			chain(ctx, handler, arg)

			log.AssertMessages(tt.outputLogs)
		})
	}
}
