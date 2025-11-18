package services

import (
	"fmt"

	"github.com/Smart-Door-Lock-IoT/api/features/control/domains"
	"github.com/Smart-Door-Lock-IoT/api/features/control/dto/responses"
	"github.com/Smart-Door-Lock-IoT/api/features/control/repositories"
	httputils "github.com/Smart-Door-Lock-IoT/api/pkg/http/utils"
)

type Log struct {
	logRepo *repositories.Log
}

func NewLog(logRepo *repositories.Log) *Log {
	return &Log{
		logRepo: logRepo,
	}
}

func (s *Log) CreateLog(data domains.Log) {
	if res, err := s.logRepo.Create(data); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Log created:", res)
	}
}

func (s *Log) GetAllLogs() (*responses.GetAllLogsResponse, *httputils.Error) {
	if res, err := s.logRepo.GetAll(); err != nil {
		return nil, httputils.NewInternalError(err)
	} else {
		return &responses.GetAllLogsResponse{
			Logs: *res,
		}, nil
	}
}

func (s *Log) DeleteAllLogs() (*responses.DeleteAllLogsResponse, *httputils.Error) {
	if err := s.logRepo.DeleteAll(); err != nil {
		return nil, httputils.NewInternalError(err)
	} else {
		return &responses.DeleteAllLogsResponse{
			Message: "ok",
		}, nil
	}
}
