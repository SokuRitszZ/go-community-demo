package topic

import (
	"community-demo/model/topic"
	"errors"
	"strings"
)

type publishFlow struct {
	Title   string
	UserID  uint
	Content string
	Topic   *topic.Topic
}

func PublishService(title string, userID uint, content string) (*topic.Topic, error) {
	return PublishFlow(title, userID, content).do()
}

func PublishFlow(title string, userID uint, content string) *publishFlow {
	return &publishFlow{
		title,
		userID,
		content,
		nil,
	}
}

func (f *publishFlow) do() (*topic.Topic, error) {
	if err := f.checkTitle(); err != nil {
		return nil, err
	}
	if err := f.checkContent(); err != nil {
		return nil, err
	}
	if err := f.publish(); err != nil {
		return nil, err
	}
	return f.Topic, nil
}

func (f *publishFlow) checkTitle() error {
	title := strings.Trim(f.Title, " ")
	if len(title) == 0 {
		return errors.New("标题为空")
	}
	if len(title) > 16 {
		return errors.New("标题太长，长度超过16")
	}
	return nil
}

func (f *publishFlow) checkContent() error {
	content := strings.Trim(f.Content, " ")
	if len(content) == 0 {
		return errors.New("内容为空")
	}
	if len(content) > 256 {
		return errors.New("内容太长，长度超过256")
	}
	return nil
}

func (f *publishFlow) publish() error {
	nTopic := topic.Topic{
		Title:   strings.Trim(f.Title, " "),
		UserID:  f.UserID,
		Content: strings.Trim(f.Content, " "),
	}
	if err := topic.Dao().Publish(&nTopic); err != nil {
		return err
	}
	f.Topic = &nTopic
	return nil
}
