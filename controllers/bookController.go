package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID string `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBK Book

	if err := ctx.ShouldBindJSON(&newBK); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBK.BookID = fmt.Sprintf("c%d", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBK)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBK,
	})
}

func GetAllBooks(ctx *gin.Context) {

	var all = []Book{}

	for i := range BookDatas {
		all = append(all, BookDatas[i])
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": all,
	})
}

func GetBook(ctx *gin.Context) {
	bkID := ctx.Param("bookID")
	condition := false
	var bkData Book

	for i, v := range BookDatas {
		if bkID == v.BookID {
			condition = true
			bkData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", bkID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bkData,
	})
}

func UpdateBook(ctx *gin.Context) {
	bkID := ctx.Param("bookID")
	condition := false
	var UpdateBook Book

	if err := ctx.ShouldBindJSON(&UpdateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	for i, v := range BookDatas {
		if bkID == v.BookID {
			condition = true
			BookDatas[i] = UpdateBook
			BookDatas[i].BookID = bkID
			break
		}
	}
	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bkID),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully updated", bkID),
	})
}

func DeleteBook(ctx *gin.Context) {
	bkID := ctx.Param("bookID")
	condition := false
	var bk_i int

	for i, car := range BookDatas {
		if bkID == car.BookID {
			condition = true
			bk_i = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bkID),
		})
		return
	}

	copy(BookDatas[bk_i:], BookDatas[bk_i+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully deleted", bkID),
	})
}
