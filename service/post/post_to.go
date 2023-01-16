package post

import (
	"community-demo/model/post"
	"community-demo/model/topic"
	"community-demo/model/user"
	"errors"
	"strings"
)

type postToFlow struct {
	UserID  uint
	TopicID uint
	Content string
	Post    *post.Post
}

func ToService(userID, topicID uint, content string) (*post.Post, error) {
	return ToFlow(userID, topicID, content).do()
}

func ToFlow(userID, topicID uint, content string) *postToFlow {
	return &postToFlow{
		userID,
		topicID,
		content,
		nil,
	}
}

func (f *postToFlow) do() (*post.Post, error) {
	if err := f.checkTopicExist(); err != nil {
		return nil, err
	}
	if err := f.checkUserExist(); err != nil {
		return nil, err
	}
	if err := f.checkContent(); err != nil {
		return nil, err
	}
	if err := f.to(); err != nil {
		return nil, err
	}
	return f.Post, nil
}

func (f *postToFlow) checkTopicExist() error {
	if !topic.Dao().FindIfExist(f.TopicID) {
		return errors.New("不存在此帖子")
	}
	return nil
}

func (f *postToFlow) checkUserExist() error {
	if !user.Dao().FindIfExist(f.UserID) {
		return errors.New("不存在此用户")
	}
	return nil
}

func (f *postToFlow) checkContent() error {
	f.Content = strings.Trim(f.Content, " ")
	if len(f.Content) == 0 {
		return errors.New("回复内容为空")
	}
	if len(f.Content) > 64 {
		return errors.New("内容太长，超过64")
	}
	return nil
}

func (f *postToFlow) to() error {
	nPost := post.Post{
		TopicID: f.TopicID,
		UserID:  f.UserID,
		Content: f.Content,
	}
	if err := post.Dao().PostTo(&nPost); err != nil {
		return err
	}
	f.Post = &nPost
	return nil
}
