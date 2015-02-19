package gowork

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	Worker struct {
		ch chan *Worker
	}
	WorkerPool struct {
		numWorkers int
		ch         chan *Worker
	}
)

func NewPool(size int) *WorkerPool {
	ch := make(chan *Worker)
	for i := 0; i < size; i++ {
		go func() { ch <- &Worker{ch} }()
	}
	return &WorkerPool{ch: ch}
}

func (w *Worker) Do(cb func()) {
	cb()
	go func() { w.ch <- w }()
}

//waits until a worker is available
func (p *WorkerPool) GetWorker() *Worker {
	return <-p.ch
}

//waits until a worker is available
func (p *WorkerPool) Workers() <-chan *Worker {
	return p.ch
}

//releases
func (p *WorkerPool) ReleaseWorker() {
	go func() {
		p.ch <- &Worker{ch: p.ch}
	}()
}

func (p *WorkerPool) Run(cb func()) {
	for w := range p.ch {
		go w.Do(cb)
	}
}
