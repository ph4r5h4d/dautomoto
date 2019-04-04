package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	q := Queue{}
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/queue", q.addToQueue)

	r.GET("/get", func(c *gin.Context) {
		queue := getQueue(&q)
		if links, err := queue.getLinks(1); err != nil {
			c.JSON(400, gin.H{
				"links": links,
			})
		}
	})

	if err := r.Run(":9595"); err != nil {
		panic("Unable to start serve")
	}
}

func (q *Queue) addToQueue(c *gin.Context) {
	var d DownloadRequest
	if err := c.BindJSON(&d); err != nil {
		c.JSON(400, gin.H{
			"err": err,
		})
		return
	}

	queue := getQueue(q)

	queue.enqueue(d.Links)

	c.JSON(200, gin.H{
		"message": "Added to queue",
	})

}
