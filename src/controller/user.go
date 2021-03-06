package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/src/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	//调用service层注册服务，返回id
	token := username + password
	userId, err := service.CreateUser(username, password)
	if err != nil {
		log.Println("insert err", err)
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   userId,
		Token:    token,
	})
	//hello
	/*
		if _, exist := usersLoginInfo[token]; exist {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
			})
		} else {
			atomic.AddInt64(&userIdSequence, 1)
			newUser := User{
				Id:   userIdSequence,
				Name: username,
			}
			usersLoginInfo[token] = newUser
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   userIdSequence,
				Token:    username + password,
			})
		}

	*/
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username
	userInfo, err := service.QueryUserInfo(username)
	if err != nil || userInfo == nil {
		log.Println(err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	fmt.Println(userInfo.User.Password)
	if password != userInfo.User.Password {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   userInfo.User.Id,
		Token:    token,
	})

	/*
		if user, exist := usersLoginInfo[token]; exist {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.Id,
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
		}
	*/
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userInfo, err := service.QueryUserInfo(token)
	if err != nil || userInfo == nil {
		log.Println(err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
	user := User{
		Id:            userInfo.User.Id,
		Name:          userInfo.User.Name,
		FollowCount:   userInfo.User.FollowCount,
		FollowerCount: userInfo.User.FollowerCount,
		IsFollow:      false,
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})
	/*
		if user, exist := usersLoginInfo[token]; exist {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 0},
				User:     user,
			})
		} else {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
		}

	*/
}
