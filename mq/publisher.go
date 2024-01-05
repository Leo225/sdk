package mq

import "context"

type Publisher interface {
	Publish(ctx context.Context, msg interface{}) error
}

type publisher struct {
	topic string
	c     Clienter
}

func NewPublisher(topic string, client Clienter) Publisher {
	return &publisher{
		topic: topic,
		c:     client,
	}
}

func (p *publisher) Publish(ctx context.Context, msg interface{}) error {
	return p.c.Publish(ctx, p.topic, msg)
}
