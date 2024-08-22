package common_api

type (
	GetTariffsListResponse struct {
		// Список тарифов
		Tariffs []Tariff `json:"tariffs"`
	}

	Tariff struct {
		// ИД тарифа
		Id int `json:"id"`
		// Название тарифа
		Name string `json:"name"`
		// Активный тариф
		IsActive bool `json:"is_active"`
	}
)

// Запрос списка тарифов
func (cl *Client) GetTariffsList() (response GetTariffsListResponse, err error) {
	err = cl.Get("get_tariffs_list", nil, nil, &response)
	return
}
