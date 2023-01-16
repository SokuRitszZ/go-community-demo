package user

import (
	"community-demo/model/user"
	"community-demo/utils"
	"errors"
	"strings"
)

type registerFlow struct {
	Name     string
	Password string
	User     *user.User
}

func RegisterService(name, password string) (*user.User, error) {
	return RegisterFlow(name, password).do()
}

func RegisterFlow(name, password string) *registerFlow {
	return &registerFlow{
		Name:     name,
		Password: password,
	}
}

func (f *registerFlow) do() (*user.User, error) {
	if err := f.checkExist(); err != nil {
		return nil, err
	}
	if err := f.checkEmpty(); err != nil {
		return nil, err
	}
	if err := f.Register(); err != nil {
		return nil, err
	}
	return f.User, nil
}

func (f *registerFlow) checkExist() error {
	user, err := user.Dao().FindByName(f.Name)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("用户名已被使用")
	}
	return nil
}

func (f *registerFlow) checkEmpty() error {
	name := strings.Trim(f.Name, " ")
	password := strings.Trim(f.Password, " ")
	if len(name) == 0 {
		return errors.New("用户名为空")
	}
	if len(name) > 16 {
		return errors.New("用户名太长，超过16")
	}
	if len(password) == 0 {
		return errors.New("密码为空")
	}
	if len(password) > 32 {
		return errors.New("密码太长，超过32")
	}
	return nil
}

func (f *registerFlow) Register() error {
	nUser := user.User{
		Name:     f.Name,
		Password: utils.Encode(f.Password),
		Avatar:   "https://sdfsdf.dev/100x100.png",
	}
	err := user.Dao().Register(&nUser)
	if err != nil {
		return err
	}
	f.User = &nUser
	return nil
}
