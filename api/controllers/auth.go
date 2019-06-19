package controllers

import (
	"encoding/json"
	"github.com/blend/go-sdk/uuid"
	"golang.org/x/crypto/bcrypt"
	"net"
	"net/http"
	"strconv"
	"time"
	"webshop/api/models"
)

func(c Controller) GetUser(w http.ResponseWriter, req *http.Request) *models.User {
	session_id, err := req.Cookie("ses_id")
	if err != nil {
		return nil
	}
	var user = models.User{}
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	c.DB.Data.First(&user, "session_id = ? AND session_ip = ?", session_id.Value, host)
	if user.Username != "" {
		return &user
	}
	return nil
}

func(c Controller) AuthRegister(w http.ResponseWriter, req *http.Request) {
	var check = models.User{}
	c.DB.Data.First(&check, "username = ?", req.URL.Query()["username"][0])
	if check.Password != "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " User already exists!"))
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(req.URL.Query()["password"][0]), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " Could not hash password: " + err.Error()))
	}

	var ses_id = uuid.V4().String()

	host, _, _ := net.SplitHostPort(req.RemoteAddr)

	var user = &models.User{
		Username: req.URL.Query()["username"][0],
		Password: string(pass),
		Email: req.URL.Query()["email"][0],
		Session_ip: host,
		Session_id: ses_id,
	}
	c.DB.Data.Save(user)
	http.SetCookie(w, &http.Cookie{Name: "ses_id", Value: ses_id, Expires: time.Now().Add(time.Hour*24), Path: "/"})
	w.Write([]byte("{\"success\": true}"))
}

func(c Controller) AuthLogin(w http.ResponseWriter, req *http.Request) {
	var user = models.User{}
	c.DB.Data.First(&user, "username = ?", req.URL.Query()["username"][0])
	if user.Password != "" {
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.URL.Query()["password"][0])) != nil {
			goto err
		}else{
			host, _, _ := net.SplitHostPort(req.RemoteAddr)
			user.Session_ip = host
			var ses_id = uuid.V4().String()
			user.Session_id = ses_id
			c.DB.Data.Save(user)
			http.SetCookie(w, &http.Cookie{Name: "ses_id", Value: ses_id, Expires: time.Now().Add(time.Hour*24), Path: "/"})
			w.Write([]byte("{\"success\": true}"))
			return
		}
	}
	err:
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("{\"success\": false}"))
}

func(c Controller) AuthLogout(w http.ResponseWriter, req *http.Request) {
	if user := c.GetUser(w, req); user != nil {
		user.Session_id = ""
		user.Session_ip = ""
		http.SetCookie(w, &http.Cookie{Name: "ses_id", Value: "", Expires: time.Now().Add(time.Hour*24), Path: "/"})
		c.DB.Data.Save(user)
		w.Write([]byte("{\"success\": true}"))
	}else{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " Could not find session cookie."))
	}
}

func(c Controller) GetUserInfo(w http.ResponseWriter, req *http.Request) {
	if user := c.GetUser(w, req); user != nil {
		type UserInfo struct {
			Name string
			Email string
			Address string
			Zip_code string
			Province string
			Country string
		}
		data, err := json.Marshal(UserInfo{user.Username, user.Email, user.Address, user.ZipCode, user.Province, user.Country})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " Could not convert user data to json " + err.Error()))
		}
		w.Write(data)
	}else{
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(strconv.Itoa(http.StatusForbidden) + " Not logged in."))
	}
}

func(c Controller) UpdateUserInfo(w http.ResponseWriter, req *http.Request) {
	if user := c.GetUser(w, req); user != nil {
		user.Email = req.URL.Query()["Email"][0]
		user.Address = req.URL.Query()["Address"][0]
		user.ZipCode = req.URL.Query()["Zip_code"][0]
		user.Province = req.URL.Query()["Province"][0]
		user.Country = req.URL.Query()["Country"][0]
		c.DB.Data.Save(user)
		w.Write([]byte("{\"success\":true}"))
	}else{
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(strconv.Itoa(http.StatusForbidden) + " Not logged in."))
	}
}