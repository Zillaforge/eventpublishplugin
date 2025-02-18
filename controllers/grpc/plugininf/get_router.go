package plugininf

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"pegasus-cloud.com/aes/eventpublishpluginclient/pb"
	"pegasus-cloud.com/aes/toolkits/tracer"
	tkUtils "pegasus-cloud.com/aes/toolkits/utilities"
)

/*
GetRouter ...

errors:
*/
func (m *Method) GetRouter(ctx context.Context, input *emptypb.Empty) (output *pb.GetRouterResponseList, err error) {
	var (
		funcName = tkUtils.NameOfFunction().Name()
	)

	_, f := tracer.StartWithContext(ctx, funcName)
	defer f(tracer.Attributes{
		"input":  &input,
		"output": &output,
		"err":    &err,
	})

	return &pb.GetRouterResponseList{}, nil
}
