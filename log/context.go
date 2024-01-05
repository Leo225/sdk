package log

import "context"

type (
	traceIDKey struct{}
	spanIDKey  struct{}
	userIDKey  struct{}
)

func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func FromTraceIDContext(ctx context.Context) (traceID string, ok bool) {
	traceID, ok = ctx.Value(traceIDKey{}).(string)
	return
}

func NewSpanIDContext(ctx context.Context, spanID string) context.Context {
	return context.WithValue(ctx, spanIDKey{}, spanID)
}

func FromSpanIDContext(ctx context.Context) (spanID string, ok bool) {
	spanID, ok = ctx.Value(spanIDKey{}).(string)
	return
}

func NewUserIDContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func FromUserIDContext(ctx context.Context) (userID string, ok bool) {
	userID, ok = ctx.Value(userIDKey{}).(string)
	return
}

func CopyContext(ctx context.Context) context.Context {
	parent := context.Background()
	if traceID, ok := FromTraceIDContext(ctx); ok {
		parent = NewTraceIDContext(ctx, traceID)
	}
	if spanID, ok := FromSpanIDContext(ctx); ok {
		parent = NewSpanIDContext(ctx, spanID)
	}
	if userID, ok := FromUserIDContext(ctx); ok {
		parent = NewUserIDContext(ctx, userID)
	}

	return parent
}
