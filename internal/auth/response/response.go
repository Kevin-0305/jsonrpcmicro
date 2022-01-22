package response

type UserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	State   int    `json:"state"`
	Message string `json:"message"`
}
