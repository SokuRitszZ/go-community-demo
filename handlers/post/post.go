package post

import (
	"community-demo/service/post"
	"community-demo/utils"
)

func ToHandler(userID, topicID uint, content string) *utils.Message {
	post, err := post.ToService(userID, topicID, content)
	if err != nil {
		return &utils.Message{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &utils.Message{
		Code: 0,
		Msg:  "OK",
		Data: post,
	}
}

func DeleteHandler(userID, id uint) *utils.Message {
	err := post.DeleteService(userID, id)
	if err != nil {
		return &utils.Message{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &utils.Message{
		Code: 0,
		Msg:  "OK",
		Data: nil,
	}
}
