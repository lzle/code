package module

import (
	"bytes"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	log "go-calltask/log"
	"io/ioutil"
	"net/http"
	Url "net/url"
	"strings"
	"time"
)

var (
	HTTP_Client *http.Client
)

func init() {
	HTTP_Client = new(http.Client)
	transport := &http.Transport{
		MaxIdleConns:        0,
		IdleConnTimeout:     30 * time.Second,
		DisableCompression:  true,
		MaxIdleConnsPerHost: 500,
	}
	HTTP_Client.Transport = transport
	HTTP_Client.Timeout = time.Second
}

// get
func HttpRequestGet(url string, params map[string]string) (bodyJson *simplejson.Json) {
	if url == "" {
		return nil
	}
	//fmt.Println(HTTP_Client)
	//if HTTP_Client == nil{
	//	newHttpClient()
	//}
	//fmt.Println(HTTP_Client)
	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v get url[%s] params[%v]", r, url, params)
		}
	}()
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.LOGGER.Error("new get request failed %s", err.Error())
	}
	// 添加参数
	query := request.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
	//httpClient := new(http.Client)
	//httpClient.Timeout = time.Second *5
	response, err := HTTP_Client.Do(request)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.LOGGER.Error("couldn't parse response body %s", err.Error())
		return nil
	}
	log.LOGGER.Info("request get url[%s] data[%v] recv response status[%s] body %s",
		url, params, response.Status, string(body))
	// json解析
	result, err := simplejson.NewJson(body)
	if err != nil {
		log.LOGGER.Info("%s: %s[%s]", "failed to JSON data.Body", err.Error(), body)
		return nil
	}
	return result
}

// post json
func HttpRequestPost(url string, data map[string]string) (bodyJson *simplejson.Json) {
	if url == "" {
		return nil
	}
	//if HTTP_Client == nil {
	//	newHttpClient()
	//}
	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v post url[%s] data[%v]", r, url, data)
		}
	}()
	// 添加参数
	//postValue := Url.Values{}
	//for key,value := range data{
	//	postValue.Set(key,value)
	//}
	//postString := postValue.Encode()
	jsonData, _ := json.Marshal(data)
	request, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		log.LOGGER.Error("new post request failed %s", err.Error())
	}
	request.Header.Add("Content-Type", "application/json")
	//httpClient := new(http.Client)
	//httpClient.Timeout = time.Second *5
	response, err := HTTP_Client.Do(request)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.LOGGER.Error("couldn't parse response body %s", err.Error())
	}
	log.LOGGER.Info("request post url[%s] data[%v] recv response status[%s] body %s",
		url, data, response.Status, string(body))
	// json解析
	result, err := simplejson.NewJson(body)
	if err != nil {
		log.LOGGER.Info("%s: %s[%s]", "failed to JSON data.Body", err.Error(), body)
		return nil
	}
	return result
}


// post formData
func HttpPostForm(url string, data map[string]string) (bodyJson *simplejson.Json) {
	if url == "" {
		return nil
	}
	defer func() {
		if r := recover(); r != nil {
			log.LOGGER.Error("%v post url[%s] data[%v]", r, url, data)
		}
	}()

	// 添加参数
	postValue := Url.Values{}
	for key, value := range data {
		postValue.Set(key, value)
	}
	postString := postValue.Encode()
	request, err := http.NewRequest("POST", url, strings.NewReader(postString))
	if err != nil {
		log.LOGGER.Error("new post request failed %s", err.Error())
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded ")
	response, err := HTTP_Client.Do(request)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.LOGGER.Error("couldn't parse response body %s", err.Error())
	}
	log.LOGGER.Info("request post url[%s] data[%v] recv response status[%s] body %s",
		url, data, response.Status, string(body))
	// json解析
	result, err := simplejson.NewJson(body)
	if err != nil {
		log.LOGGER.Error("%s: %s[%s]", "failed to JSON data.Body", err.Error(), body)
		return nil
	}
	return result
}
