package user

import (
	"community-demo/model/user"
	"community-demo/utils"
	"errors"
)

type loginFlow struct {
	Name     string
	Password string
	User     *user.User
	Token    string
}

func LoginService(name, password string) (string, error) {
	return LoginFlow(name, password).do()
}

func LoginFlow(name, password string) *loginFlow {
	return &loginFlow{
		Name:     name,
		Password: password,
	}
}

func (f *loginFlow) do() (string, error) {
	if err := f.checkMatch(); err != nil {
		return "", err
	}
	return f.Token, nil
}

func (f *loginFlow) checkMatch() error {
	name := f.Name
	password := utils.Encode(f.Password)
	var nUser *user.User
	var err error
	if nUser, err = user.NewUserDaoInstance().CheckMatch(name, password); err != nil {
		return err
	}
	if nUser == nil {
		return errors.New("用户名或密码错误")
	}
	token, _ := utils.ParseJWT(nUser.ID, nUser.Name)
	f.Token = token
	return nil
}
