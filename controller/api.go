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
		var params_post string
		header := make(map[string]string)
		cookies := []*http.Cookie{}
		is_mobile := false
		is_post := false
		if request.Method_data != "" {
			is_post = true
			if !utils.IsJSON(request.Method_data) {
				params := url.Values{}
				method_post_arr := strings.Split(request.Method_data, "&")
				for _, v := range method_post_arr {
					kv := strings.Split(v, "=")
					if len(kv) > 0 {
						params[string(kv[0])] = []string{string(v[len(kv[0])+1:])}
					} else {
						params[string(v)] = []string{""}
					}
				}
				params_post = params.Encode()
				header["Content-Type"] = "application/x-www-form-urlencoded"
			} else {
				params_post = request.Method_data
				header["Content-Type"] = "application/json"
			}
		}
		if request.Cookie != "" {
			if strings.Contains(request.Cookie, "User-Agent") {
				is_mobile = true
			}
			cookie_arr := strings.Split(request.Cookie, "&")
			for _, v := range cookie_arr {
				kv := strings.Split(v, "=")
				if len(kv) > 0 {
					cookies = append(cookies, &http.Cookie{Name: string(kv[0]), Value: string(v[len(kv[0])+1:])})
				} else {
					cookies = append(cookies, &http.Cookie{Name: string(kv[0]), Value: ""})
				}

			}
		}
		header["User-Agent"] = utils.RandomUserAgent(is_mobile)
		if is_post {
			data, err = utils.HttpPOSTWithHeader(request.Url, params_post, header, cookies, request.Proxy)
		} else {
			data, err = utils.HttpGetWithHeader(request.Url, header, cookies, request.Proxy)
		}

		fmt.Println(err)
		return c.HTML(http.StatusOK, string(data))
	} else if c.Method() == echo.GET {
		var (
			test        string
			url         string
			method_data string
			cookie      string
			proxy       string
			proxy_auth  string
			proxy_type  string
			params_post string
			is_mobile   bool
		)
		test = c.QueryParam("test")
		url = c.QueryParam("url")
		if url == "" {
			url = `http://batdongsan.com.vn/nha-dat-ban`
		}
		method_data = c.QueryParam("method_data")
		cookie = c.QueryParam("cookie")
		proxy = c.QueryParam("proxy")
		proxy_auth = c.QueryParam("proxy_auth")
		if proxy_auth != "" {
			proxy = fmt.Sprintf("%s@%s", proxy_auth, proxy)
		}
		proxy_type = c.QueryParam("proxy_type")
		if proxy_type == "http" || proxy_type == "" {
			proxy = fmt.Sprintf("http://%s", proxy)
		} else if proxy_type == "sock5" { //not support yet

		}
		if test == "true" {
			header := make(map[string]string)
			cookies := []*http.Cookie{}

			if !utils.IsJSON(method_data) {
				params_post = method_data
				header["Content-Type"] = "application/x-www-form-urlencoded"
			} else {
				params_post = method_data
				header["Content-Type"] = "application/json"
			}
			if cookie != "" {
				if strings.Contains(cookie, "User-Agent") {
					is_mobile = true
				}
				cookie_arr := strings.Split(cookie, "&")
				for _, v := range cookie_arr {
					kv := strings.Split(v, "=")
					if len(kv) > 0 {
						cookies = append(cookies, &http.Cookie{Name: string(kv[0]), Value: string(v[len(kv[0])+1:])})
					} else {
						cookies = append(cookies, &http.Cookie{Name: string(kv[0]), Value: ""})
					}
				}
			}
			header["User-Agent"] = utils.RandomUserAgent(is_mobile)
			data, err = utils.HttpPOSTWithHeader(url, params_post, header, cookies, proxy)
			fmt.Println(err)
			return c.HTML(http.StatusOK, string(data))
		}
	}

	return nil
}
