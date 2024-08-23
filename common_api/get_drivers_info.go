package common_api

import (
	"net/url"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetDriversInfoRequest struct {
		// Включить в ответ запроса заблокированных водителей
		LockedDrivers bool `validate:"omitempty"`
		// Включить в ответ запроса уволенных водителей
		DismissedDrivers bool `validate:"omitempty"`
		// Список возвращаемых полей через запятую
		Fields string `validate:"omitempty"`
	}

	GetDriversInfoResponse struct {
		// Массив не удаленных водителей
		DriversInfo []GetDriverInfoResponse `json:"drivers_info"`
	}
)

// Запрос списка водителей
func (cl *Client) GetDriversInfo(req GetDriversInfoRequest) (response GetDriversInfoResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	if req.LockedDrivers {
		v.Add("locked_drivers", "true")
	}
	if req.DismissedDrivers {
		v.Add("dismissed_drivers", "true")
	}
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	err = cl.Get("get_drivers_info", nil, v, &response)

	return
}
