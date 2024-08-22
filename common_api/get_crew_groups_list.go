package common_api

type (
	GetCrewGroupsListResponse struct {
		// Список групп экипажей
		CrewGroups []CrewGroup `json:"crew_groups"`
	}

	CrewGroup struct {
		// ИД группы экипажей
		Id int `json:"id"`
		// Название группы экипажей
		Name string `json:"name"`
	}
)

// Запрос списка групп экипажей
func (cl *Client) GetCrewGroupsList() (response GetCrewGroupsListResponse, err error) {
	err = cl.Get("get_crew_groups_list", nil, nil, &response)
	return
}
