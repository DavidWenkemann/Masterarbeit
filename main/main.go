package main

import (
	"fmt"

	//"errors"
	"example/Masterarbeit/basedata"
	//"example/Masterarbeit/basedata"
	//"github.com/gin-gonic/gin"
)

/*
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, basedata.GetProducts)
}
*/

func main() {
	//router := gin.Default()
	//router.GET("/products", getProducts)
	//router.Run("localhost:8080")

	fmt.Printf(basedata.GetProductsByEAN("4029764001807").Name)
}
