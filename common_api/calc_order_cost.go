package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	CalcOrderCostRequest struct {
		// ИД тарифа
		TariffID int `json:"tariff_id" validate:"required"`

		// Время подачи
		SourceTime string `json:"source_time,omitempty" validate:"omitempty,datetime=20060102150405"`
		// Предварительный заказ
		IsPrior bool `json:"is_prior,omitempty" validate:"omitempty"`
		// ИД клиента
		ClientId *int `json:"client_id,omitempty" validate:"omitempty"`
		// ИД сотрудника клиента
		ClientEmployeeId *int `json:"client_employee_id,omitempty" validate:"omitempty"`
		// ИД скидки
		DiscountId *int `json:"discount_id,omitempty" validate:"omitempty"`
		// ИД дисконтной карты
		DiscCardId *int `json:"disc_card_id,omitempty" validate:"omitempty"`
		// ИД района подачи
		SourceZoneId *int `json:"source_zone_id,omitempty" validate:"omitempty"`
		// ИД района назначения
		DestZoneId *int `json:"dest_zone_id,omitempty" validate:"omitempty"`
		// Километраж по городу
		DistanceCity *float64 `json:"distance_city,omitempty" validate:"omitempty"`
		// Километраж за городом
		DistanceCountry *float64 `json:"distance_country,omitempty" validate:"omitempty"`
		// Километраж до подачи за городом
		SourceDistanceCountry *float64 `json:"source_distance_country,omitempty" validate:"omitempty"`
		// Загородный заказ
		IsCountry bool `json:"is_country,omitempty" validate:"omitempty"`
		// Время ожидания посадки клиента в минутах
		WaitingMinutes *int `json:"waiting_minutes,omitempty" validate:"omitempty"`
		// Почасовой заказ
		IsHourly bool `json:"is_hourly,omitempty" validate:"omitempty"`
		// Длительность почасового заказа в минутах
		HourlyMinutes *int `json:"hourly_minutes,omitempty" validate:"omitempty"`
		// Призовой заказ
		IsPrize bool `json:"is_prize,omitempty" validate:"omitempty"`
		// Обратный путь за городом
		BackWay bool `json:"back_way,omitempty" validate:"omitempty"`
		// Список ИД услуг через точку с запятой, пример: «1;2;3» Устарело. Рекомендуется использовать параметр order_params.
		Services string `json:"services,omitempty" validate:"omitempty"`
		// Список ИД параметров заказа через точку с запятой, пример: «1;2;3»
		OrderParams string `json:"order_params,omitempty" validate:"omitempty"`
		// Признак безналичного заказа
		Cashless bool `json:"cashless,omitempty" validate:"omitempty"`
	}

	CalcOrderCostResponse struct {
		// Рассчитанная общая сумма заказа
		Sum float64 `json:"sum"`
		// Дополнительная информация по расчету суммы заказа
		Info []struct {
			// Описание позиции дополнительной информации по расчету суммы заказа
			Comment string `json:"comment"`
			// Сумма позиции дополнительной информации по расчету суммы заказа
			Sum string `json:"sum"`
		} `json:"info"`
	}
)

// Расчет суммы заказа
func (cl *Client) CalcOrderCost(req CalcOrderCostRequest) (response CalcOrderCostResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	v.Add("tariff_id", strconv.Itoa(req.TariffID))
	if req.SourceTime != "" {
		v.Add("source_time", req.SourceTime)
	}
	if req.IsPrior {
		v.Add("is_prior", "true")
	}
	if req.SourceTime != "" {
		v.Add("source_time", req.SourceTime)
	}
	if req.ClientId != nil {
		v.Add("client_id", strconv.Itoa(*req.ClientId))
	}
	if req.ClientEmployeeId != nil {
		v.Add("client_employee_id", strconv.Itoa(*req.ClientEmployeeId))
	}
	if req.DiscountId != nil {
		v.Add("discount_id", strconv.Itoa(*req.DiscountId))
	}
	if req.DiscCardId != nil {
		v.Add("disc_card_id", strconv.Itoa(*req.DiscCardId))
	}
	if req.SourceZoneId != nil {
		v.Add("source_zone_id", strconv.Itoa(*req.SourceZoneId))
	}
	if req.DestZoneId != nil {
		v.Add("dest_zone_id", strconv.Itoa(*req.DestZoneId))
	}
	if req.DistanceCity != nil {
		v.Add("distance_city", strconv.FormatFloat(*req.DistanceCity, 'g', -1, 64))
	}
	if req.DistanceCountry != nil {
		v.Add("distance_country", strconv.FormatFloat(*req.DistanceCountry, 'g', -1, 64))
	}
	if req.SourceDistanceCountry != nil {
		v.Add("source_distance_country", strconv.FormatFloat(*req.SourceDistanceCountry, 'g', -1, 64))
	}
	if req.IsCountry {
		v.Add("is_country", "true")
	}
	if req.WaitingMinutes != nil {
		v.Add("waiting_minutes", strconv.Itoa(*req.WaitingMinutes))
	}
	if req.IsHourly {
		v.Add("is_hourly", "true")
	}
	if req.HourlyMinutes != nil {
		v.Add("hourly_minutes", strconv.Itoa(*req.HourlyMinutes))
	}
	if req.IsPrize {
		v.Add("is_prize", "true")
	}
	if req.BackWay {
		v.Add("back_way", "true")
	}
	if req.Services != "" {
		v.Add("services", req.Services)
	}
	if req.OrderParams != "" {
		v.Add("order_params", req.OrderParams)
	}
	if req.Cashless {
		v.Add("cashless", "true")
	}

	/*
		100	Тариф не найден
		101	Ошибка при расчете по тарифу
		102	Скидка не найдена
		103	Клиент не найден
		104	Район подачи не найден
		105	Район назначения не найден
		106	Дисконтная карта не найдена
		107	Район остановки не найден
		108	Группа экипажа не найдена
		109	Служба ЕДС не найдена
		110	Дисконтная карта не действительна
		111	Не найден сотрудник клиента
	*/
	e := errorMap{
		100: ErrTariffNotFound,
		101: ErrCalculationByTariff,
		102: ErrDiscountNotFound,
		103: ErrClientNotFound,
		104: ErrZoneSourceNotFound,
		105: ErrZoneDestinationNotFound,
		106: ErrDiscountCardNotFound,
		107: ErrZoneStopNotFound,
		108: ErrCrewNotFound,
		109: ErrUdsNotFound,
		110: ErrDiscountCardIsNotValid,
		111: ErrCustomerClientNotFound,
	}

	err = cl.Get("calc_order_cost", e, v, &response)

	return
}
