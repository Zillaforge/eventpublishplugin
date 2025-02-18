package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"pegasus-cloud.com/aes/toolkits/mviper"
	"pegasus-cloud.com/aes/toolkits/tracer"

	cnt "eventpublishplugin/constants"
	"eventpublishplugin/logger"
	"eventpublishplugin/services"

	mod "eventpublishplugin/module"

	tkErr "pegasus-cloud.com/aes/toolkits/errors"
)

var (
	srvGRPC     *grpc.Server
	_socketFile = ""
)

func Run() {
	_socketFile = cnt.UnixSocket
	prepareUpstreamServices()
	startInitPlugin()
	startGRPCServer()
	signAction()
}

func signAction() {
	zap.L().Info("EventPublish Plugin process is running !!")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sign := <-quit
		switch sign {
		case syscall.SIGINT, syscall.SIGTERM:
			zap.L().Info("Shutdown Service.")
			// If plugin is stopped. the socket file must be removed.
			os.Remove(_socketFile)
			return
		case syscall.SIGUSR2:
			startGRPCServer()
		default:
			zap.L().Info(fmt.Sprintln("Other Signal ", sign))
			return
		}
	}
}

func startInitPlugin() {
	pluginConfigs := mviper.Get("plugin")
	if pluginConfigs == nil {
		zap.L().With(
			zap.String(cnt.Server, "pluginConfigs.(map[string]interface{})"),
		).Error(cnt.ServerPluginConfigsIsEmptyErrMsg)
		os.Exit(1)
	}

	configs, ok := pluginConfigs.(map[string]interface{})
	if !ok {
		zap.L().With(
			zap.String(cnt.Server, "pluginConfigs.(map[string]interface{})"),
		).Error(cnt.ServerPluginConfigsConvertTypeIsFailedErrMsg)
		os.Exit(1)
	}
	serviceName, ok := configs["service"].(string)
	if ok {
		mod.Init(serviceName)
	}
	socketPath, ok := configs["socket_path"].(string)
	if ok {
		if socketPath != "" {
			_socketFile = socketPath
		}
	}
}

func startGRPCServer() {
	tlsEnable := false
	if tlsEnable {
		c, err := credentials.NewServerTLSFromFile("", "")
		if err != nil {
			zap.L().With(
				zap.String(cnt.Server, "credentials.NewServerTLSFromFile(...)"),
				zap.String("certFile", ""),
				zap.String("keyFile", ""),
			).Error(err.Error())
			os.Exit(1)
		}
		srvGRPC = grpc.NewServer(
			grpc.Creds(c),
			grpc.ChainUnaryInterceptor(tracer.RequestIDParser()),
			grpc.UnaryInterceptor(tracer.NewGRPCUnaryServerInterceptor()),
		)
	} else {
		// add action parser middleware to array
		srvGRPC = grpc.NewServer(
			grpc.ChainUnaryInterceptor(tracer.RequestIDParser()),
			grpc.UnaryInterceptor(tracer.NewGRPCUnaryServerInterceptor()),
		)
	}
	registerGRPCServer(srvGRPC)
	os.Remove(_socketFile)
	lis2, err := net.Listen("unix", _socketFile)
	if err != nil {
		zap.L().With(
			zap.String(cnt.Server, "net.Listen(...)"),
			zap.String("network", "unix"),
			zap.String("address", _socketFile),
		).Error(err.Error())
		os.Exit(1)
	}
	go func() {
		if err := srvGRPC.Serve(lis2); err != nil {
			zap.L().With(
				zap.String(cnt.Server, "srvGRPC.Serve(...)"),
			).Error(err.Error())
			os.Exit(1)
		}
	}()
}

func prepareUpstreamServices() {
	// Initialize Logger
	logger.Init("eventpublish_plugin.log")
	logger.InitEventLogger("eventpublish_events.log")

	// Initialize Tracer
	if mviper.GetBool("plugin.tracer.enable") {
		tracer.Init(&tracer.Config{
			ServiceName: mviper.GetString("plugin.instance"),
			Endpoint:    mviper.GetString("plugin.tracer.collector_endpoint"),
			Timeout:     mviper.GetInt("plugin.tracer.timeout"),
		})
	}

	// Initialize all of Services
	if err := services.InitServices(); err != nil {
		if e, ok := tkErr.IsError(err); ok {
			switch e.Code() {
			}
		}
		zap.L().With(
			zap.String(cnt.Server, "services.InitServices()"),
		).Error(err.Error())
		os.Exit(1)
	}
}
