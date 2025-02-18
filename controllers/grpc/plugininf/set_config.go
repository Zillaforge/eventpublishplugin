package plugininf

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"pegasus-cloud.com/aes/eventpublishpluginclient/pb"
	"pegasus-cloud.com/aes/toolkits/tracer"
	tkUtils "pegasus-cloud.com/aes/toolkits/utilities"
)

/*
SetConfig ...

errors:
*/
func (m *Method) SetConfig(ctx context.Context, input *pb.SetConfigRequest) (output *emptypb.Empty, err error) {
	var (
		funcName = tkUtils.NameOfFunction().Name()
	)

	_, f := tracer.StartWithContext(ctx, funcName)
	defer f(tracer.Attributes{
		"input":  &input,
		"output": &output,
		"err":    &err,
	})

	return &emptypb.Empty{}, nil
}
