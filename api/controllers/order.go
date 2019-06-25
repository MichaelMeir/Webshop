package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webshop/api/models"
)

func(c Controller) RegisterOrder(w http.ResponseWriter, req *http.Request) {
	type Products []struct {
		UUID string `json:uuid`
	}
	var products = Products{}
	json.Unmarshal([]byte(req.URL.Query()["products"][0]), &products)
	var total uint = 0
	for i := 0; i < len(products); i++ {
		var product = models.Product{}
		c.DB.Data.First(&product, "uuid = ?", products[i].UUID)
		total += product.Price
	}
	user := c.GetUser(w, req)
	if user == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " Could not find session cookie."))
	}else{
		data, err := json.Marshal(*user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " Could not parse user."))
		}
		var order = &models.Order{
			Products: req.URL.Query()["products"][0],
			Price: total,
			User: string(data),
		}
		c.DB.Data.Save(order)
		w.Write([]byte("{\"success\": true}"))
	}
}
