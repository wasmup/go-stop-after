package main

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const d = 200 * time.Millisecond //  To stop a task after a period of time

func BenchmarkTimeSince(b *testing.B) {
	t0 := time.Now()
	var count = 0
	for i := 0; i < b.N; i++ {
		if time.Since(t0) < d {
			count++
		}
	}
	_ = count
}

func BenchmarkContext(b *testing.B) {
	var ctx, cancel = context.WithTimeout(context.Background(), d)
	defer cancel()
	var count = 0
	for i := 0; i < b.N; i++ {
		select {
		case <-ctx.Done():
			// break
		default:
			count++
		}
	}
	_ = count
}
func BenchmarkContextErr(b *testing.B) {
	var ctx, cancel = context.WithTimeout(context.Background(), d)
	defer cancel()
	var count = 0
	for i := 0; i < b.N; i++ {
		if ctx.Err() == nil {
			count++
		}
	}
	_ = count
}

func BenchmarkAfterFunc(b *testing.B) {
	var done uint32
	time.AfterFunc(d, func() { atomic.StoreUint32(&done, 1) })
	var count = 0
	for i := 0; i < b.N; i++ {
		if atomic.LoadUint32(&done) == 0 {
			count++
		}
	}
	_ = count
}

func BenchmarkDoneChannel(b *testing.B) {
	var done = make(chan struct{})
	time.AfterFunc(d, func() { close(done) })
	var count = 0
	for i := 0; i < b.N; i++ {
		select {
		case <-done:
			// break
		default:
			count++
		}
	}
	_ = count
}

type foo struct {
	sync.Mutex
	state bool
}

func (p *foo) end() {
	p.Lock()
	p.state = true
	p.Unlock()
}
func (p *foo) isDone() bool {
	var b bool
	p.Lock()
	b = p.state
	p.Unlock()
	return b
}
func BenchmarkAfterFuncMutex(b *testing.B) {
	var it = foo{}
	time.AfterFunc(d, func() { it.end() })
	var count = 0
	for i := 0; i < b.N; i++ {
		if it.isDone() {
			count++
		}
	}
	_ = count
}
