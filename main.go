package main

// todo 导入 go - gin
import (
	. "github.com/tattoo1880/testgoinit/config"
	"github.com/tattoo1880/testgoinit/service"

	"github.com/gin-gonic/gin"
)

func main() {

	NewRabbitMQ()
	defer MyRabbitMQ.Close()

	go service.Suckit("test")

	r := gin.Default()
	r.GET("/pub", func(c *gin.Context) {

		myquery := c.Query("msg")
		if myquery == "" {
			c.JSON(400, gin.H{"error": "msg is required"})
			return
		}

		// err := MyRabbitMQ.Publish("test", "hello world")
		err := MyRabbitMQ.Publish("test", myquery)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "published"})
	})

	err := r.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

}
