package common

import "context"

type Provider interface {
	Name() string
	Publish(ctx context.Context, input *PublishInput) (output *PublishOutput, err error)
}

var RedisChannel string

var RabbitMQExchange string
var RabbitMQRoutingKey string
