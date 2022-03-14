package main

import "github.com/gin-gonic/gin"


type User struct {
	Id int64
	userName string
	
}
func main() {
	r := gin.Default()

	userRouter :=r.Group("/users")
	{
		userRouter.GET("/", getUsers)
	}

	r.Run(":3000")
}

func getUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"error":   false,
		"message": "Hello",
	})
}
