package plugininf

import (
	"context"

	"github.com/Zillaforge/eventpublishpluginclient/pb"
	"github.com/Zillaforge/toolkits/tracer"
	tkUtils "github.com/Zillaforge/toolkits/utilities"
)

/*
EnableHttpRouter ...

errors:
*/
func (m *Method) EnableHttpRouter(ctx context.Context, input *pb.HttpRequestInfo) (output *pb.HttpResponseInfo, err error) {
	var (
		funcName = tkUtils.NameOfFunction().Name()
	)

	_, f := tracer.StartWithContext(ctx, funcName)
	defer f(tracer.Attributes{
		"input":  &input,
		"output": &output,
		"err":    &err,
	})
	return &pb.HttpResponseInfo{}, nil
}
