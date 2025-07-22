package rabbitmq

import (
	"eventpublishplugin/services"

	"github.com/Zillaforge/pegasusmsgqueueclient/msgqueue"
)

type Message struct {
	Conn       *msgqueue.Handler
	Exchange   string
	RoutingKey string
	Headers    map[string]interface{}
	_          struct{}
}

func (Message) Name() string {
	return services.RabbitMQKind
}
