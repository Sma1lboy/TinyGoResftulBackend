package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	//old school way to loading html page
	// file, err := os.Open("./index.html")
	// fmt.Println("opened file")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer file.Close()
	// bytes, err := io.ReadAll(file)
	// if err != nil {
	// 	panic(err.Error())

	// }
	// fmt.Fprintln(c.Writer, string(bytes))

	c.HTML(http.StatusOK, "index.html", gin.H{"users": getAllUsers()})
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

func removeUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	fmt.Println(idStr)
	fmt.Println("asds")
	id, _ := strconv.Atoi(idStr)
	user := removeUser(id)
	c.JSON(http.StatusOK, user)
}

// func userUpdate(res http.ResponseWriter, req *http.Request) {

// }
func getAllHandler(c *gin.Context) {
	c.JSON(http.StatusOK, getAllUsers())
}
func main() {
	r := gin.Default()
	r.Static("/static", "./static/")
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", indexHandler)
	//getall
	r.GET("/user/all", getAllHandler)
	r.POST("/user/add", addUserHandler)
	r.GET("/user/remove/:id", removeUserHandler)

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
