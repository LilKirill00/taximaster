package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateWayBillCheckRequest struct {
		// ИД путевого листа (должен быть задан либо ИД либо номер)
		WayBillId int `json:"way_bill_id" validate:"required"`
		// Номер путевого листа (должен быть задан либо ИД либо номер)
		WayBillNumber string `json:"way_bill_number" validate:"required"`
		// Тип осмотра ("med/tech")
		Kind string `json:"kind" validate:"required,eq=med|eq=tech"`
		// Имя пользователя
		UserName string `json:"user_name" validate:"required"`
		// Результат осмотра
		Success bool `json:"success" validate:"required"`

		// Номер осмотра
		Number string `json:"number,omitempty" validate:"omitempty"`
		// Комментарий
		Comment string `json:"comment,omitempty" validate:"omitempty"`
	}
)

// Создание осмотра по путевому листу
func (cl *Client) CreateWayBillCheck(req CreateWayBillCheckRequest) (response EmptyResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100 Нет лицензии на использование путевых листов
		101 Не найден путевой лист
	*/
	e := errorMap{
		100: ErrNoLicenseToUseWayBill,
		101: ErrWayBillNotFound,
	}

	err = cl.PostJson("create_way_bill_check", e, req, &response)

	return
}
