package _httpUtils

import (
	"encoding/json"
	"fmt"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	_domain "github.com/easysoft/zagent/internal/pkg/domain"
	_i118Utils "github.com/easysoft/zagent/internal/pkg/lib/i118"
	_logUtils "github.com/easysoft/zagent/internal/pkg/lib/log"
	"github.com/easysoft/zagent/internal/pkg/var"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string, requestTo string) (interface{}, bool) {
	client := &http.Client{}

	if _var.Verbose {
		_logUtils.Info(url)
	}

	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return nil, false
	}

	resp, respErr := client.Do(req)

	if respErr != nil {
		_logUtils.Error(respErr.Error())
		return nil, false
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if _var.Verbose {
		_logUtils.PrintUnicode(bodyStr)
	}
	defer resp.Body.Close()

	if requestTo == "farm" {
		var bodyJson _domain.RpcResp
		jsonErr := json.Unmarshal(bodyStr, &bodyJson)
		if jsonErr != nil {
			if strings.Index(string(bodyStr), "<html>") > -1 {
				_logUtils.Error(_i118Utils.Sprintf("wrong_response_format", "html"))
				return nil, false
			} else {
				_logUtils.Error(jsonErr.Error())
				return nil, false
			}
		}
		code := bodyJson.Code
		return bodyJson.Payload, code == _const.ResultSuccess
	} else {
		var bodyJson map[string]interface{}
		jsonErr := json.Unmarshal(bodyStr, &bodyJson)
		if jsonErr != nil {
			_logUtils.Error(jsonErr.Error())
			return nil, false
		} else {
			return bodyJson, true
		}
	}
}

func Post(url string, params interface{}) (interface{}, bool) {
	if _var.Verbose {
		_logUtils.Info(url)
	}
	client := &http.Client{}

	paramStr, err := json.Marshal(params)
	if err != nil {
		_logUtils.Error(err.Error())
		return nil, false
	}

	req, reqErr := http.NewRequest("POST", url, strings.NewReader(string(paramStr)))
	if reqErr != nil {
		_logUtils.Error(reqErr.Error())
		return nil, false
	}

	req.Header.Set("Content-Type", "application/json")

	resp, respErr := client.Do(req)
	if respErr != nil {
		_logUtils.Error(respErr.Error())
		return nil, false
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)
	if _var.Verbose {
		_logUtils.PrintUnicode(bodyStr)
	}

	var result _domain.RpcResp
	json.Unmarshal(bodyStr, &result)

	defer resp.Body.Close()

	code := result.Code
	return result, code == _const.ResultSuccess
}

func GenUrl(server string, path string) string {
	server = UpdateUrl(server)
	url := fmt.Sprintf("%sapi/v1/%s", server, path)
	return url
}

func UpdateUrl(url string) string {
	if strings.LastIndex(url, "/") < len(url)-1 {
		url += "/"
	}
	return url
}
