package requests

type CreateLogRequest struct {
	DeviceName string `json:"device_name"`
	Status     bool   `json:"status"`
} // @name CreateLogRequest
