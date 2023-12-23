package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	file, err := os.Open("./index.html")
	fmt.Println("opened file")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err.Error())

	}
	fmt.Fprintln(c.Writer, string(bytes))
}
func cssHandler(c *gin.Context) {
	file, err := os.Open("./style.css")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintln(c.Writer, string(bytes))
}

func addUserHandler(c *gin.Context) {
	var user User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	user.ID = id
	id++
	addUser(user)
}

// func userRemove(res http.ResponseWriter, req *http.Request) {

// }
// func userUpdate(res http.ResponseWriter, req *http.Request) {

// }
func getAllHandler(c *gin.Context) {
	c.JSON(http.StatusOK, getAllUsers())
}
func main() {
	r := gin.Default()
	r.GET("/", indexHandler)
	r.GET("/style.css", cssHandler)

	//getall
	r.GET("/user/all", getAllHandler)
	r.POST("/user/add", addUserHandler)

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
