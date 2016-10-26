package gowork

import (
	"time"
)

type (
	worker struct {
		pool  *workerPool
		Index int
	}
	workerPool struct {
		numWorkers int
		ch         chan *worker
	}
	iterator struct {
		minIterationInterval int
	}
)

func NewPool(size int) *workerPool {

	ch := make(chan *worker, size)
	p := workerPool{ch: ch, numWorkers: size}
	for i := 0; i < size; i++ {
		p.ch <- &worker{pool: &p, Index: i}
	}
	return &p
}

func (w *worker) Do(cb func()) {
	go func() {
		cb()
		w.release()
	}()
}
func (w *worker) DoWithIndex(cb func(i int)) {
	go func() {
		cb(w.Index)
		w.release()
	}()
}

//waits until a worker is available
func (p *workerPool) GetWorker() *worker {
	w := <-p.ch
	return w
}

//waits until a worker is available
func (p *workerPool) Workers() <-chan *worker {
	return p.ch
}

//releases
func (w *worker) release() {
	w.pool.ch <- w
}

func (p *workerPool) Sync() {

	workers := make([]*worker, p.numWorkers)
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
func (p *workerPool) RunInParallel(cb func()) {
	for i := 0; i < p.numWorkers; i++ {
		p.GetWorker().Do(cb)
	}
	p.Sync()
}
func (p *workerPool) RunInParallelWithIndex(cb func(i int)) {
	for i := 0; i < p.numWorkers; i++ {
		p.GetWorker().DoWithIndex(cb)
	}
	p.Sync()
}
func (p *workerPool) Size() int {
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
