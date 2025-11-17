package requests

type ValidatePinRequest struct {
	Pin string `json:"pin"`
} // @name ValidatePinRequest

type ChangePinRequest struct {
	NewPin string `json:"new_pin"`
} // @name ChangePinRequest
