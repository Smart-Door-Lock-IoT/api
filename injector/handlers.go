//go:build wireinject
// +build wireinject

package injector

import (
	configuration "github.com/Smart-Door-Lock-IoT/api/features/configuration/injector"
	"github.com/google/wire"
)

func InitConfigurationHandlers() *configuration.ConfigurationHandlers {
	wire.Build(
		configuration.NewConfigurationHandlers,
		configuration.Set,
	)
	return nil
}
