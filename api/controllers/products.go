package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func(c Controller) BindCategory(w http.ResponseWriter, req *http.Request) {
	//if user := c.GetUser(w, req); user != nil || user.Admin {
	//	w.WriteHeader(http.StatusForbidden)
	//	w.Write([]byte(strconv.Itoa(http.StatusForbidden) + " No Access"))
	//}
	var product = models.Product{}
	c.DB.Data.First(&product, "uuid = ?", req.URL.Query().Get("product_id"))
	var category = models.Category{}
	c.DB.Data.First(&category, "uuid = ?", req.URL.Query().Get("category_id"))
	if product.Name != "" && category.Name != "" {
		c.DB.Data.Model(&product).Association("Categories").Append(&category)
		out, err := json.Marshal(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " could not parse database model"))
		}
		w.Write(out)
	}else{
		var out = "Product not found"
		if category.Name != "" {
			out = "Category and " + out
			if product.Name != "" {
				out = "Category not found"
			}
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(strconv.Itoa(http.StatusNotFound) + " " + out))
	}
}

func(c Controller) BindImage(w http.ResponseWriter, req *http.Request) {
	if user := c.GetUser(w, req); user != nil || user.Admin {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(strconv.Itoa(http.StatusForbidden) + " No Access"))
	}
	var product = models.Product{}
	c.DB.Data.First(&product, "uuid = ?", req.URL.Query().Get("product_id"))
	var category = models.Category{}
	c.DB.Data.First(&category, "uuid = ?", req.URL.Query().Get("category_id"))
	if product.Name != "" && category.Name != "" {
		c.DB.Data.Model(&product).Association("Categories").Append(c.DB.Data.Model(&category))
	}else{
		var out = "Product not found"
		if category.Name != "" {
			out = "Category and " + out
			if product.Name != "" {
				out = "Category not found"
			}
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(strconv.Itoa(http.StatusNotFound) + " " + out))
	}
}