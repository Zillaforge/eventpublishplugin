package plugininf

import (
	"context"

	"pegasus-cloud.com/aes/eventpublishpluginclient/pb"
	"pegasus-cloud.com/aes/toolkits/tracer"
	tkUtils "pegasus-cloud.com/aes/toolkits/utilities"
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
