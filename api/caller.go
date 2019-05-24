package api

import (
	"net/http"
	"reflect"
	"strconv"
	"webshop/api/controllers"
	"webshop/api/models"
)

var controller controllers.Controller = controllers.Controller{models.OpenDatabase()}

func callFunc(name string, w http.ResponseWriter, req *http.Request) {
	if _, ok := reflect.TypeOf(controller).MethodByName(name); ok {
		method := reflect.ValueOf(controller).MethodByName(name)
		in := make([]reflect.Value, method.Type().NumIn())
		in[0] = reflect.ValueOf(w)
		in[1] = reflect.ValueOf(req)
		method.Call(in)
	}else{
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(strconv.Itoa(http.StatusInternalServerError) + " api call not defined."))
		if err != nil {
			println(err.Error())
		}
	}
}
