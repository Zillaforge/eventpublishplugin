package plugininf

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"github.com/Zillaforge/eventpublishpluginclient/pb"
	"github.com/Zillaforge/toolkits/tracer"
	tkUtils "github.com/Zillaforge/toolkits/utilities"
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
