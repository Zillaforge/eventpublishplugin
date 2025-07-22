package plugininf

import (
	"context"

	"github.com/Zillaforge/eventpublishpluginclient/pb"
	"github.com/Zillaforge/toolkits/tracer"
	tkUtils "github.com/Zillaforge/toolkits/utilities"
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
