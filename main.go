package main

import (
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
