package post

import (
	post2 "community-demo/handlers/post"
	"community-demo/model"
	"community-demo/model/post"
	"community-demo/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitPostGroup(group *gin.RouterGroup) error {
	err := model.Migrate(&post.Post{})
	if err != nil {
		return err
	}
	group.Use(utils.JWTAuth())
	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"group": "post",
			"msg":   "PONG",
		})
	})
	group.POST("/to", func(c *gin.Context) {
		userID := c.GetUint("id")
		topicIDStr, _ := c.GetPostForm("topic_id")
		topicID, _ := strconv.ParseUint(topicIDStr, 10, 64)
		content, _ := c.GetPostForm("content")
		c.JSON(200, post2.ToHandler(userID, uint(topicID), content))
	})
	group.DELETE("/delete", func(c *gin.Context) {
		userID := c.GetUint("id")
		idStr, _ := c.GetQuery("id")
		id, _ := strconv.ParseUint(idStr, 10, 64)
		c.JSON(200, post2.DeleteHandler(userID, uint(id)))
	})
	return nil
}
