package sentinel

import (
	"eventpublishplugin/services"

	"github.com/go-redis/redis/v8"
)

type Sentinel struct {
	Conn    *redis.Client
	Channel string
	_       struct{}
}

func (Sentinel) Name() string {
	return services.RedisSentinelKind
}
