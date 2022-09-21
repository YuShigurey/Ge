package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
func main() {
	r := gin.Default()
	r.GET("/", getIndex)
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,
			gin.H{
				"message": "pong",
			})
	})

	err := r.Run()

	if err != nil {
		return
	} else {
		fmt.Println(err)
	}
}

func getIndex(ctx *gin.Context) {

	var listOfRoutePath []string
	for _, route := range gin.Default().Routes() {
		listOfRoutePath = append(listOfRoutePath, route.Path)
	}

	ctx.JSON(http.StatusOK,
		map[string][]string{"apis": listOfRoutePath})

}
