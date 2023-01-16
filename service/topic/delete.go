package topic

import (
	"community-demo/model/topic"
	"community-demo/model/user"
	"errors"
)

type deleteFlow struct {
	UserID uint
	ID     uint
}

func DeleteService(userID, id uint) error {
	return DeleteFlow(userID, id).do()
}

func DeleteFlow(userID, id uint) *deleteFlow {
	return &deleteFlow{
		userID,
		id,
	}
}

func (f *deleteFlow) do() error {
	if err := f.checkUserExist(); err != nil {
		return err
	}
	if err := f.delete(); err != nil {
		return err
	}
	return nil
}

func (f *deleteFlow) checkUserExist() error {
	info, err := user.Dao().GetInfoById(int64(f.UserID))
	if err != nil {
		return err
	}
	if info.ID == 0 {
		return errors.New("不存在此用户")
	}
	return nil
}

func (f *deleteFlow) delete() error {
	if err := topic.Dao().DeleteById(f.UserID, f.ID); err != nil {
		return err
	}
	return nil
}
