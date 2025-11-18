package mqttclient

import (
	"fmt"
	"os"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	mqttInstance *mqtt.Client
	mqttOnce     sync.Once
)

func New() *mqtt.Client {
	mqttOnce.Do(
		func() {
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
			} else {
				fmt.Println("Connected to MQTT broker")
			}

			mqttInstance = &client
		},
	)

	return mqttInstance
}
