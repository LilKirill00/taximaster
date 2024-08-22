package common_api

type (
	GetClientGroupsListResponse struct {
		// Список групп клиентов
		ClientGroups []ClientGroup `json:"client_groups"`
	}

	ClientGroup struct {
		// ИД группы клиентов
		Id int `json:"id"`
		// Название группы клиентов
		Name string `json:"name"`
	}
)

// Запрос списка групп клиентов
func (cl *Client) GetClientGroupsList() (response GetClientGroupsListResponse, err error) {
	err = cl.Get("get_client_groups_list", nil, nil, &response)
	return
}
