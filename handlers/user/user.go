package user

import (
	"community-demo/service/user"
	"community-demo/utils"
)

func RegisterHandler(name string, password string) *utils.Message {
	user, err := user.RegisterService(name, password)
	if err != nil {
		return &utils.Message{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &utils.Message{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
}

func LoginHandler(name string, password string) *utils.Message {
	token, err := user.LoginService(name, password)
	if err != nil {
		return &utils.Message{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &utils.Message{
		Code: 0,
		Msg:  "OK",
		Data: token,
	}
}

func GetInfoHandler(id uint) *utils.Message {
	info, err := user.GetInfoService(id)
	if err != nil {
		return &utils.Message{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &utils.Message{
		Code: 0,
		Msg:  "OK",
		Data: info,
	}
}
