package logs

import (
	"context"

	"github.com/nhannt315/real_estate_api/pkg/errors"
)

// LogPanic is called in recovery, and print panic's stack log
func LogPanic(ctx context.Context, log Logger, panicErr interface{}) {
	if err, ok := panicErr.(error); ok {
		log.Error(ctx, err)
		return
	}

	log.Error(ctx, errors.Errorf("panic. %+v", panicErr))
}
