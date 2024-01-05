package storage

import "context"

type mdIncomingKey struct {
}

type Metadata map[string]string

func (md Metadata) Get(k string) string {
	return md[k]
}

func (md Metadata) Set(k, v string) {
	md[k] = v
}

func NewIncomingContext(ctx context.Context, md Metadata) context.Context {
	return context.WithValue(ctx, mdIncomingKey{}, md)
}

func FromIncomingContext(ctx context.Context) (md Metadata, ok bool) {
	md, ok = ctx.Value(mdIncomingKey{}).(Metadata)
	return
}
