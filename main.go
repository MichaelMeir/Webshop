package main

import (
	"io/ioutil"
	"mime"
	"net/http"
	"strings"
	"webshop/Env"
	"webshop/FileHandler"
	"webshop/api"
)

var env Env.ENV = Env.LoadENV()

func main() {
	api.LoadCalls(env)
	api.LoadRules(env)
	http.HandleFunc("/", Handler)
	if !env.IsEmpty("SSL_KEY", "SSL_CERT", "HTTPS_PORT") {
		key, err := ioutil.ReadFile(env["SSL_KEY"])
		cert, err := ioutil.ReadFile(env["SSL_CERT"])
		if err != nil {
			panic(err)
		}
		go http.ListenAndServeTLS(":" + env["HTTPS_PORT"], string(key), string(cert), nil)
	}
	println("starting host on port " + env["HTTP_PORT"])
	err := http.ListenAndServe(":" + env["HTTP_PORT"], nil)
	if err != nil {
		panic(err)
	}
}

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if strings.HasPrefix(req.URL.Path, env["API"]) {
		req.URL.Path = strings.TrimPrefix(req.URL.Path, env["API"])
		api.Handle(w, req)
		return
	}else if strings.HasPrefix(req.URL.Path, env["STATIC"]) {
		w.Header().Set("Content-Type", mime.TypeByExtension(req.URL.Path))
		_, err := w.Write([]byte(FileHandler.GetFile(env["ROOT"], req.URL.Path)))
		if err != nil {
			panic(err)
		}
		return
	}
	_, err := w.Write([]byte(FileHandler.GetFile(env["ROOT"], env["INDEX"])))
	if err != nil {
		panic(err)
	}
}