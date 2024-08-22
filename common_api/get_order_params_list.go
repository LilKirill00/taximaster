package common_api

type (
	GetOrderParamsListResponse struct {
		// Список параметров заказа
		OrderParams []OrderParam `json:"order_params"`
	}

	OrderParam struct {
		// ИД параметра
		Id int `json:"id"`
		// Название параметра
		Name string `json:"name"`
		// Абсолютная сумма параметра, руб
		Sum float64 `json:"sum"`
		// Процент параметра от стоимости заказа, %
		Percent float64 `json:"percent"`
		// Регулирует доступ к заказу
		OrderAccessControl bool `json:"order_access_control"`
	}
)

// Запрос списка атрибутов
func (cl *Client) GetOrderParamsList() (response GetOrderParamsListResponse, err error) {
	err = cl.Get("get_order_params_list", nil, nil, &response)
	return
}
