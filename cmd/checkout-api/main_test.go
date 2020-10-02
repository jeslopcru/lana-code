package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"lana-challenge/internal/handler"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateABasket(t *testing.T) {
	router := getGinEngine()

	resp := createBasket(router)

	assert.Contains(t, resp.Body.String(), "basket created")
	assert.Equal(t, resp.Code, 201)
}

func TestAddProductToBasket(t *testing.T) {

	router := getGinEngine()
	basketResponse := createBasket(router)
	var basketCreated handler.BasketPayload
	_ = json.NewDecoder(basketResponse.Body).Decode(&basketCreated)

	resp := addProduct(basketCreated, router)

	assert.Equal(t, resp.Code, 200)
	assert.Contains(t, resp.Body.String(), "product added")

}

func addProduct(basketCreated handler.BasketPayload, router *gin.Engine) *httptest.ResponseRecorder {
	payload := strings.NewReader(fmt.Sprintf("{\"uuid\":\"%s\",\"code\":\"PEN\"}", basketCreated.Uuid))
	req, _ := http.NewRequest("POST", "/basket/add", payload)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	return resp
}

func createBasket(router *gin.Engine) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/basket", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	return resp
}

func TestEmptyBasket(t *testing.T) {

	router := getGinEngine()
	basketResponse := createBasket(router)
	var basketCreated handler.BasketPayload
	_ = json.NewDecoder(basketResponse.Body).Decode(&basketCreated)
	_ = addProduct(basketCreated, router)

	resp := emptyBasket(basketCreated, router)

	assert.Equal(t, resp.Code, 200)
	assert.Contains(t, resp.Body.String(), "basket empty")
}

func emptyBasket(basketCreated handler.BasketPayload, router *gin.Engine) *httptest.ResponseRecorder {
	payload := strings.NewReader(fmt.Sprintf("{\"uuid\":\"%s\"}", basketCreated.Uuid))
	req, _ := http.NewRequest("POST", "/basket/delete", payload)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	return resp
}

func TestTotal(t *testing.T) {

	router := getGinEngine()
	basketResponse := createBasket(router)
	var basketCreated handler.BasketPayload
	_ = json.NewDecoder(basketResponse.Body).Decode(&basketCreated)
	_ = addProduct(basketCreated, router)

	payload := strings.NewReader(fmt.Sprintf("{\"uuid\":\"%s\"}", basketCreated.Uuid))
	req, _ := http.NewRequest("POST", "/basket/total", payload)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, 200)
	assert.Contains(t, resp.Body.String(), "basket total")
	assert.Contains(t, resp.Body.String(), `"total":5`)
	assert.Contains(t, resp.Body.String(), `"products"`)
}
