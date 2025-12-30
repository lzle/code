package src

import (
	"github.com/bitly/go-simplejson"
	"go-callin/core"
	"io/ioutil"
	"net/http"
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
	defer func() {
		if r := recover(); r != nil {
			core.LOGGER.Error("%v get url[%s] params[%v]", r, url, params)
		}
	}()
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		core.LOGGER.Error("new get request failed . %+v", err)
	}
	// 添加参数
	query := request.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
	response, err := HTTP_Client.Do(request)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		core.LOGGER.Error("couldn't parse response body. %+v", err)
		return nil
	}
	core.LOGGER.Info("request get url[%s] data[%v] recv response status[%s] body %s",
		url, params, response.Status, string(body))
	// json解析
	result, err := simplejson.NewJson(body)
	if err != nil {
		core.LOGGER.Error("%s: %s[%s]", "failed to JSON data.Body", err.Error(), body)
		return nil
	}
	return result
}
