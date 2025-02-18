package plugininf

import (
	"context"

	"pegasus-cloud.com/aes/eventpublishpluginclient/pb"
	"pegasus-cloud.com/aes/toolkits/tracer"
	tkUtils "pegasus-cloud.com/aes/toolkits/utilities"
)

/*
CallGRPCRouter ...

errors:
*/
func (m *Method) CallGRPCRouter(ctx context.Context, input *pb.RPCRouterRequest) (output *pb.RPCRouterResponse, err error) {
	var (
		funcName = tkUtils.NameOfFunction().Name()
	)

	_, f := tracer.StartWithContext(ctx, funcName)
	defer f(tracer.Attributes{
		"input":  &input,
		"output": &output,
		"err":    &err,
	})
	return &pb.RPCRouterResponse{}, nil
}
