package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/orders", func(c *gin.Context) {
		resp, err := http.Post("http://localhost:8081/orders", "application/json", c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "order service down"})
			return
		}
		defer resp.Body.Close()

		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	})

	r.GET("/products", func(c *gin.Context) {
		resp, err := http.Post("http://localhost:8082/inventory", "application/json", c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "inventory service down"})
			return
		}
		defer resp.Body.Close()

		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	})

	r.Run(":8080")
}
