package common_api

import (
	"net/url"
	"strconv"

	"github.com/ros-tel/taximaster/validator"
)

type (
	GetClientsInfoRequest struct {
		// Текст для поиска по названию или по номеру договора клиента
		Text string `validate:"omitempty"`
		// Максимальное количество клиентов, которое надо вернуть. Если не указано, то 10.
		MaxClientsCount int `validate:"omitempty"`
		// Фильтр по группе клиентов
		ClientGroupId int `validate:"omitempty"`
		// Фильтр по вышестоящему подразделению, возвращаются все подчиненные отделы и сотрудники на всю глубину иерархии
		ParentID int `validate:"omitempty"`
		// Список возвращаемых полей через запятую. По умолчанию возвращаются поля "name" и "number". Поле "client_id" возвращается всегда
		Fields string `validate:"omitempty"`
	}

	GetClientsInfoResponse struct {
		ClientsInfo []GetClientInfoResponse `json:"clients_info"`
	}
)

// Запрос информации по клиенту
func (cl *Client) GetClientsInfo(req GetClientsInfoRequest) (response GetClientsInfoResponse, err error) {
	err = validator.Validate(req)
	if err != nil {
		return
	}

	v := url.Values{}
	if req.Text != "" {
		v.Add("fields", req.Text)
	}
	if req.MaxClientsCount != 0 {
		v.Add("fields", strconv.Itoa(req.MaxClientsCount))
	}
	if req.ClientGroupId != 0 {
		v.Add("fields", strconv.Itoa(req.ClientGroupId))
	}
	if req.ParentID != 0 {
		v.Add("fields", strconv.Itoa(req.ParentID))
	}
	if req.Fields != "" {
		v.Add("fields", req.Fields)
	}

	err = cl.Get("get_client_info", nil, v, &response)

	return
}
