package mqtt

import (
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func messageHandler(c mqtt.Client, m mqtt.Message) {

}

func New() {
	broker := os.Getenv("MQTT_BROKER")
	clientID := "smart-door-api-client-" + time.Now().Format("20060102-150405")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	// opts.SetDefaultPublishHandler(messageHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic("failed to connect to MQTT broker: " + token.Error().Error())
	}

	if token := client.Subscribe(
		"smart-door-lock-iot/#", 2, messageHandler,
	); token.Wait() && token.Error() != nil {
		panic("failed to subscribe to topic: " + token.Error().Error())
	}
}
