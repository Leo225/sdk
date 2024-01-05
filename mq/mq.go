package mq

import "context"

type Clienter interface {
	Publish(ctx context.Context, sub string, msg interface{}) error
	Subscribe(topic string, h SubscribeHandler, queue ...string) error
}
