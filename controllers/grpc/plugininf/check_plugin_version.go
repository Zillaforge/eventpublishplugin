package plugininf

import (
	"context"
	cnt "eventpublishplugin/constants"

	"google.golang.org/protobuf/types/known/emptypb"
	"pegasus-cloud.com/aes/eventpublishpluginclient/pb"
	"pegasus-cloud.com/aes/toolkits/mviper"
	"pegasus-cloud.com/aes/toolkits/tracer"
	tkUtils "pegasus-cloud.com/aes/toolkits/utilities"
)

/*
CheckPluginVersion ...

error:
*/
func (m *Method) CheckPluginVersion(ctx context.Context, input *emptypb.Empty) (output *pb.CheckVersionResponse, err error) {
	var (
		funcName = tkUtils.NameOfFunction().Name()
	)

	_, f := tracer.StartWithContext(ctx, funcName)
	defer f(tracer.Attributes{
		"input":  &input,
		"output": &output,
		"err":    &err,
	})

	isMatched := false
	pluginMap := mviper.GetStringMap("plugin")
	if pluginMap != nil {
		if version, ok := pluginMap["version"]; ok {
			if value, ok := version.(string); ok {
				if value == cnt.Version {
					isMatched = true
				}
			}
		}
	}
	if isMatched {
		return &pb.CheckVersionResponse{IsMatch: true}, nil
	}
	return &pb.CheckVersionResponse{IsMatch: false}, nil
}
