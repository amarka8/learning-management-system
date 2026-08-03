// this package is meant to reduce the presence of global variables and package together everything necessary to
// test this application easily. the environment includes the logger and a config to be used by the server package
package environment

import (
	"github.com/amarka8/learning-management-system/config"
	logging "github.com/amarka8/learning-management-system/logger"
)

type Env struct {
	logging.Logger
	config.Config
}

func Environment(logger logging.Logger, cfg config.Config) *Env {
	if logger == nil {
		logger = logging.NullLogger{}
	}
	if cfg == nil {
		cfg = config.NullConfig{}
	}

	return &Env{Logger: logger, Config: cfg}
}

func Null() *Env {
	return Environment(nil, nil)
}
