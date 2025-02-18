package plugininf

import (
	"context"
	"strings"

	cnt "eventpublishplugin/constants"

	"google.golang.org/protobuf/types/known/emptypb"
	"pegasus-cloud.com/aes/eventpublishpluginclient/pb"
	"pegasus-cloud.com/aes/toolkits/tracer"
	tkUtils "pegasus-cloud.com/aes/toolkits/utilities"
)

/*
GetName ...

errors:
*/
func (m *Method) GetName(ctx context.Context, input *emptypb.Empty) (output *pb.GetNameResponse, err error) {
	var (
		funcName = tkUtils.NameOfFunction().Name()
	)

	_, f := tracer.StartWithContext(ctx, funcName)
	defer f(tracer.Attributes{
		"input":  &input,
		"output": &output,
		"err":    &err,
	})
	return &pb.GetNameResponse{Name: strings.ToLower(cnt.Kind)}, nil
}
