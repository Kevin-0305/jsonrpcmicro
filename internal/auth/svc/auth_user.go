package svc

type UserService struct {
}

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (that *UserService) Login(loginRequest LoginRequest, reply *UserResponse) error {
	if loginRequest.Account == "Admin" && loginRequest.Password == "123456" {
		reply.ID = 1
		reply.Name = "管理员"
	} else {
		reply.ID = 0
		reply.Name = "该账户不存在"
	}
	return nil
}
