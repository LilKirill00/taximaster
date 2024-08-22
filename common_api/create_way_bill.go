package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateWayBillRequest struct {
		// Время начала
		StartTime string `json:"start_time" validate:"required, datetime=20060102150405"`
		// Время завершения
		FinishTime string `json:"finish_time" validate:"required, datetime=20060102150405"`
		// ИД водителя
		DriverId int `json:"driver_id" validate:"required"`
		// ИД автомобиля
		CarId int `json:"car_id" validate:"required"`

		// Номер путевого листа
		Number string `json:"number,omitempty" validate:"omitempty"`
		// Комментарий
		Comment string `json:"comment,omitempty" validate:"omitempty"`
	}

	CreateWayBillResponse struct {
		// ИД созданного путевого листа
		WayBillId int `json:"way_bill_id"`
	}
)

// Создание путевого листа
func (cl *Client) CreateWayBill(req CreateWayBillRequest) (response CreateWayBillResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100 Нет лицензии на использование путевых листов
		101 Не найден водитель
		102 Не найден автомобиль
	*/
	e := errorMap{
		100: ErrNoLicenseToUseWayBill,
		101: ErrDriverNotFound,
		102: ErrCarNotFound,
	}

	err = cl.PostJson("create_way_bill", e, req, &response)

	return
}
