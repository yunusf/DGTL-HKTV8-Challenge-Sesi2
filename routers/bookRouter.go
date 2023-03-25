package routers

import (
	"challenge-use-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// create
	router.POST("/book", controllers.CreateBook)
	// fetch all
	router.GET("/books", controllers.GetAllBooks)
	// fetch by id
	router.GET("/book/:bookID", controllers.GetBook)
	// update
	router.PUT("/book/:bookID", controllers.UpdateBook)
	// delete
	router.DELETE("/book/:bookID", controllers.DeleteBook)

	return router
}
