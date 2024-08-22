package common_api

type (
	GetDiscountsListResponse struct {
		// Список скидок
		Discounts []Discount `json:"discounts"`
	}

	Discount struct {
		// ИД скидки
		Id int `json:"id"`
		// Название скидки
		Name string `json:"name"`
		// Абсолютная сумма скидки, руб
		Sum float64 `json:"sum"`
		// Процент скидки от стоимости заказа, %
		Percent float64 `json:"percent"`
	}
)

// Запрос списка скидок
func (cl *Client) GetDiscountsList() (response GetDiscountsListResponse, err error) {
	err = cl.Get("get_discounts_list", nil, nil, &response)
	return
}
