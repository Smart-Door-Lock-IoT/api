run:
	go run github.com/Smart-Door-Lock-IoT/api/cmd/api

di:
	wire gen github.com/Smart-Door-Lock-IoT/api/injector

docs:
	swag init -g cmd/api/main.go --requiredByDefault

.PHONY: run di docs