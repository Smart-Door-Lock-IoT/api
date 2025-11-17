package services

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Control struct {
	mqttClient *mqtt.Client
}

func NewControl(mqttClient *mqtt.Client) *Control {
	return &Control{
		mqttClient: mqttClient,
	}
}

func (s *Control) TriggerOpenDoor()        {}
func (s *Control) TriggerFingerprintMode() {}
func (s *Control) TriggerRFIDMode()        {}
