package controllers

import (
	"github.com/blend/go-sdk/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
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

func(c Controller) AuthRegister(w http.ResponseWriter, req *http.Request) {
	pass, err := bcrypt.GenerateFromPassword([]byte(req.URL.Query()["password"][0]), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " Could not hash password: " + err.Error()))
	}

	var ses_id = uuid.V4().String()

	var user = &models.User{
		Username: req.URL.Query()["username"][0],
		Password: string(pass),
		Email: req.URL.Query()["email"][0],
		Session_ip: req.RemoteAddr,
		Session_id: ses_id,
	}
	c.DB.Data.Save(user)
	req.AddCookie(&http.Cookie{Name: "ses_id", Value: ses_id})
	w.Write([]byte("{\"success\": true}"))
}

func(c Controller) AuthLogin(w http.ResponseWriter, req *http.Request) {
	var user = models.User{}
	c.DB.Data.First(&user, "username = ?", req.URL.Query()["username"][0])
	if user.Password != "" {
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.URL.Query()["password"][0])) != nil {
			goto err
		}else{
			user.Session_ip = req.RemoteAddr
			var ses_id = uuid.V4().String()
			user.Session_id = ses_id
			c.DB.Data.Save(user)
			req.AddCookie(&http.Cookie{Name: "ses_id", Value: ses_id})
			w.Write([]byte("{\"success\": true}"))
		}
	}
	err:
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("{\"success\": false}"))
}

func(c Controller) AuthLogout(w http.ResponseWriter, req *http.Request) {
	if cookie, err := req.Cookie("ses_id"); err == nil {
		var user = models.User{}
		c.DB.Data.First(&user, "ses_id = ?, ses_ip = ?", cookie.Value, req.RemoteAddr)
		if user.Username == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " Could not find user linked to session"))
		}else{
			user.Session_id = ""
			user.Session_ip = ""
			req.AddCookie(&http.Cookie{Name: "ses_id", Value: ""})
			c.DB.Data.Save(user)
			w.Write([]byte("{\"success\": true}"))
		}
	}else{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " Could not find session cookie: " + err.Error()))
	}
}