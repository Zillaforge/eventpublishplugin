package module

import (
	cnt "eventpublishplugin/constants"
	modCom "eventpublishplugin/module/common"
	"eventpublishplugin/module/rabbitmq"
	"eventpublishplugin/module/sentinel"
	"eventpublishplugin/services"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"pegasus-cloud.com/aes/pegasusmsgqueueclient/msgqueue"
)

var _provider modCom.Provider

func Init(name string) {
	_provider = new(name)
}

func New(name string) (p modCom.Provider) {
	return new(name)
}

func new(name string) (p modCom.Provider) {
	switch services.ServiceMap[name].Kind {
	case services.RedisSentinelKind:
		conn, ok := services.ServiceMap[name].Conn.(*redis.Client)
		if ok {
			p = &sentinel.Sentinel{
				Conn:    conn,
				Channel: modCom.RedisChannel,
			}
			return p
		}
		zap.L().With(
			zap.String(cnt.Module, "services.ServiceMap[name].Conn.(*redis.Client)"),
		).Error(cnt.ModuleConnectionTypeIsIllegalMsg)
	case services.RabbitMQKind:
		conn, ok := services.ServiceMap[name].Conn.(*msgqueue.Handler)
		if ok {
			p = &rabbitmq.Message{
				Conn:       conn,
				Exchange:   modCom.RabbitMQExchange,
				RoutingKey: modCom.RabbitMQRoutingKey,
			}
			return p
		}
		zap.L().With(
			zap.String(cnt.Module, "services.ServiceMap[name].Conn.(*msgqueue.Handler)"),
		).Error(cnt.ModuleConnectionTypeIsIllegalMsg)
	default:
		zap.L().Error(cnt.ModuleServiceTypeIsNotSupportedMsg)
	}
	return nil
}

func Use() (p modCom.Provider) {
	return _provider
}
