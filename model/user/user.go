package user

import (
	"community-demo/model"
	"community-demo/utils"
	"errors"
	"gorm.io/gorm"
	"sync"
)

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;not null"`
	Password string `gorm:"column:password;not null" json:"-"`
	Avatar   string `gorm:"column:avatar"`
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

// Register C
func (*DAO) Register(user *User) error {
	if err := model.DB.Create(user).Error; err != nil {
		utils.Logger.Printf("[User]insert error: %s\n", err.Error())
		return err
	}
	return nil
}

// GetInfoById R
func (*DAO) GetInfoById(id int64) (*User, error) {
	var user User
	err := model.DB.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		utils.Logger.Printf("[User]find error: %s\n", err.Error())
		return nil, err
	}
	return &user, nil
}

func (*DAO) FindByName(name string) (*User, error) {
	var user User
	err := model.DB.Where("name = ?", name).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		utils.Logger.Printf("[User]find error: %s\n", err.Error())
		return nil, err
	}
	if user.ID != 0 {
		return &user, nil
	}
	return nil, nil
}

func (*DAO) CheckMatch(name, password string) (*User, error) {
	var user User
	err := model.DB.Where("name = ? and password = ?", name, password).Find(&user).Error
	if err == gorm.ErrRecordNotFound || user.ID == 0 {
		return nil, nil
	}
	if err != nil {
		utils.Logger.Printf("[User]find error: %s\n", err.Error())
		return nil, err
	}
	return &user, nil
}

func (*DAO) GetBaseData(id uint) (*User, error) {
	var user User
	err := model.DB.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound || user.ID == 0 {
		return nil, errors.New("找不到这个用户")
	}
	if err != nil {
		utils.Logger.Printf("[User]find error: %s\n", err.Error())
		return nil, err
	}
	return &user, nil
}

func (*DAO) FindIfExist(id uint) bool {
	var user User
	err := model.DB.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound || user.ID == 0 {
		return false
	}
	if err != nil {
		utils.Logger.Printf("[User]find error: %s\n", err.Error())
		return false
	}
	return true
}
