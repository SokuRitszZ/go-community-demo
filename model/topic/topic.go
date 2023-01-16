package topic

import (
	"community-demo/model"
	"community-demo/model/user"
	"community-demo/utils"
	"gorm.io/gorm"
	"sync"
)

type Topic struct {
	gorm.Model
	Title   string    `gorm:"column:title" json:"title"`
	UserID  uint      `gorm:"column:user_id" json:"user_id"`
	Content string    `gorm:"column:content" json:"content"`
	User    user.User `gorm:"ForeignKey:UserID" json:"-"`
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

// Publish C
func (*DAO) Publish(topic *Topic) error {
	if err := model.DB.Create(topic).Error; err != nil {
		utils.Logger.Printf("[Topic]insert error: %s\n", err.Error())
		return err
	}
	return nil
}

// GetById R
func (*DAO) GetById(id uint) (*Topic, error) {
	var topic Topic
	err := model.DB.Where("id = ?", id).Find(&topic).Error
	if err == gorm.ErrRecordNotFound || topic.ID == 0 {
		return nil, nil
	}
	if err != nil {
		utils.Logger.Printf("[Topic]find error: %s\n", err.Error())
		return nil, err
	}
	return &topic, nil
}

func (*DAO) GetByUserId(userID uint) ([]Topic, error) {
	var topics []Topic
	err := model.DB.Where("user_id = ?", userID).Find(&topics).Error
	if err == gorm.ErrRecordNotFound {
		return topics, nil
	}
	if err != nil {
		utils.Logger.Printf("[Topic]find error: %s\n", err.Error())
		return topics, err
	}
	return topics, nil
}
