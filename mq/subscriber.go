package mq

import "context"

type SubscribeHandler func(ctx context.Context, data []byte) error

type Subscriber interface {
	Register(topic string, h SubscribeHandler, queue ...string) error
}

type subscriber struct {
	c Clienter
}

func NewSubscriber(c Clienter) Subscriber {
	return &subscriber{
		c: c,
	}
}

func (s *subscriber) Register(topic string, h SubscribeHandler, queue ...string) error {
	return s.c.Subscribe(topic, h, queue...)
}
