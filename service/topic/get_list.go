package topic

import (
	"community-demo/model/topic"
	"community-demo/model/user"
	"errors"
)

type getListFlow struct {
	UserID uint
	List   []topic.Topic
}

func GetListService(userID uint) ([]topic.Topic, error) {
	return GetListFlow(userID).do()
}

func GetListFlow(userID uint) *getListFlow {
	return &getListFlow{
		userID,
		nil,
	}
}

func (f *getListFlow) do() ([]topic.Topic, error) {
	if err := f.checkUserExist(); err != nil {
		return nil, err
	}
	if err := f.getList(); err != nil {
		return nil, err
	}
	return f.List, nil
}

func (f *getListFlow) checkUserExist() error {
	info, err := user.Dao().GetInfoById(int64(f.UserID))
	if err != nil {
		return err
	}
	if info.ID == 0 {
		return errors.New("不存在此用户")
	}
	return nil
}

func (f *getListFlow) getList() error {
	id := f.UserID
	topics, err := topic.Dao().GetByUserId(id)
	if err != nil {
		return err
	}
	if topics == nil || len(topics) == 0 {
		return errors.New("该用户没有发过帖子")
	}
	f.List = topics
	return nil
}
