package task

import (
	"errors"
	"sync/atomic"
)

const (
	Running = 1
	Stop    = 0
)

type TaskHandler = func(closer chan struct{})

func NewTask(handler TaskHandler) *Task {
	return &Task{
		handler: handler,
		state:   Stop,
	}
}

type Task struct {
	handler TaskHandler
	closer  chan struct{}
	state   uint32
}

func (t *Task) Run() error {
	if !atomic.CompareAndSwapUint32(&t.state, Stop, Running) {
		return errors.New("task is already running")
	}
	t.closer = make(chan struct{})
	go func() {
		t.handler(t.closer)
	}()
	return nil
}

func (t *Task) Stop() error {
	if !atomic.CompareAndSwapUint32(&t.state, Running, Stop) {
		return errors.New("task is already stop")
	}
	close(t.closer)
	return nil
}
