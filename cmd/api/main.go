package main

import (
	"os"

	"github.com/Smart-Door-Lock-IoT/api/pkg/http"
	"github.com/Smart-Door-Lock-IoT/api/pkg/mqtt"
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
		mqtt.New()
	}()

	http.NewServer()
}
