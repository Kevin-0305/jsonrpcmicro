package request

type Login struct {
	Accout   string `json:"accout"`
	Password string `json:"password"`
}
