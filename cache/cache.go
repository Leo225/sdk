package cache

import (
	"context"
	"time"
)

type Cacher interface {
	Getter
	Setter
	Remover
}

type Getter interface {
	Get(ctx context.Context, key string, v interface{}) (err error)
	GetString(ctx context.Context, key string) (v string, err error)
	GetBytes(ctx context.Context, key string) (bytes []byte, err error)
}

type Setter interface {
	Set(ctx context.Context, key string, v interface{}, expiration ...time.Duration) (err error)
	SetString(ctx context.Context, key string, value string, expiration ...time.Duration) (err error)
	SetBytes(ctx context.Context, key string, bytes []byte, expiration ...time.Duration) (err error)
}

type Remover interface {
	Remove(ctx context.Context, keys ...string) (err error)
	RemoveMatch(ctx context.Context, match string) (err error)
}
