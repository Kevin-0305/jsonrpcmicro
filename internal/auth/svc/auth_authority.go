package svc

func CheckAuthority() bool {
	obj := "AuthUser"
	// 获取请求方法
	act := "Post"
	// 获取用户的角色
	sub := 1
	e := Casbin()
	// 判断策略中是否存在
	success, _ := e.Enforce(sub, obj, act)
	if success {
		return true
	} else {
		return false
	}
}
