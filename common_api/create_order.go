package common_api

import "github.com/ros-tel/taximaster/validator"

type (
	CreateOrderRequest struct {
		// Номер телефона
		Phone string `json:"phone" validate:"required,max=30"`
		// Адрес подачи
		Source string `json:"source" validate:"required"`
		// Время подачи
		SourceTime string `json:"source_time" validate:"required,datetime=20060102150405"`

		// Адрес назначения
		Dest string `json:"dest,omitempty" validate:"omitempty"`
		// Заказчик
		Customer string `json:"customer,omitempty" validate:"omitempty"`
		// Комментарий
		Comment string `json:"comment,omitempty" validate:"omitempty"`
		// ИД группы экипажей
		CrewGroupID int `json:"crew_group_id,omitempty" validate:"omitempty"`
		// ИД службы ЕДС
		UdsID int `json:"uds_id,omitempty" validate:"omitempty"`
		// ИД тарифа
		TariffID int `json:"tariff_id,omitempty" validate:"omitempty"`
		// Предварительный заказ
		IsPrior bool `json:"is_prior,omitempty" validate:"omitempty"`
		// Долгота адреса подачи
		SourceLon float64 `json:"source_lon,omitempty" validate:"omitempty"`
		// Широта адреса подачи
		SourceLat float64 `json:"source_lat,omitempty" validate:"omitempty"`
		// Долгота адреса назначения
		DestLon float64 `json:"dest_lon,omitempty" validate:"omitempty"`
		// Широта адреса назначения
		DestLat float64 `json:"dest_lat,omitempty" validate:"omitempty"`
	}

	CreateOrderResponse struct {
		// ИД созданного заказа
		OrderID int `json:"order_id"`
	}
)

// Создание нового заказа
func (cl *Client) CreateOrder(req CreateOrderRequest) (response CreateOrderResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	/*
		100	Заказ с такими параметрами уже создан
		101	Тариф не найден
		102	Группа экипажа не найдена
		103	Служба ЕДС не найдена
		110	Клиент заблокирован
		111	Не найден клиент, который может использовать собственный счет для оплаты заказов
		114	Недостаточно средств на безналичном счете клиента в ТМ
		115	Отрицательный баланс на безналичном счете клиента в ТМ
		116	Для клиента запрещена оплата заказа наличными. Клиент должен максимально использовать в заказе безналичную оплату (оплату с основного счета)
	*/

	e := errorMap{
		100: ErrOrderExistsWithParametrs,
		101: ErrTariffNotFound,
		102: ErrCrewGroupsNotFound,
		103: ErrUdsNotFound,
		110: ErrClientBlocked,
		111: ErrClientwhoCanUseTheirOwnNotFound,
		114: ErrInsufficientFundsCashless,
		115: ErrNegativeBalanceCashless,
		116: ErrCashPaymentNotAllowed,
	}

	err = cl.PostJson("create_order", e, req, &response)

	return
}
