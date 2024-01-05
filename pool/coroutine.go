package pool

import "sync"

// Coroutine Pool
type Coroutine struct {
	queue chan int
	wg    *sync.WaitGroup
}

func NewCoroutine(size int) *Coroutine {
	if size <= 0 {
		size = 1
	}
	return &Coroutine{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *Coroutine) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}

	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *Coroutine) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *Coroutine) Wait() {
	p.wg.Wait()
}
