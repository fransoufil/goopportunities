package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//inicializa o Router utilizando as configuracoes default do gin
	r := gin.Default()
	//Definindo uma Rota
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
