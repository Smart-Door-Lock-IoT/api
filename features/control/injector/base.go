package injector

import (
	"github.com/Smart-Door-Lock-IoT/api/features/control/handlers"
	"github.com/Smart-Door-Lock-IoT/api/features/control/repositories"
	"github.com/Smart-Door-Lock-IoT/api/features/control/services"
	"github.com/Smart-Door-Lock-IoT/api/pkg/database"
	"github.com/Smart-Door-Lock-IoT/api/pkg/mqttclient"
	"github.com/google/wire"
)

type ControlHandlers struct {
	Control *handlers.Control
	Log     *handlers.Log
}

func NewControlHandlers(
	control *handlers.Control,
	log *handlers.Log,
) *ControlHandlers {
	return &ControlHandlers{
		Control: control,
		Log:     log,
	}
}

var (
	Set = wire.NewSet(
		handlers.NewControl,
		handlers.NewLog,
		services.NewControl,
		services.NewLog,
		repositories.NewLog,
		database.New,
		mqttclient.New,
	)
)
