package main

import (
	"github.com/gin-gonic/gin"
	"lana-challenge/internal/handler"
)

func main() {

	r := getGinEngine()

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func getGinEngine() *gin.Engine {
	r := gin.Default()
	h := handler.Handler{}
	r.POST("/basket", h.BasketCreated)
	r.POST("/basket/add", h.BasketAddedProduct)
	r.POST("/basket/delete", h.BasketEmpty)
	r.POST("/basket/total", h.BasketTotal)
	return r
}
