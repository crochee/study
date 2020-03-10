package arithmetic

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"strings"
	"time"
	jsoniter "github.com/json-iterator/go"
)

/**
* @ Description:
* @Author:
* @Date: 2020/3/10 17:32
 */
const (
	GET      = "GET"
	POST     = "POST"
	PUT      = "PUT"
	PATCH    = "PATCH"
	DELETE   = "DELETE"
	COPY     = "COPY"
	OPTIONS  = "OPTIONS"
	LINK     = "LINK"
	UNLINK   = "UNLINK"
	PURGE    = "PURGE"
	LOCK     = "LOCK"
	UNLOCK   = "UNLOCK"
	PROPFIND = "PROPFIND"
	VIEW     = "VIEW"
	HEAD     = "HEAD"
)

func Handle(method string, body map[string]interface{}) (result map[string]interface{}) {
	start := time.Now()
	// 获取环境
	env := viper.GetString("params.env")
	defer func() {
		if r := recover(); r != nil {
			relog.EorrorWrite("inner", map[string]interface{}{
				"req":  body,
				"resp": r,
				"t":    time.Now().Format("2006-01-02T15:04:05+08:00"),
				"cost": time.Since(start),
				"env":  env,
			})
		} else {
			relog.InfoWrite("inner", map[string]interface{}{
				"req":  body,
				"resp": result,
				"t":    time.Now().Format("2006-01-02T15:04:05+08:00"),
				"cost": time.Since(start),
				"env":  env,
			})
		}
	}()
	// 处理url
	tempURL, ok := body["url"].(string)
	if !ok {
		panic("url is not right!")
	}
	if !strings.HasPrefix(tempURL, "http://") {
		tempURL = "http://" + tempURL
	}
	client := &http.Client{
		Timeout: 10,
	}
	var (
		rep *http.Request
		err error
	)
	delete(body, "url")
	switch strings.ToUpper(method) {
	case POST:
		bodyData, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(body)
		if err != nil {
			return common.Error(-1, err)
		}
		rep, err = http.NewRequest(POST, tempURL, bytes.NewReader(bodyData))
		if err != nil {
			return common.Error(-1, err)
		}
	case GET:
		rep, err = http.NewRequest(GET, buildGetURL(body), nil)
		if err != nil {
			return common.Error(-1, err)
		}
	case HEAD:
		rep, err = http.NewRequest(HEAD, tempURL, nil)
		if err != nil {
			return common.Error(-1, err)
		}
	}
	rep.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(rep)
	if err != nil {
		return common.Error(-1, err)
	}
	defer resp.Body.Close()
	if err = jsoniter.ConfigCompatibleWithStandardLibrary.NewDecoder(resp.Body).Decode(&result); err != nil {
		return common.Error(-1, err)
	}
	return
}

// 组织url格式
func buildGetURL(data map[string]interface{}) string {
	par := url.Values{}
	for k, v := range data {
		if k == "query" {
			mv, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(v)
			v = string(mv)
		}
		par.Add(k, fmt.Sprintf("%s", v))
	}
	return "?" + par.Encode()
}
