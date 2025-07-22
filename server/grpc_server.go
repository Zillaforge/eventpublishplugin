package server

import (
	ctlPluginInf "eventpublishplugin/controllers/grpc/plugininf"

	"google.golang.org/grpc"
	"github.com/Zillaforge/eventpublishpluginclient/pb"
)

func registerGRPCServer(srv *grpc.Server) {
	pb.RegisterEventPublishPluginInterfaceCRUDControllerServer(srv, new(ctlPluginInf.Method))
}
