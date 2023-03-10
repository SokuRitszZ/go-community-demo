package main

import (
	"community-demo/controllers/post"
	"community-demo/controllers/topic"
	"community-demo/controllers/user"
	"community-demo/model"
	"community-demo/utils"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	engine := gin.Default()
	engine.Use(gin.Logger())

	_ = user.InitUserGroup(engine.Group("/user"))
	_ = topic.InitTopicGroup(engine.Group("/topic"))
	_ = post.InitPostGroup(engine.Group("/post"))

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "PONG",
		})
	})
	engine.Run()
}

func Init() error {
	if err := model.Init(); err != nil {
		return err
	}
	if err := utils.InitLogger(); err != nil {
		return err
	}
	return nil
}
