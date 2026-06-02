package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"method":"GET"})
}
func posting(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"method":"Post"})
}
func putting(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"method":"put"})
}
func deleting(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"method":"delete"})
}
func patching(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"method":"patch"})
}
func head(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"method":"head"})
}
func options(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"method":"options"})
}
func main(){
	router:=gin.Default()
	router.GET("/someget",getting)
	router.POST("/somepost",posting)
	router.PUT("/someput",putting)
	router.DELETE("/somedelete",deleting)
	router.PATCH("/somepatch",patching)
	router.HEAD("/somehead",head)
	router.OPTIONS("/someoptions",options)
	router.Run(":3030")
}
