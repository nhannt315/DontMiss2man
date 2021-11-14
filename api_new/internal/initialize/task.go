package initialize

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/nhannt315/real_estate_api/pkg/errors"
	"github.com/nhannt315/real_estate_api/pkg/logs"
)

// DefaultStopSigs is signals that server stop when received by default
var DefaultStopSigs = []os.Signal{syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}

// Tasks 終了時に行う処理(task)の管理/実行
type Tasks struct {
	log   logs.Logger
	tasks []Task
}

// Task is interface to control function's init / start / shutdown
type Task interface {
	Shutdown(ctx context.Context) error
	Name() string
}

// NewShutdownTasks shutdownTasksを返す
func NewShutdownTasks(log logs.Logger) *Tasks {
	return &Tasks{
		log:   log,
		tasks: make([]Task, 0),
	}
}

// Add taskを登録する
func (st *Tasks) Add(task Task) {

	st.tasks = append(st.tasks, task)
}

// ExecuteAll taskは後ろ(後から登録されたもの)から順に実行
func (st *Tasks) ExecuteAll(ctx context.Context) {
	for i := len(st.tasks) - 1; i >= 0; i-- {
		task := st.tasks[i]
		if task == nil {
			continue // 実行済みのものは無視
		}
		st.log.Infof(ctx, "execute shutdown task: %s", task.Name())
		if err := task.Shutdown(ctx); err != nil {
			st.log.Error(ctx, errors.Wrapf(err, "shutdown task: %s", task.Name()))
		}
		st.tasks[i] = nil
	}
}

// WaitForServerStop waits for server stop
func WaitForServerStop(ctx context.Context, log logs.Logger) {

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, DefaultStopSigs...)

	select {
	case sig := <-sigChan:
		log.Infof(ctx, "got stop sig: %s", sig.String())
	case <-ctx.Done():
		log.Infof(ctx, "context is done: %s", ctx.Err())
	}
}
