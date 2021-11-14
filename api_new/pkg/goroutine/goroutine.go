package goroutine

import (
	"context"
	"sync"
)

// ContextKey コンテキストのキー
type ContextKey string

const (
	goroutineKey ContextKey = "goroutine_id"
)

var (
	m            = new(sync.Mutex)
	maxID uint64 = 1
)

// GetContextKey Context内のgoroutineID用のキー(testからの利用向け)
func GetContextKey() ContextKey {
	return goroutineKey
}

// Go Contextを作成した上で引数にinterface{}を持つGoroutineを作成する
func Go(f func(ctx context.Context)) {
	ctx := NewContextWithGoroutineID()

	f1 := func(ctx context.Context, arg interface{}) {
		f(ctx)
	}
	goWithInterceptor(ctx, nil, f1)
}

// GoWithArgument Contextを作成した上で引数にcontextとinterface{}を持つGoroutineを作成する
func GoWithArgument(g interface{}, f func(ctx context.Context, arg interface{})) {
	ctx := NewContextWithGoroutineID()
	goWithInterceptor(ctx, g, f)
}

// GoWithContext Contextを受け取り引数にinterface{}を持つGoroutineを作成する
func GoWithContext(ctx context.Context, f func(ctx context.Context)) {
	ctx = SetGoroutineID(ctx)
	f1 := func(ctx context.Context, arg interface{}) {
		f(ctx)
	}
	goWithInterceptor(ctx, nil, f1)
}

// GoWithContextAndArgument  Contextを受け取り引数にcontextとinterface{}を持つGoroutineを作成する
func GoWithContextAndArgument(ctx context.Context, g interface{}, f func(ctx context.Context, arg interface{})) {
	ctx = SetGoroutineID(ctx)

	goWithInterceptor(ctx, g, f)
	//	go f(ctx, g)
}

// IDValue GoroutineIDを取得する
//	取得できない場合、 0  が返る。
func IDValue(ctx context.Context) uint64 {
	if ctx == nil {
		return 0
	}
	goroutineID, ok := ctx.Value(goroutineKey).(uint64)
	if !ok || goroutineID == 0 {
		return 0
	}
	return goroutineID
}

// SetGoroutineID GoroutineIDをcontextにsetする
func SetGoroutineID(ctx context.Context) context.Context {
	m.Lock()
	ctx = context.WithValue(ctx, goroutineKey, maxID)
	maxID++
	m.Unlock()
	return ctx
}

// SetGoroutineIDWithValue 指定されたGoroutineIDをcontextにsetする
func SetGoroutineIDWithValue(parent context.Context, gid uint64) context.Context {
	return context.WithValue(parent, goroutineKey, gid)
}

// NewContextWithGoroutineID GoroutineIDを持つContextを生成する
func NewContextWithGoroutineID() context.Context {
	ctx := context.Background()
	return SetGoroutineID(ctx)
}
