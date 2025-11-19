package responses

import "github.com/Smart-Door-Lock-IoT/api/features/control/domains"

type GetAllLogsResponse struct {
	Logs []domains.Log `json:"logs"`
} // @name GetAllLogsResponse

type GetAllLatestLogsResponse struct {
	Logs []domains.Log `json:"logs"`
} // @name GetAllLatestLogsResponse

type DeleteAllLogsResponse struct {
	Message string `json:"message"`
} // @name DeleteAllLogsResponse
