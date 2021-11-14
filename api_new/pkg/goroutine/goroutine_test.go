package goroutine

import (
	"context"
	"testing"
)

func BenchmarkSetGoroutineID(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SetGoroutineID(ctx)
	}
}
