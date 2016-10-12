package gowork

import (
	"time"
)

type (
	Worker struct {
		pool  *WorkerPool
		Index int
	}
	WorkerPool struct {
		numWorkers int
		ch         chan *Worker
	}
	iterator struct {
		minIterationInterval int
	}
)

func NewPool(size int) *WorkerPool {

	ch := make(chan *Worker, size)
	p := WorkerPool{ch: ch, numWorkers: size}
	for i := 0; i < size; i++ {
		go func(i int) {
			p.ch <- &Worker{pool: &p, Index: i}
		}(i)
	}
	return &p
}

func (w *Worker) Do(cb func()) {
	go func() {
		cb()
		w.release()
	}()
}

//waits until a worker is available
func (p *WorkerPool) GetWorker() *Worker {
	w := <-p.ch
	return w
}

//waits until a worker is available
func (p *WorkerPool) Workers() <-chan *Worker {
	return p.ch
}

//releases
func (w *Worker) release() {
	w.pool.ch <- w
}

func (p *WorkerPool) Sync() {

	workers := make([]*Worker, p.numWorkers)
	// obtained all the workers - all done
	for i := 0; i < p.numWorkers; i++ {
		workers[i] = p.GetWorker()
	}
	// release workers again
	go func() {
		for i := 0; i < p.numWorkers; i++ {
			workers[i].release()
		}
	}()
}
func (p *WorkerPool) RunInParallel(cb func()) {
	for i := 0; i < p.numWorkers; i++ {
		p.GetWorker().Do(cb)
	}
	p.Sync()
}

func (p *WorkerPool) Size() int {
	return p.numWorkers
}
func IterateEvery(minIterationInterval int) *iterator {
	return &iterator{minIterationInterval: minIterationInterval}
}
func (i *iterator) Run(cb func()) {
	for {
		iterationStartTime := time.Now()
		cb()
		elapsedTime := int(time.Since(iterationStartTime).Seconds())
		sleepInterval := i.minIterationInterval - elapsedTime
		if sleepInterval < 0 {
			sleepInterval = 0
		}
		time.Sleep(time.Duration(sleepInterval) * time.Second)
	}
}
