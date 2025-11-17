package services

import (
	"time"

	"github.com/Smart-Door-Lock-IoT/api/features/control/dto/responses"
	httputils "github.com/Smart-Door-Lock-IoT/api/pkg/http/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Control struct {
	mqttClient *mqtt.Client
}

func NewControl(mqttClient *mqtt.Client) *Control {
	return &Control{
		mqttClient: mqttClient,
	}
}

func (s *Control) TriggerOpenDoor() (*responses.TriggerOpenDoorResponse, *httputils.Error) {
	if token := (*s.mqttClient).Publish(
		"smart-door-lock-iot/open-door",
		2,
		false,
		"triggered",
	); token.WaitTimeout(time.Second*5) && token.Error() != nil {
		return nil, httputils.NewInternalError(token.Error())
	} else {
		return &responses.TriggerOpenDoorResponse{
			Message: "oke",
		}, nil
	}
}

func (s *Control) TriggerFingerprintMode() (
	*responses.TriggerFingerprintModeResponse, *httputils.Error,
) {
	if token := (*s.mqttClient).Publish(
		"smart-door-lock-iot/fingerprint-mode",
		2,
		false,
		"triggered",
	); token.WaitTimeout(time.Second*5) && token.Error() != nil {
		return nil, httputils.NewInternalError(token.Error())
	} else {
		return &responses.TriggerFingerprintModeResponse{
			Message: "oke",
		}, nil
	}
}

func (s *Control) TriggerRFIDMode() (*responses.TriggerRFIDModeResponse, *httputils.Error) {
	if token := (*s.mqttClient).Publish(
		"smart-door-lock-iot/rfid-mode",
		2,
		false,
		"triggered",
	); token.WaitTimeout(time.Second*5) && token.Error() != nil {
		return nil, httputils.NewInternalError(token.Error())
	} else {
		return &responses.TriggerRFIDModeResponse{
			Message: "oke",
		}, nil
	}
}
