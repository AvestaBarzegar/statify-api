package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var persons = []person{
	{ID: "1", Name: "Avesta"},
}

func mockResponse(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

func main() {
	fmt.Println("Hello World")
	router := gin.Default()
	router.GET("/", mockResponse)

	router.Run()
}
