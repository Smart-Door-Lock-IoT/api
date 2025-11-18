package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/Smart-Door-Lock-IoT/api/features/control/dto/requests"
	"github.com/Smart-Door-Lock-IoT/api/injector"
	"github.com/Smart-Door-Lock-IoT/api/pkg/http"
	"github.com/Smart-Door-Lock-IoT/api/pkg/mqttclient"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

// @title 		API Smart Door Lock IoT
// @version 	1.0
func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	if appEnv == "development" {
		if err := godotenv.Load(); err != nil {
			panic("gagal memuat file .env: " + err.Error())
		}
	}

	go func() {
		client := mqttclient.New()

		handlers := injector.InitControlHandlers()

		// subscribe log activities
		if token := (*client).Subscribe(
			"smart-door-lock-iot/log-activities",
			2,
			func(c mqtt.Client, m mqtt.Message) {
				var logReq requests.CreateLogRequest
				if err := json.Unmarshal(m.Payload(), &logReq); err != nil {
					println("failed to unmarshal logReq activity: " + err.Error())
				} else {
					handlers.Log.CreateLog(logReq)
				}
			},
		); token.WaitTimeout(time.Second*5) && token.Error() != nil {
			panic("failed to subscribe to log activities topic: " + token.Error().Error())
		} else {
			println("Subscribed to log activities topic")
		}
	}()

	http.NewServer()
}
