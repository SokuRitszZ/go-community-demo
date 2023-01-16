package post

import (
	"community-demo/model"
	"community-demo/model/topic"
	"community-demo/model/user"
	"community-demo/utils"
	"errors"
	"gorm.io/gorm"
	"sync"
)

type Post struct {
	gorm.Model
	TopicID uint         `gorm:"column:topic_id" json:"topic_id"`
	UserID  uint         `gorm:"column:user_id" json:"user_id"`
	Content string       `gorm:"column:content" json:"content"`
	Topic   *topic.Topic `gorm:"ForeignKey:TopicID" json:"-"`
	User    *user.User   `gorm:"ForeignKey:UserID" json:"-"`
}

type DAO struct{}

var dao *DAO
var once sync.Once

func Dao() *DAO {
	once.Do(
		func() {
			dao = &DAO{}
		})
	return dao
}

func (*DAO) PostTo(post *Post) error {
	err := model.DB.Create(post).Error
	if err != nil {
		utils.Logger.Printf("[Post]create error: %s\n", err.Error())
		return err
	}
	return nil
}

func (*DAO) DeleteById(userID, id uint) error {
	var post Post
	err := model.DB.Where("user_id = ? and id = ?", userID, id).First(&post).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("不存在此回帖，或者没有权限删除")
	}
	if err != nil {
		return errors.New("不存在此回帖，或者没有权限删除")
	}
	err = model.DB.Where("id = ?", id).Unscoped().Delete(&post).Error
	if err != nil {
		return errors.New("不存在此回帖，或者没有权限删除")
		//utils.Logger.Printf("[Post]delete error: %s\n", err.Error())
		//return err
	}
	return nil
}
