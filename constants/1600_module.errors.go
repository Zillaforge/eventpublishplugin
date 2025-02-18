package constants

import tkErr "pegasus-cloud.com/aes/toolkits/errors"

const (
	// 1600xxxx: module

	ModuleInternalServerErrCode         = 16000000
	ModuleInternalServerErrMsg          = "internal server error"
	ModuleConnectionTypeIsIllegalCode   = 16000001
	ModuleConnectionTypeIsIllegalMsg    = "connection type is illegal"
	ModuleServiceTypeIsNotSupportedCode = 16000002
	ModuleServiceTypeIsNotSupportedMsg  = "service type is not supported"
)

var (
	// 1600xxxx: module

	// 16000000(internal server error)
	ModuleInternalServerErr = tkErr.Error(ModuleInternalServerErrCode, ModuleInternalServerErrMsg)
	// 16000001(connection type is illegal)
	ModuleConnectionTypeIsIllegal = tkErr.Error(ModuleConnectionTypeIsIllegalCode, ModuleConnectionTypeIsIllegalMsg)
	// 16000002(service type is not supported)
	ModuleServiceTypeIsNotSupported = tkErr.Error(ModuleServiceTypeIsNotSupportedCode, ModuleServiceTypeIsNotSupportedMsg)
)
