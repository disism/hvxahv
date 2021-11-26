package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/activity"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"io/ioutil"
)

func InboxHandler(c *gin.Context) {
	name := c.Param("actor")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	fmt.Println(string(body))

	activity.Types(name, body)
}

func OutboxHandler(c *gin.Context) {
	name := middleware.GetUserName(c)
	t := c.PostForm("type")
	o := c.PostForm("object")

	fmt.Println(name, t, o)
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "ok!",
	})
}
