package main

import (
	"fmt"
	"go-crawler/utils"
	"net/http"
	"testing"
)

type Proxy struct {
	Region string
	DNS    string
	Port   string
}

func TestGetContent(t *testing.T) {
	var (
		url         = "http://batdongsan.com.vn/"
		params_post = ""
		header      map[string]string
		cookies     []*http.Cookie
		proxy_auth  string
	)
	proxies := []Proxy{
		{"VN", "p1.vietpn.co", "1808"},
		{"VN", "p1.vietpn.co", "1808"},
		{"VN", "p2.vietpn.co", "1808"},
		{"VN", "p4.vietpn.co", "1808"},
		{"VN", "p10.vietpn.co", "1808"},
		{"VN", "p14.vietpn.co", "1808"},
		{"VN", "p15.vietpn.co", "1808"},
		{"VN", "s4.vietpn.co", "1808"},
		{"VN", "123.31.45.40", "1808"},
		{"VN", "p20.vietpn.co", "1808"},
		{"VN", "s9.vietpn.co", "1808"},
		{"US", "172.245.249.126", "1808"},
		{"VN", "42.112.30.20", "3100"},
		{"VN", "42.112.30.23", "3101"},
		{"VN", "42.112.30.30", "3102"},
		{"SG", "149.28.133.243", "1808"},
		{"VN", "v1.vietpn.co", "3100"},
		{"VN", "v2.vietpn.co", "3101"},
		{"VN", "v3.vietpn.co", "3102"},
		{"VN", "v4.vietpn.co", "3103"},
		{"VN", "v5.vietpn.co", "3104"},
		{"VN", "v6.vietpn.co", "3105"},
		{"VN", "v7.vietpn.co", "3106"},
		{"VN", "v8.vietpn.co", "3107"},
		{"VN", "v9.vietpn.co", "3108"},
		{"VN", "v10.vietpn.co", "3109"},
		{"VN", "v11.vietpn.co", "3110"},
		{"VN", "v12.vietpn.co", "3111"},
		{"VN", "v13.vietpn.co", "3112"},
		{"VN", "v14.vietpn.co", "3113"},
		{"VN", "v15.vietpn.co", "3114"},
		{"VN", "v16.vietpn.co", "3115"},
		{"VN", "v17.vietpn.co", "3116"},
		{"VN", "v18.vietpn.co", "3117"},
		{"VN", "v19.vietpn.co", "3118"},
		{"VN", "v20.vietpn.co", "3119"},
		{"VN", "v21.vietpn.co", "3120"},
		{"VN", "v22.vietpn.co", "3121"},
		{"VN", "v23.vietpn.co", "3122"},
		{"VN", "v24.vietpn.co", "3123"},
		{"VN", "v25.vietpn.co", "3124"},
		{"VN", "v26.vietpn.co", "3125"},
		{"VN", "v27.vietpn.co", "3126"},
		{"VN", "v28.vietpn.co", "3127"},
		{"VN", "v29.vietpn.co", "3128"},
		{"VN", "v30.vietpn.co", "3129"},
		{"VN", "v31.vietpn.co", "3130"},
		{"VN", "v32.vietpn.co", "3131"},
		{"VN", "v33.vietpn.co", "3100"},
		{"VN", "p16.vietpn.co", "1808"},
		{"VN", "p22.vietpn.co", "1808"},
		{"VN", "p23.vietpn.co", "1808"},
		{"VN", "p3.vietpn.co", "1808"},
		{"VN", "p21.vietpn.co", "1808"},
		{"JP", "s5.vietpn.co", "1808"},
		{"VN", "p6.vietpn.co", "1808"},
		{"VN", "v34.vietpn.co", "1808"},
		{"VN", "v35.vietpn.co", "1808"},
		{"VN", "v36.vietpn.co", "1808"},
		{"VN", "v39.vietpn.co", "1808"},
		{"VN", "v40.vietpn.co", "1808"},
		{"VN", "v41.vietpn.co", "1808"},
		{"VN", "v41.vietpn.co", "1808"},
		{"VN", "125.212.251.75", "1808"},
	}
	proxy_auth = ""
	for i, proxy := range proxies {
		if proxy.Region == "VN" {
			_proxy := fmt.Sprintf("http://%s@%s:%s", proxy_auth, proxy.DNS, proxy.Port)
			data, err := utils.HttpPOSTWithHeader(url, params_post, header, cookies, _proxy)
			if err != nil {
				t.Errorf("HttpPOSTWithHeader using proxy %s have err: %v", _proxy, err)
				return
			}
			t.Logf("using proxy: %s %v", _proxy, string(data))
		}
		if i == 1 {
			break
		}

	}
}
