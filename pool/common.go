package pool

import (
	"io"
	"sync"
)

type CommonPool struct {
	sync.Mutex
	pool    chan io.Closer
	maxOpen int
	numOpen int
	minOpen int
	closed  bool
	factory factoryFunc
}

func NewCommonPool(minOpen, maxOpen int, factory factoryFunc) (*CommonPool, error) {
	if maxOpen <= 0 || minOpen > maxOpen {
		return nil, ErrInvalidConfig
	}

	p := &CommonPool{
		maxOpen: maxOpen,
		minOpen: minOpen,
		factory: factory,
		pool:    make(chan io.Closer, maxOpen),
	}

	for i := 0; i < minOpen; i++ {
		closer, err := factory()
		if err != nil {
			continue
		}
		p.numOpen++
		p.pool <- closer
	}
	return p, nil
}

func (p *CommonPool) Get() (io.Closer, error) {
	if p.closed {
		return nil, ErrPoolClosed
	}
	for {
		closer, err := p.getOrCreate()
		return closer, err
	}
}

func (p *CommonPool) Put(closer io.Closer) error {
	if p.closed {
		return ErrPoolClosed
	}
	p.Lock()
	defer p.Unlock()

	p.pool <- closer
	return nil
}

func (p *CommonPool) Close(closer io.Closer) error {
	p.Lock()
	defer p.Unlock()

	err := closer.Close()
	p.numOpen--
	return err
}

func (p *CommonPool) Shutdown() error {
	if p.closed {
		return ErrPoolClosed
	}
	p.Lock()
	defer p.Unlock()

	close(p.pool)
	for closer := range p.pool {
		err := closer.Close()
		if err != nil {
			return err
		}
		p.numOpen--
	}
	p.closed = true
	return nil
}

func (p *CommonPool) getOrCreate() (io.Closer, error) {
	select {
	case closer := <-p.pool:
		return closer, nil
	default:
	}
	p.Lock()
	defer p.Unlock()

	if p.numOpen >= p.maxOpen {
		closer := <-p.pool
		return closer, nil
	}

	//New Connect
	closer, err := p.factory()
	if err != nil {
		return nil, err
	}
	p.numOpen++
	return closer, err
}
