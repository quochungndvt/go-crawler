package controller

import (
	"fmt"
	"go-crawler/schema"
	"go-crawler/utils"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo"
)

//GetContent main func get content
func GetContent(cc echo.Context) error {
	c := cc.(*schema.CustomContext)
	errRequire := c.DoRequire()
	if errRequire != nil {
		return c.ResponseJson(errRequire, nil)
	}
	var (
		data []byte
		err  error
	)
	if c.Method() == echo.POST {
		var request schema.RequestData
		if err := c.ParseRequest(&request); err != nil {
			return c.ResponseJson(err, nil)
		}
		fmt.Println("request", request)
		var paramsPost string
		header := make(map[string]string)
		cookies := []*http.Cookie{}
		isMobile := false
		isPost := false
		if request.MethodData != "" {
			isPost = true
			if !utils.IsJSON(request.MethodData) {
				params := url.Values{}
				methodPostArr := strings.Split(request.MethodData, "&")
				for _, v := range methodPostArr {
					kv := strings.Split(v, "=")
					if len(kv) > 0 {
						params[string(kv[0])] = []string{string(kv[1])}
					} else {
						params[string(v)] = []string{""}
					}
				}
				paramsPost = params.Encode()
				header["Content-Type"] = "application/x-www-form-urlencoded"
			} else {
				paramsPost = request.MethodData
				header["Content-Type"] = "application/json"
			}
		}
		if request.Cookie != "" {
			if strings.Contains(request.Cookie, "User-Agent") {
				isMobile = true
			}
			cookieArr := strings.Split(request.Cookie, "&")
			for _, v := range cookieArr {
				kv := strings.Split(v, "=")
				if len(kv) > 0 {
					cookies = append(cookies, &http.Cookie{Name: string(kv[0]), Value: string(kv[1])})
				} else {
					cookies = append(cookies, &http.Cookie{Name: string(kv[0]), Value: ""})
				}

			}
		}
		if request.Header != "" {
			headerArr := strings.Split(request.Header, "&")
			for _, v := range headerArr {
				kv := strings.Split(v, "=")
				if len(kv) > 0 {
					header[string(kv[0])] = string(kv[1])
				} else {
					header[string(kv[0])] = ""
				}
			}
		}
		header["User-Agent"] = utils.RandomUserAgent(isMobile)
		if isPost {
			data, err = utils.HttpPOSTWithHeader(request.URL, paramsPost, header, cookies, request.Proxy)
		} else {
			data, err = utils.HttpGetWithHeader(request.URL, header, cookies, request.Proxy)
		}

		fmt.Println(err)
		return c.HTML(http.StatusOK, string(data))
	} else if c.Method() == echo.GET {
		var (
			test       string
			url        string
			methodData string
			cookie     string
			proxy      string
			proxyAuth  string
			proxyType  string
			paramsPost string
			isMobile   bool
		)
		test = c.QueryParam("test")
		url = c.QueryParam("url")
		if url == "" {
			url = `https://www.google.com.vn/`
		}
		methodData = c.QueryParam("method_data")
		cookie = c.QueryParam("cookie")
		proxy = c.QueryParam("proxy")
		proxyAuth = c.QueryParam("proxy_auth")
		if proxyAuth != "" {
			proxy = fmt.Sprintf("%s@%s", proxyAuth, proxy)
		}
		proxyType = c.QueryParam("proxy_type")
		if proxyType == "http" || proxyType == "" {
			proxy = fmt.Sprintf("http://%s", proxy)
		} else if proxyType == "sock5" { //not support yet

		}
		if test == "true" {
			header := make(map[string]string)
			cookies := []*http.Cookie{}

			if !utils.IsJSON(methodData) {
				paramsPost = methodData
				header["Content-Type"] = "application/x-www-form-urlencoded"
			} else {
				paramsPost = methodData
				header["Content-Type"] = "application/json"
			}
			if cookie != "" {
				if strings.Contains(cookie, "User-Agent") {
					isMobile = true
				}
				cookieArr := strings.Split(cookie, "&")
				for _, v := range cookieArr {
					kv := strings.Split(v, "=")
					if len(kv) > 0 {
						cookies = append(cookies, &http.Cookie{Name: string(kv[0]), Value: string(kv[1])})
					} else {
						cookies = append(cookies, &http.Cookie{Name: string(kv[0]), Value: ""})
					}
				}
			}

			header["User-Agent"] = utils.RandomUserAgent(isMobile)
			data, err = utils.HttpPOSTWithHeader(url, paramsPost, header, cookies, proxy)
			fmt.Println(err)
			return c.HTML(http.StatusOK, string(data))
		}
	}

	return nil
}
