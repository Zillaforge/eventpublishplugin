package sentinel

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
func (s *Sentinel) Publish(ctx context.Context, input *modCom.PublishInput) (output *modCom.PublishOutput, err error) {
	var (
		funcName  = tkUtils.NameOfFunction().Name()
		requestID = utility.MustGetContextRequestID(ctx)
	)

	ctx, f := tracer.StartWithContext(ctx, funcName)
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
	if err := s.Conn.Publish(ctx, s.Channel, message).Err(); err != nil {
		zap.L().With(
			zap.String(cnt.Module, "s.Conn.Publish(...)"),
			zap.String(cnt.RequestID, requestID),
			zap.String("channel", s.Channel),
			zap.Any("message", message),
		).Error(err.Error())
		return nil, tkErr.New(cnt.ModuleInternalServerErr)
	}
	return nil, nil
}
