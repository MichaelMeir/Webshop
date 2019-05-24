package api

import (
	"encoding/json"
	"io/ioutil"
	"webshop/Env"
)

type Call struct {
	Path     string `json:"path"`
	Method   string `json:"method"`
	Required map[string] string `json:"required"`
	Perform string `json:"perform"`
}

type Rules map[string] string

var rules Rules
var calls []Call

func LoadCalls(env Env.ENV) {
	data, err := ioutil.ReadFile(env["CALLS"])
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &calls)
	if err != nil {
		panic(err)
	}
}

func LoadRules(env Env.ENV) {
	data, err := ioutil.ReadFile(env["RULES"])
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &rules)
	if err != nil {
		panic(err)
	}
}
