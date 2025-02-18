package services

import (
	cnt "eventpublishplugin/constants"

	"pegasus-cloud.com/aes/pegasusmsgqueueclient/msgqueue"
	"pegasus-cloud.com/aes/pegasusmsgqueueclient/msgqueue/rabbitmq"
	"pegasus-cloud.com/aes/pegasusmsgqueueclient/msgqueue/rabbitmq/core"
)

type RabbitMQInput struct {
	Name          string
	Account       interface{}
	Password      interface{}
	Host          interface{}
	ManageHost    interface{}
	Timeout       interface{}
	RPCTimeout    interface{}
	Vhost         interface{}
	ConnectionNum interface{}
	ChannelNum    interface{}
	ConsumerConn  interface{}
	ReplicaNum    interface{}
}

func UnmarshalRabbitMQ(s map[string]interface{}) {
	InitRabbitMQ(&RabbitMQInput{
		Name:          s["name"].(string),
		Account:       s["account"],
		Password:      s["password"],
		Host:          s["host"],
		ManageHost:    s["manage_host"],
		Timeout:       s["timeout"],
		RPCTimeout:    s["rpc_timeout"],
		Vhost:         s["vhost"],
		ConnectionNum: s["connection_num"],
		ChannelNum:    s["channel_num"],
		ConsumerConn:  s["consumer_conn"],
		ReplicaNum:    s["replica_num"],
	})

}

func InitRabbitMQ(input *RabbitMQInput) {
	conn := newRabbitMQ(&newRabbitMQInput{
		account:       input.Account,
		password:      input.Password,
		host:          input.Host,
		manageHost:    input.ManageHost,
		timeout:       input.Timeout,
		rpcTimeout:    input.RPCTimeout,
		vhost:         input.Vhost,
		connectionNum: input.ConnectionNum,
		channelNum:    input.ChannelNum,
		consumerConn:  input.ConsumerConn,
		replicaNum:    input.ReplicaNum,
	})
	ServiceMap[input.Name] = &Service{
		Kind: RabbitMQKind,
		Conn: conn,
	}

}
func NewRabbitMQ(input *RabbitMQInput) (conn *msgqueue.Handler) {
	return newRabbitMQ(&newRabbitMQInput{
		account:       input.Account,
		password:      input.Password,
		host:          input.Host,
		manageHost:    input.ManageHost,
		timeout:       input.Timeout,
		rpcTimeout:    input.RPCTimeout,
		vhost:         input.Vhost,
		connectionNum: input.ConnectionNum,
		channelNum:    input.ChannelNum,
		consumerConn:  input.ConsumerConn,
		replicaNum:    input.ReplicaNum,
	})

}

type newRabbitMQInput struct {
	account       interface{}
	password      interface{}
	host          interface{}
	manageHost    interface{}
	timeout       interface{}
	rpcTimeout    interface{}
	vhost         interface{}
	connectionNum interface{}
	channelNum    interface{}
	consumerConn  interface{}
	replicaNum    interface{}
}

func newRabbitMQ(input *newRabbitMQInput) (conn *msgqueue.Handler) {
	handler := msgqueue.New(
		&rabbitmq.Provider{
			AMQP: core.AMQP{
				Kind:                   cnt.Kind,
				Account:                interface2string(input.account),
				Password:               interface2string(input.password),
				Host:                   interface2string(input.host),
				ManageHost:             interface2string(input.manageHost),
				Timeout:                interface2integer(input.timeout),
				RPCTimeout:             interface2integer(input.rpcTimeout),
				Vhost:                  interface2string(input.vhost),
				OperationConnectionNum: interface2integer(input.connectionNum),
				ChannelNum:             interface2integer(input.channelNum),
				ConsumerConnectionNum:  interface2integer(input.consumerConn),
				ReplicaNum:             interface2integer(input.replicaNum),
			},
		},
	)
	return &handler
}
