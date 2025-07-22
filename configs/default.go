package configs

import (
	tkCfg "github.com/Zillaforge/toolkits/configs"
	"github.com/Zillaforge/toolkits/mviper"
)

func init() {
	mviper.SetDefault("plugin", map[string]interface{}{}, "-", tkCfg.TypeRestart, tkCfg.RegionGlobal)
}
