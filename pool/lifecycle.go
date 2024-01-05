package pool

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"
)

type LifeCyclePool struct {
	sync.Mutex
	ticker       *time.Ticker
	tickerCtx    context.Context
	tickerCancel context.CancelFunc
	pool         []io.Closer
	poolIndex    int
	numOpen      int
	maxOpen      int
	minOpen      int
	closed       bool
	factory      factoryFunc
	maxLifeTime  time.Duration
}

type Connection struct {
	io.Closer
	time time.Time
	use  bool
}

func NewLifeCyclePool(minOpen, maxOpen int, maxLifeTime time.Duration, factory factoryFunc) (*LifeCyclePool, error) {
	if maxOpen <= 0 || minOpen > maxOpen {
		return nil, ErrInvalidConfig
	}

	tickerCtx, tickerCancel := context.WithCancel(context.Background())
	p := &LifeCyclePool{
		maxOpen:      maxOpen,
		minOpen:      minOpen,
		maxLifeTime:  maxLifeTime,
		factory:      factory,
		ticker:       time.NewTicker(1 * time.Second),
		tickerCtx:    tickerCtx,
		tickerCancel: tickerCancel,
		pool:         make([]io.Closer, 0),
		poolIndex:    0,
	}
	p.inactivate()

	for i := 0; i < minOpen; i++ {
		closer, err := factory()
		if err != nil {
			continue
		}
		p.numOpen++
		p.pool = append(p.pool, &Connection{
			time:   time.Now(),
			Closer: closer,
			use:    false,
		})
	}

	return p, nil
}

func (p *LifeCyclePool) Get() (io.Closer, error) {
	if p.closed {
		return nil, ErrPoolClosed
	}

	for {
		closer, err := p.getOrCreate()
		if err != nil {
			return nil, err
		}
		if closer == nil && err == nil {
			time.Sleep(time.Millisecond)
			continue
		}
		return closer, nil
	}
}

func (p *LifeCyclePool) Put(closer io.Closer) error {
	if p.closed {
		return ErrPoolClosed
	}
	p.Lock()
	defer p.Unlock()

	if len(p.pool) == p.maxOpen {
		return ErrInvalidConfig
	}

	c := closer.(*Connection)
	c.use = false
	c.time = time.Now()
	p.pool = append(p.pool, c)
	return nil
}

func (p *LifeCyclePool) Close(closer io.Closer) error {
	p.Lock()
	defer p.Unlock()

	err := closer.Close()
	p.numOpen--
	return err
}

func (p *LifeCyclePool) Shutdown() error {
	if p.closed {
		return ErrPoolClosed
	}
	p.Lock()
	defer p.Unlock()

	p.ticker.Stop()
	for i := 0; i < len(p.pool); i++ {
		c := p.pool[i].(*Connection)
		c.use = false
		err := c.Closer.Close()
		if err != nil {
			return err
		}
		p.numOpen--
	}
	p.tickerCancel()
	p.closed = true
	p.pool = make([]io.Closer, 0)
	p.poolIndex = 0
	return nil
}

func (p *LifeCyclePool) inactivate() {
	go func() {
		for {
			select {
			case <-p.ticker.C:
				p.Lock()
				for i := 0; i < len(p.pool); i++ {
					c := p.pool[i].(*Connection)
					d := time.Since(c.time)
					if !c.use && d >= p.maxLifeTime {
						c.Closer.Close()
						p.pool = append(p.pool[:i], p.pool[i+1:]...)
						p.numOpen--
					}
				}
				p.Unlock()
			case <-p.tickerCtx.Done():
				return
			}
		}
	}()
}

func (p *LifeCyclePool) getOrCreate() (io.Closer, error) {
	p.Lock()
	defer p.Unlock()

	if len(p.pool) > 0 {
		if p.poolIndex >= len(p.pool) {
			p.poolIndex = 0
		}
		fmt.Printf("len: %d, index: %d, numOpen: %d\n", len(p.pool), p.poolIndex, p.numOpen)

		c := p.pool[p.poolIndex].(*Connection)
		c.use = true
		c.time = time.Now()

		p.pool = append(p.pool[:p.poolIndex], p.pool[p.poolIndex+1:]...)
		p.poolIndex++
		return c, nil
	}

	if p.numOpen < p.maxOpen {
		closer, err := p.factory()
		if err != nil {
			return nil, err
		}
		p.numOpen++
		return &Connection{
			Closer: closer,
			time:   time.Now(),
			use:    true,
		}, nil
	}

	return nil, nil
}
