package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
)

// StartServer load homepage from html file, and responds to the html request.
func StartServer(cache *cache.Cache) *gin.Engine {
	server := gin.Default()
	//Loading home page
	server.LoadHTMLFiles("./static/form.html")
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", gin.H{})
	})
	//Query answer
	server.POST("/answer", func(c *gin.Context) {
		info, ok := cache.Get(c.PostForm("id"))
		if !ok {
			c.JSON(http.StatusOK, "Заданный UID заказа не найден")
			return
		}
		c.JSON(http.StatusOK, info)
	})

	return server
}
