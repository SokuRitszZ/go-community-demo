package post

import (
	"community-demo/model/post"
)

type deleteFlow struct {
	ID     uint
	UserID uint
}

func DeleteService(userID, id uint) error {
	return DeleteFlow(userID, id).do()
}

func DeleteFlow(userID, id uint) *deleteFlow {
	return &deleteFlow{
		id,
		userID,
	}
}

func (f *deleteFlow) do() error {
	if err := post.Dao().DeleteById(f.UserID, f.ID); err != nil {
		return err
	}
	return nil
}
