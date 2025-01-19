package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	AnalyzePhoneRequest struct {
		// Номер телефона
		Phone string `validate:"required"`

		// Искать среди телефонов водителей
		SearchInDriversMobile *bool `validate:"omitempty"`
		// Искать среди телефонов клиентов
		SearchInClients *bool `validate:"omitempty"`
		// Искать в справочнике телефонов
		SearchInPhones *bool `validate:"omitempty"`
	}

	AnalyzePhoneResponse struct {
		// Может принимать значения: "driver_mobile", "client", "phone"
		PhoneType string `json:"phone_type"`
		// ИД водителя, клиента, телефона из справочника
		ID int `json:"id"`
		// ИД сотрудника клиента (если телефон найден среди телефонов клиента)
		ClientEmployeeID int `json:"client_employee_id"`
	}
)

// Анализ телефона
func (cl *Client) AnalyzePhone(req AnalyzePhoneRequest) (response AnalyzePhoneResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("phone", req.Phone)
	if req.SearchInDriversMobile != nil {
		v.Add("search_in_drivers_mobile", strconv.FormatBool(*req.SearchInDriversMobile))
	}
	if req.SearchInClients != nil {
		v.Add("search_in_clients", strconv.FormatBool(*req.SearchInClients))
	}
	if req.SearchInPhones != nil {
		v.Add("search_in_phones", strconv.FormatBool(*req.SearchInPhones))
	}

	/*
		100 Телефон не найден
	*/
	e := errorMap{
		100: ErrPhoneNotFound,
	}

	err = cl.Get("analyze_phone", e, v, &response)

	return
}
