package user

import (
	user2 "community-demo/handlers/user"
	"community-demo/model"
	"community-demo/model/user"
	"community-demo/utils"
	"github.com/gin-gonic/gin"
)

func InitUserGroup(group *gin.RouterGroup) error {
	err := model.Migrate(&user.User{})
	if err != nil {
		return err
	}
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"group": "user",
			"msg":   "PONG",
		})
	})
	group.POST("/register", func(c *gin.Context) {
		name, _ := c.GetPostForm("name")
		password, _ := c.GetPostForm("password")
		c.JSON(200, user2.RegisterHandler(name, password))
	})
	group.POST("/login", func(c *gin.Context) {
		name, _ := c.GetPostForm("name")
		password, _ := c.GetPostForm("password")
		c.JSON(200, user2.LoginHandler(name, password))
	})
	group.GET("/get_info", utils.JWTAuth(), func(c *gin.Context) {
		id := c.GetUint("id")
		c.JSON(200, user2.GetInfoHandler(id))
	})

	return nil
}
