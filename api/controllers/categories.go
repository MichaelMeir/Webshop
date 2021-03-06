package controllers

import (
	"encoding/json"
	"net/http"
	"webshop/api/models"
)

func(c Controller) GetAllCategories(w http.ResponseWriter, req *http.Request) {
	var categories = []models.Category{}
	c.DB.Data.Find(&categories)
	data, err := json.Marshal(categories)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(data)
}

func(c Controller) GetSpecificCategory(w http.ResponseWriter, req *http.Request) {
	var category = models.Category{}
	c.DB.Data.Preload("Products").First(&category, "name = ?", req.URL.Query().Get("name"))
	data, err := json.Marshal(category)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(data)
}