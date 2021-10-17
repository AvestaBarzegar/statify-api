package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func goDotEnvVar(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	PORT := goDotEnvVar("PORT")

	router := gin.Default()
	router.GET("/", mockResponse)
	router.Run(":" + PORT)
}
