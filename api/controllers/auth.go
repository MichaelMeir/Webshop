package controllers

import (
	"net/http"
	"webshop/api/models"
)

func(c Controller) GetUser(w http.ResponseWriter, req *http.Request) *models.User {
	session_id, err := req.Cookie("ses_id")
	if err != nil {
		return nil
	}
	var user = models.User{}
	c.DB.Data.First(&user, "session_id = ?, session_ip = ?", session_id, req.RemoteAddr)
	if user.Username != "" {
		return &user
	}
	return nil
}