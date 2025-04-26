package validation

type LoginInfo struct {
	Account      string `json:"account"`
	AccountCheck string `json:"account_check"`
}

func LoginMethod(loginMethod uint8, userEmail, phone, emailCode, userPassword string) (LoginInfo, bool) {
	switch loginMethod {
	case 1: //邮箱 密码
		return LoginInfo{
			Account:      userEmail,
			AccountCheck: userPassword,
		}, true
	case 2: //邮箱 验证码
		return LoginInfo{
			Account:      userEmail,
			AccountCheck: emailCode,
		}, true
	default:
		return LoginInfo{}, false
	}
}
