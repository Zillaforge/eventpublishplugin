package constants

const (
	Name           = "EventPublishPlugin"
	PascalCaseName = "EventPublishPlugin"
	SnakeCaseName  = "event_publish_plugin"
	KebabCaseName  = "event-publish-plugin"
	UpperAbbrName  = "EventPublishPlugin"
	LowerAbbrName  = "eventpublishplugin"

	Kind                 = PascalCaseName
	Version              = "0.1.1"
	GlobalConfigPath     = "etc/ASUS"
	GlobalConfigFilename = LowerAbbrName + ".yaml"

	UnixSocket = "/run/" + LowerAbbrName + ".sock"

	RequestID = "RequestID"
	Server    = "Server"
	GRPC      = "GRPC"
	Module    = "Module"
)
