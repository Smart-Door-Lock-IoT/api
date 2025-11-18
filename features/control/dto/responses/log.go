package responses

import "github.com/Smart-Door-Lock-IoT/api/features/control/domains"

type GetAllLogsResponse struct {
	Logs []domains.Log `json:"logs"`
} // @name GetAllLogsResponse

type DeleteAllLogsResponse struct {
	Message string `json:"message"`
} // @name DeleteAllLogsResponse
