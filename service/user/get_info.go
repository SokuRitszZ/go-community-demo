package user

import "community-demo/model/user"

type getInfoFlow struct {
	ID   uint
	User *user.User
}

func GetInfoService(id uint) (*user.User, error) {
	return GetInfoFlow(id).do()
}

func GetInfoFlow(id uint) *getInfoFlow {
	return &getInfoFlow{
		ID: id,
	}
}

func (f *getInfoFlow) do() (*user.User, error) {
	if err := f.getBaseData(); err != nil {
		return nil, err
	}
	return f.User, nil
}

func (f *getInfoFlow) getBaseData() error {
	id := f.ID
	data, err := user.NewUserDaoInstance().GetBaseData(id)
	if err != nil {
		return err
	}
	data.Password = ""
	f.User = data
	return nil
}
