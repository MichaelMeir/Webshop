package controllers

import (
	"encoding/json"
	"net/http"
	"webshop/api/models"
)

func(c Controller) GetAllProducts(w http.ResponseWriter, req *http.Request) {
	var products = []models.Product{}
	c.DB.Data.Preload("Categories").Find(&products)
	data, err := json.Marshal(products)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}

func(c Controller) GetSpecificProduct(w http.ResponseWriter, req *http.Request) {
	var product = models.Product{}
	c.DB.Data.Preload("Categories").First(&product, "uuid = ?", req.URL.Query().Get("id"))
	data, err := json.Marshal(product)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}

func(c Controller) GetProductImage(w http.ResponseWriter, req *http.Request) {
	var products = []models.Product{}
	c.DB.Data.Preload("Images").Find(&products)
	data, err := json.Marshal(products)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}