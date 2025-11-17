package injector

import (
	"github.com/Smart-Door-Lock-IoT/api/features/control/handlers"
	"github.com/Smart-Door-Lock-IoT/api/features/control/services"
	"github.com/Smart-Door-Lock-IoT/api/pkg/mqtt"
	"github.com/google/wire"
)

type ControlHandlers struct {
	Control *handlers.Control
}

func NewControlHandlers(
	control *handlers.Control,
) *ControlHandlers {
	return &ControlHandlers{
		Control: control,
	}
}

var (
	Set = wire.NewSet(
		handlers.NewControl,
		services.NewControl,
		mqtt.New,
	)
)
