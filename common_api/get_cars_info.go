package common_api

import (
	"net/url"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetCarsInfoRequest struct {
		// Включить в ответ заблокированных автомобилей (по умолчанию false)
		LockedCars bool `validate:"omitempty"`
		// Список возвращаемых полей через запятую
		Fields string `validate:"fields"`
	}

	GetCarsInfoResponse struct {
		// Массив автомобилей
		CrewsInfo []GetCarInfoRequest `json:"crews_info"`
	}
)

// Запрос информации об автомобиле
func (cl *Client) GetCarsInfo(req GetCarsInfoRequest) (response GetCarsInfoResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	if req.LockedCars {
		v.Add("locked_cars", "true")
	}
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	/*
		100 Автомобиль не найден
	*/
	e := errorMap{
		100: ErrCarNotFound,
	}

	err = cl.Get("get_car_info", e, v, &response)

	return
}
