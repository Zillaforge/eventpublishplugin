package plugininf

import (
	"context"
	cnt "eventpublishplugin/constants"
	"eventpublishplugin/logger"
	mod "eventpublishplugin/module"
	modCom "eventpublishplugin/module/common"
	"eventpublishplugin/utility"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	cCnt "pegasus-cloud.com/aes/eventpublishpluginclient/constants"
	"pegasus-cloud.com/aes/eventpublishpluginclient/pb"
	tkErr "pegasus-cloud.com/aes/toolkits/errors"
	"pegasus-cloud.com/aes/toolkits/tracer"
	tkUtils "pegasus-cloud.com/aes/toolkits/utilities"
)

/*
Reconcile ...

errors:
- 14000000(internal server error)
*/
func (m *Method) Reconcile(ctx context.Context, input *pb.ReconcileRequest) (output *emptypb.Empty, err error) {
	var (
		funcName  = tkUtils.NameOfFunction().Name()
		requestID = utility.MustGetContextRequestID(ctx)
	)

	reqValue := string(input.Request)
	respValue := string(input.Response)
	ctx, f := tracer.StartWithContext(ctx, funcName)
	defer f(tracer.Attributes{
		"action":    &input.Action,
		"metadata":  &input.Metadata,
		"reqValue":  &reqValue,
		"respValue": &respValue,
		"err":       &err,
	})

	logger.Use().Info(fmt.Sprintf("%s | %s | %s | %s, | %s",
		input.Metadata[tracer.RequestID],
		mod.Use().Name(),
		input.Action,
		reqValue,
		respValue,
	))

	publishInput := &modCom.PublishInput{
		Action:   input.Action,
		Metadata: input.Metadata,
		Request:  input.Request,
		Response: input.Response,
	}
	if _, err := mod.Use().Publish(ctx, publishInput); err != nil {
		zap.L().With(
			zap.String(cnt.GRPC, "mod.Use().Publish(...)"),
			zap.String(cnt.RequestID, requestID),
			zap.Any("input", publishInput),
		).Error(err.Error())
		return &emptypb.Empty{}, tkErr.New(cCnt.GRPCInternalServerErr)
	}
	return &emptypb.Empty{}, nil
}
