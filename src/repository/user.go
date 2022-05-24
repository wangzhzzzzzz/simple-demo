package repository

import (
	"github.com/RaymondCode/simple-demo/src/database"
	"log"
	"sync"
)

type User struct {
	Id            int64  `gorm:"column:id"`
	Name          string `gorm:"column:name"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao //空结构体节省内存
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	//Do里面的函数只会执行一次
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) CreateUser(user *User) error {
	if err := database.MySqlDb.Create(user).Error; err != nil {
		log.Println("insert err", err.Error())
	}
	return nil
}
