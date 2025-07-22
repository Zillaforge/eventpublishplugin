package rabbitmq

import (
	"context"
	"encoding/json"
	cnt "eventpublishplugin/constants"
	modCom "eventpublishplugin/module/common"
	"eventpublishplugin/utility"

	"go.uber.org/zap"
	tkErr "github.com/Zillaforge/toolkits/errors"
	"github.com/Zillaforge/toolkits/tracer"
	tkUtils "github.com/Zillaforge/toolkits/utilities"
)

/*
Publish ...

errors:
- 16000000(internal server error)
*/
func (m *Message) Publish(ctx context.Context, input *modCom.PublishInput) (output *modCom.PublishOutput, err error) {
	var (
		funcName  = tkUtils.NameOfFunction().Name()
		requestID = utility.MustGetContextRequestID(ctx)
	)

	_, f := tracer.StartWithContext(ctx, funcName)
	defer f(tracer.Attributes{
		"input":  &input,
		"output": &output,
		"err":    &err,
	})

	message, err := json.Marshal(input)
	if err != nil {
		zap.L().With(
			zap.String(cnt.Module, "json.Marshal(...)"),
			zap.String(cnt.RequestID, requestID),
			zap.Any("v", input),
		).Error(err.Error())
		return nil, tkErr.New(cnt.ModuleInternalServerErr).WithInner(err)
	}

	if err := m.Conn.Queue().SendMessage(m.Exchange, m.RoutingKey, requestID, message, m.Headers); err != nil {
		zap.L().With(
			zap.String(cnt.Module, "m.Conn.Queue().SendMessage(...)"),
			zap.String(cnt.RequestID, requestID),
			zap.String("exchange", m.Exchange),
			zap.String("routing-key", m.RoutingKey),
			zap.Any("message", message),
		).Error(err.Error())
		return nil, tkErr.New(cnt.ModuleInternalServerErr)
	}
	return nil, nil
}
