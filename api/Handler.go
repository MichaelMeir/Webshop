package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func getCall(path string) *Call {
	for i := 0; i < len(calls); i++ {
		if equalsIgnoreSlashes(calls[i].Path, path) {
			return &calls[i]
		}
	}
	return nil
}

func equalsIgnoreSlashes(a string, b string) bool {
	a = strings.TrimSuffix(strings.TrimPrefix(a, "/"), "/")
	b = strings.TrimSuffix(strings.TrimPrefix(b, "/"), "/")
	return a == b
}

func Handle(w http.ResponseWriter, req *http.Request) {
	var call = getCall(req.URL.Path)
	if call == nil {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(strconv.Itoa(http.StatusNotFound) + " api call not found."))
		if err != nil {
			println(err.Error())
		}
		return
	}
	query := req.URL.Query()
	err := AddBodyQuery(&query, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(strconv.Itoa(http.StatusNotFound) + " error while reading body: " + err.Error()))
		return
	}
	if strings.ToLower(req.Method) == strings.ToLower(call.Method) {
		for key, rule := range call.Required {
			if val, ok := req.URL.Query()[key]; ok {
				if FollowsRule(val[0], rule) {
					continue
				}
			}
			w.WriteHeader(http.StatusExpectationFailed)
			_, err := w.Write([]byte(strconv.Itoa(http.StatusExpectationFailed) + " missing expected parameters."))
			if err != nil {
				println(err.Error())
			}
			return
		}
		callFunc(call.Perform, w, req)
	}else{
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(strconv.Itoa(http.StatusBadRequest) + " unexpected method used for api call."))
		if err != nil {
			println(err.Error())
		}
		return
	}
}

func AddBodyQuery(query *url.Values, req *http.Request) error {
	data, err := ioutil.ReadAll(req.Body)
	if len(data) == 0 {
		return nil
	}
	if err != nil {
		return err
	}
	body := map[string]interface{}{}
	err = json.Unmarshal(data, &body)
	if err != nil {
		return err
	}
	for key, value := range body {
		if val, ok := value.(string); ok {
			query.Set(key, val)
		}
	}
	req.URL.RawQuery = query.Encode()
	return nil
}

func FollowsRule(value string, rule string) bool {
	rules := PrepareRule(rule)
	for i := 0; i < len(rules); i++ {
		matched, err := regexp.MatchString(rules[i], value)
		if err != nil || !matched {
			return false
		}
	}
	return true
}

func PrepareRule(rule string) []string {
	split := strings.Split(rule, "|")
	out := []string{}
	for i := 0; i < len(split); i++ {
		args := strings.Split(split[i], ":")
		outrule := rules[args[0]]
		for j := 1; j < len(args); j++ {
			outrule = strings.ReplaceAll(outrule, "%" + strconv.Itoa(j) + "%", args[j])
		}
		out = append(out, outrule)
	}
	return out
}