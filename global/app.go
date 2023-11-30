package global

import (
	"gofly/utils/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Config
	Log         *zap.Logger
}

var App = new(Application)
