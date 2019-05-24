package Env

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type ENV map[string]string

func LoadENV() ENV {
	bytes, err := ioutil.ReadFile(".env")
	if os.IsNotExist(err) {
		panic(".env file does not exist!")
	}else if err != nil {
		panic(err)
	}
	envcontent := strings.Split(string(bytes), "\n")
	env := make(ENV)
	for i := 0; i < len(envcontent); i++ {
		if strings.Contains(envcontent[i], "=") {
			split := strings.Split(envcontent[i], "=")
			env[split[0]] = strings.Join(split[1:], "=")
		}
	}
	return env
}

func(env ENV) IsEmpty(key ...string) bool {
	for i := 0; i < len(key); i++ {
		if len(env[key[i]]) == 0 {
			return true
		}
	}
	return false
}

func(env ENV) AsInt(key string) int {
	if i, ok := strconv.Atoi(env[key]); ok == nil {
		return i
	}
	return 0
}
