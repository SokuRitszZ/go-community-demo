package topic

import (
	"community-demo/service/topic"
	"community-demo/utils"
)

func PublishHandler(title string, userID uint, content string) *utils.Message {
	topic, err := topic.PublishService(title, userID, content)
	if err != nil {
		return &utils.Message{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &utils.Message{
		Code: 0,
		Msg:  "OK",
		Data: topic,
	}
}

func GetDetailHandler(id uint) *utils.Message {
	topic, err := topic.GetDetailService(id)
	if err != nil {
		return &utils.Message{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &utils.Message{
		Code: 0,
		Msg:  "OK",
		Data: topic,
	}
}

func GetListHandler(userID uint) *utils.Message {
	topics, err := topic.GetListService(userID)
	if err != nil {
		return &utils.Message{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &utils.Message{
		Code: 0,
		Msg:  "OK",
		Data: topics,
	}
}

func DeleteHandler(userID, id uint) *utils.Message {
	err := topic.DeleteService(userID, id)
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
