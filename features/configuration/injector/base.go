package injector

import (
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/handlers"
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/repositories"
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/services"
	"github.com/Smart-Door-Lock-IoT/api/pkg/database"
	"github.com/google/wire"
)

type ConfigurationHandlers struct {
	Pin *handlers.Pin
}

func NewConfigurationHandlers(
	pin *handlers.Pin,
) *ConfigurationHandlers {
	return &ConfigurationHandlers{
		Pin: pin,
	}
}

var (
	Set = wire.NewSet(
		handlers.NewPin,
		services.NewPin,
		repositories.NewConfig,
		database.New,
	)
)
