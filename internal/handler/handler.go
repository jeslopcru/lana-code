package handler

import (
	"github.com/gin-gonic/gin"
	"lana-challenge/internal/pricing"
	"lana-challenge/internal/storage"
	"net/http"
)

type Handler struct {
}

type BasketPayload struct {
	Uuid string `json:"uuid" binding:"required"`
}

type BasketAddedProduct struct {
	Uuid string `json:"uuid" binding:"required"`
	Code string `json:"code" binding:"required"`
}

func (h *Handler) BasketCreated(c *gin.Context) {

	storage.NewInMemoryRepository()
	basket := pricing.NewBasket(storage.InitialData())
	storage.Baskets[basket.UUID] = *basket
	c.JSON(http.StatusCreated, map[string]interface{}{
		"response_message": "basket created",
		"uuid":             basket.UUID,
	})
}

func (h *Handler) BasketAddedProduct(c *gin.Context) {
	var basketAddedProduct BasketAddedProduct

	c.BindJSON(&basketAddedProduct)

	basketStored := storage.Baskets[basketAddedProduct.Uuid]
	_ = basketStored.AddProduct(basketAddedProduct.Code)
	c.JSON(http.StatusOK, map[string]interface{}{
		"response_message": "product added",
		"uuid":             basketStored.UUID,
	})
}

func (h *Handler) BasketEmpty(c *gin.Context) {
	var basketPayload BasketPayload

	c.BindJSON(&basketPayload)

	basketStored := storage.Baskets[basketPayload.Uuid]
	_ = basketStored.RemoveProducts

	c.JSON(http.StatusOK, map[string]interface{}{
		"response_message": "basket empty",
		"uuid":             basketStored.UUID,
	})
}

func (h *Handler) BasketTotal(c *gin.Context) {
	var basketPayload BasketPayload

	c.BindJSON(&basketPayload)

	basketStored := storage.Baskets[basketPayload.Uuid]
	_ = basketStored.Total()

	c.JSON(http.StatusOK, map[string]interface{}{
		"response_message": "basket total",
		"uuid":             basketStored.UUID,
		"products":         basketStored.Products,
		"total":            basketStored.Total(),
	})
}
