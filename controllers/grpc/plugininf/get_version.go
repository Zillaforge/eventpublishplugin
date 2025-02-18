package plugininf

import (
	"context"

	cnt "eventpublishplugin/constants"

	"google.golang.org/protobuf/types/known/emptypb"
	"pegasus-cloud.com/aes/eventpublishpluginclient/pb"
	"pegasus-cloud.com/aes/toolkits/tracer"
	tkUtils "pegasus-cloud.com/aes/toolkits/utilities"
)

/*
GetVersion ...

errors:
*/
func (m *Method) GetVersion(ctx context.Context, input *emptypb.Empty) (output *pb.GetVersionResponse, err error) {
	var (
		funcName = tkUtils.NameOfFunction().Name()
	)

	_, f := tracer.StartWithContext(ctx, funcName)
	defer f(tracer.Attributes{
		"input":  &input,
		"output": &output,
		"err":    &err,
	})

	return &pb.GetVersionResponse{Version: cnt.Version}, nil
}
