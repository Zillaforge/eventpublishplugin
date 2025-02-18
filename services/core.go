package services

import (
	cnt "eventpublishplugin/constants"
	"fmt"
	"reflect"

	"go.uber.org/zap"
	tkErr "pegasus-cloud.com/aes/toolkits/errors"
	"pegasus-cloud.com/aes/toolkits/mviper"
)

type Service struct {
	Kind string
	Conn interface{}
}

const (
	RedisSentinelKind string = "redis_sentinel"
	RabbitMQKind      string = "rabbitmq"
)

var ServiceMap = make(map[string]*Service)

func InitServices() (err error) {
	zap.L().Info("Initialize all of services")
	if mviper.Get("plugin") != nil {
		for _, service := range mviper.Get("plugin.services").([]interface{}) {
			if service == nil {
				continue
			}
			s := interface2map(service)
			for _, key := range []string{"name", "kind"} {
				if s[key] == nil {
					return tkErr.New(cnt.ServiceNameIsRequiredErr)
				}
				if key == "name" {
					if reflect.TypeOf(s[key]).Kind() != reflect.String {
						return tkErr.New(cnt.ServiceNameMustBeAStringErr)
					}
					if ServiceMap[s["name"].(string)] != nil {
						return tkErr.New(cnt.ServiceNameIsRepeatedErr)
					}
				}
			}
			switch s["kind"] {
			case RedisSentinelKind:
				UnmarshalRedisSentinel(s)
			case RabbitMQKind:
				UnmarshalRabbitMQ(s)
			}
			zap.L().Info(fmt.Sprintf("Service %s is initialized", s["name"]))
		}
	}
	return nil
}

func interface2map(input interface{}) (m map[string]interface{}) {
	if m == nil {
		m = make(map[string]interface{})
	}

	val := reflect.ValueOf(input)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() == reflect.Map {
		for _, e := range val.MapKeys() {
			if k, ok := e.Interface().(string); ok {
				m[k] = val.MapIndex(e).Interface()
			}
		}
	}
	return
}

func interface2string(input interface{}) string {
	if input != nil {
		switch t := input.(type) {
		case string:
			return t
		}
	}
	return ""
}

func interface2integer(input interface{}) int {
	if input != nil {
		switch t := input.(type) {
		case int:
			return t
		case int8:
			return int(t) // standardizes across systems
		case int16:
			return int(t) // standardizes across systems
		case int32:
			return int(t) // standardizes across systems
		case int64:
			return int(t) // standardizes across systems
		}
	}
	return 0
}
