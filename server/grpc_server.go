package server

import (
	ctlPluginInf "eventpublishplugin/controllers/grpc/plugininf"

	"google.golang.org/grpc"
	"pegasus-cloud.com/aes/eventpublishpluginclient/pb"
)

func registerGRPCServer(srv *grpc.Server) {
	pb.RegisterEventPublishPluginInterfaceCRUDControllerServer(srv, new(ctlPluginInf.Method))
}
