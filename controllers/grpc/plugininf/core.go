package plugininf

import (
	"pegasus-cloud.com/aes/eventpublishpluginclient/pb"
)

// Method is implement all methods as pb.EventPublishPluginInterfaceCRUDControllerServer
type Method struct {
	// Embed EventPublishPluginInterfaceCRUDControllerServer to have UnimplementedEventPublishPluginInterfaceCRUDControllerServer()
	pb.UnimplementedEventPublishPluginInterfaceCRUDControllerServer
}

// Verify interface compliance at compile time where appropriate
var _ pb.EventPublishPluginInterfaceCRUDControllerServer = (*Method)(nil)
