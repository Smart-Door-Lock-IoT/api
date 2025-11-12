package mqtt

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func messageHandler(c mqtt.Client, m mqtt.Message) {
	fmt.Println("Received message: ", string(m.Payload()), " from topic: ", m.Topic())
}

func New() {
	broker := os.Getenv("MQTT_BROKER")
	clientID := "smart-door-api-client-" + time.Now().Format("20060102-150405")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.ResumeSubs = true
	// opts.SetDefaultPublishHandler(messageHandler)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.WaitTimeout(time.Second*5) && token.Error() != nil {
		panic("failed to connect to MQTT broker: " + token.Error().Error())
	}

	if !client.IsConnected() {
		panic("failed to connect to MQTT broker")
	}

	if token := client.Subscribe(
		"smart-door-lock-iot/#", 0, messageHandler,
	); token.WaitTimeout(time.Second*5) && token.Error() != nil {
		panic("failed to subscribe to topic: " + token.Error().Error())
	}

	for {
		time.Sleep(time.Second * 5)
		if !client.IsConnected() {
			fmt.Println("MQTT client disconnected, attempting to reconnect...")
		}
	}
}
