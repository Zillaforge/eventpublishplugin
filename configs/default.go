package configs

import (
	tkCfg "pegasus-cloud.com/aes/toolkits/configs"
	"pegasus-cloud.com/aes/toolkits/mviper"
)

func init() {
	mviper.SetDefault("plugin", map[string]interface{}{}, "-", tkCfg.TypeRestart, tkCfg.RegionGlobal)
}
