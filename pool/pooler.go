package pool

import (
	"errors"
	"io"
)

var (
	ErrInvalidConfig = errors.New("invalid pool config")
	ErrPoolClosed    = errors.New("pool closed")
)

type factoryFunc func() (io.Closer, error)

type Pooler interface {
	Get() (io.Closer, error) //Get Resource
	Put(io.Closer) error     //Release Resource
	Close(io.Closer) error   //Close Resource
	Shutdown() error         //Shutdown pool
}
