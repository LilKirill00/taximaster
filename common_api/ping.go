package common_api

// Запрос-пинг
func (cl *Client) Ping() (response EmptyResponse, err error) {
	err = cl.PostJson("ping", nil, nil, &response)
	return
}
