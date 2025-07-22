package constants

import tkErr "github.com/Zillaforge/toolkits/errors"

const (
	// 1102xxxx: server

	ServerInternalServerErrCode                   = 11020000
	ServerInternalServerErrMsg                    = "internal server error"
	ServerPluginConfigsConvertTypeIsFailedErrCode = 11020001
	ServerPluginConfigsConvertTypeIsFailedErrMsg  = "pluginConfigs convert type is failed"
	ServerPluginConfigsIsEmptyErrCode             = 11020002
	ServerPluginConfigsIsEmptyErrMsg              = "pluginConfigs is empty"
	ServerPluginProviderInfoMustBeSetErrCode      = 11010003
	ServerPluginProviderInfoMustBeSetErrMsg       = "plugin provider information must be set"
)

var (
	// 1102xxxx: server

	// 11020000(internal server error)
	ServerInternalServerErr = tkErr.Error(ServerInternalServerErrCode, ServerInternalServerErrMsg)
	// 11020001(pluginConfigs convert type is failed)
	ServerPluginConfigsConvertTypeIsFailedErr = tkErr.Error(ServerPluginConfigsConvertTypeIsFailedErrCode, ServerPluginConfigsConvertTypeIsFailedErrMsg)
	// 11020002(pluginConfigs is empty)
	ServerPluginConfigsIsEmptyErr = tkErr.Error(ServerPluginConfigsIsEmptyErrCode, ServerPluginConfigsIsEmptyErrMsg)
	// 11020003(plugin provider information must be set)
	ServerPluginProviderInfoMustBeSetErr = tkErr.Error(ServerPluginProviderInfoMustBeSetErrCode, ServerPluginProviderInfoMustBeSetErrMsg)
)
