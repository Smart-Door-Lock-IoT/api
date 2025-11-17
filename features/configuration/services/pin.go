package services

import (
	"net/http"
	"strings"

	"github.com/Smart-Door-Lock-IoT/api/features/configuration/domains"
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/dto/requests"
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/dto/responses"
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/repositories"
	httputils "github.com/Smart-Door-Lock-IoT/api/pkg/http/utils"
)

type Pin struct {
	configRepo *repositories.Config
}

func NewPin(configRepo *repositories.Config) *Pin {
	return &Pin{
		configRepo: configRepo,
	}
}

func (s *Pin) ValidatePin(req requests.ValidatePinRequest) (
	*responses.ValidatePinResponse, *httputils.Error,
) {
	if config, err := s.configRepo.GetByKey("pin"); err != nil {
		return nil, httputils.NewInternalError(err)
	} else {
		if strings.Trim(config.Value, " ") != strings.Trim(req.Pin, " ") {
			return nil, httputils.NewError(
				http.StatusForbidden,
				"PIN yang Anda masukkan salah!",
				nil,
			)
		} else {
			return &responses.ValidatePinResponse{
				Message: "oke",
			}, nil
		}
	}
}

func (s *Pin) ChangePin(req requests.ChangePinRequest) (
	*responses.ChangePinResponse, *httputils.Error,
) {
	if _, err := s.configRepo.Upsert(
		domains.Config{
			Key:   "pin",
			Value: req.NewPin,
		},
	); err != nil {
		return nil, httputils.NewInternalError(err)
	} else {
		return &responses.ChangePinResponse{
			Message: "oke",
		}, nil
	}
}
