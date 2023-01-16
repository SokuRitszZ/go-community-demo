package topic

import (
	"community-demo/model/topic"
	"errors"
)

type getDetailFlow struct {
	ID    uint
	Topic *topic.Topic
}

func GetDetailService(id uint) (*topic.Topic, error) {
	return GetDetailFlow(id).do()
}

func GetDetailFlow(id uint) *getDetailFlow {
	return &getDetailFlow{
		id,
		nil,
	}
}

func (f *getDetailFlow) do() (*topic.Topic, error) {
	if err := f.checkID(); err != nil {
		return nil, err
	}
	if err := f.getDetail(); err != nil {
		return nil, err
	}
	return f.Topic, nil
}

func (f *getDetailFlow) checkID() error {
	id := f.ID
	if id == 0 {
		return errors.New("编号不合法")
	}
	return nil
}

func (f *getDetailFlow) getDetail() error {
	nTopic, err := topic.Dao().GetById(f.ID)
	if err != nil {
		return err
	}
	if nTopic == nil {
		return errors.New("不存在此帖子")
	}
	f.Topic = nTopic
	return nil
}
