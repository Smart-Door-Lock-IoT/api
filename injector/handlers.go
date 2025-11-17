//go:build wireinject
// +build wireinject

package injector

import (
	configuration "github.com/Smart-Door-Lock-IoT/api/features/configuration/injector"
	control "github.com/Smart-Door-Lock-IoT/api/features/control/injector"
	"github.com/google/wire"
)

func InitConfigurationHandlers() *configuration.ConfigurationHandlers {
	wire.Build(
		configuration.NewConfigurationHandlers,
		configuration.Set,
	)
	return nil
}

func InitControlHandlers() *control.ControlHandlers {
	wire.Build(
		control.NewControlHandlers,
		control.Set,
	)
	return nil
}
