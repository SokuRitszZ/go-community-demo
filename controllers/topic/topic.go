package topic

import (
	topic2 "community-demo/handlers/topic"
	"community-demo/model"
	"community-demo/model/topic"
	"community-demo/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitTopicGroup(group *gin.RouterGroup) error {
	err := model.Migrate(&topic.Topic{})
	if err != nil {
		return err
	}
	group.Use(utils.JWTAuth())
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"group": "topic",
			"msg":   "PONG",
		})
	})
	group.POST("/publish", func(c *gin.Context) {
		title, _ := c.GetPostForm("title")
		userID := c.GetUint("id")
		content, _ := c.GetPostForm("content")
		c.JSON(200, topic2.PublishHandler(title, userID, content))
	})
	group.GET("/detail", func(c *gin.Context) {
		var idStr, _ = c.GetQuery("id")
		var id, _ = strconv.ParseUint(idStr, 10, 64)
		c.JSON(200, topic2.GetDetailHandler(uint(id)))
	})
	group.GET("/list/me", func(c *gin.Context) {
		id := c.GetUint("id")
		c.JSON(200, topic2.GetListHandler(id))
	})
	group.GET("/list/:userID", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("userID"), 10, 64)
		c.JSON(200, topic2.GetListHandler(uint(id)))
	})

	return nil
}
