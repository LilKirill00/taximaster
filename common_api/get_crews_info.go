package common_api

import (
	"net/url"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetCrewsInfoRequest struct {
		// Нужно ли возвращать экипажи не на линии
		// По умолчанию возвращаются только экипажи на линии
		NotWorkingCrews bool `validate:"omitempty"`
		// Список возвращаемых полей через запятую
		Fields string `validate:"fields"`
	}

	GetCrewsInfoResponse struct {
		// Массив экипажей
		CrewsInfo []GetCrewInfoResponse `json:"crews_info"`
	}
)

// Запрос информации об экипажах
func (cl *Client) GetCrewsInfo(req GetCrewsInfoRequest) (response GetCrewsInfoResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return response, err
	}

	v := url.Values{}
	if req.NotWorkingCrews {
		v.Add("not_working_crews", "true")
	}
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	err = cl.Get("get_crews_info", nil, v, &response)

	return response, err
}
