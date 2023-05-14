package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var (
	rander = make(chan string)
)

func init() {
	go randomer()
}

func randomer() {
	for {
		rander <- RandomString(10)
		time.Sleep(1)
	}
}
func initalHeader(m ...sessions) map[string][]string {
	if m != nil {
		m := m[0]
		return map[string][]string{
			"Authority":                 {"z-p4.www.instagram.com"},
			"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
			"Accept-Language":           {"en"},
			"Cache-Control":             {"no-cache"},
			"Pragma":                    {"no-cache"},
			"Sec-CH-UA":                 {`".Not/A)Brand";v="99"},  "Google Chrome";v="103"},  "Chromium";v="103"`},
			"Sec-CH-UA-Mobile":          {"?0"},
			"Content-Type":              {"application/x-www-form-urlencoded"},
			"Sec-CH-UA-Platform":        {`"Windows"`},
			"Sec-Fetch-Dest":            {"document"},
			"Sec-Fetch-Mode":            {"navigate"},
			"Sec-Fetch-Site":            {"none"},
			"Sec-Fetch-User":            {"?1"},
			"Upgrade-Insecure-Requests": {"1"},
			"User-Agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"},
			"X-IG-WWW-Claim":            {"0"},
			"X-Mid":                     {m.mid},
			"X-Instagram-AJAX":          {m.rollout_hash},
			"X-ASBD-ID":                 {m.asbd_id},
			"x-CSRFTOKEN":               {m.csrf},
			"X-IG-App-ID":               {m.fbapp_id},
			"Cookie":                    {fmt.Sprintf("csrftoken=%s; mid=%s; ig_did=%s;", m.csrf, m.mid, m.ig_did)},
		}
	}
	return map[string][]string{
		"Authority":                 {"z-p4.www.instagram.com"},
		"Accept":                    {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"},
		"Accept-Language":           {"en"},
		"Cache-Control":             {"no-cache"},
		"Pragma":                    {"no-cache"},
		"Sec-CH-UA":                 {`".Not/A)Brand";v="99"},  "Google Chrome";v="103"},  "Chromium";v="103"`},
		"Sec-CH-UA-Mobile":          {"?0"},
		"Sec-CH-UA-Platform":        {`"Windows"`},
		"Sec-Fetch-Dest":            {"document"},
		"Sec-Fetch-Mode":            {"navigate"},
		"Sec-Fetch-Site":            {"none"},
		"Sec-Fetch-User":            {"?1"},
		"Upgrade-Insecure-Requests": {"1"},
		"User-Agent":                {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"},
		"X-IG-WWW-Claim":            {"0"},
	}
}

func regex(str, rg string) []string {
	r, err := regexp.Compile(rg)
	if err != nil {
		return nil
	}
	return r.FindStringSubmatch(str)
}
func RandomString(Length int, boolean ...bool) string {
	data := "mznxbcvalskdjfhgpqowieuryt0192837465"
	len := 36
	if boolean != nil {
		len = 62
		data = "ABCDEFGHIJKLMNOPQRSTUVWXYZmznxbcvalskdjfhgpqowieuryt0192837465"
	}
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, Length)
	for i := range b {
		b[i] = data[seededRand.Intn(len)]
	}
	return string(b)
}

func getUser(proxy string) AutoGenerated {
	for {
		u, err := url.Parse(proxy)
		if err != nil {
			panic(err)
		}
		t := &http.Transport{
			Proxy: http.ProxyURL(u),
		}
		resp, err := (&http.Client{Transport: t}).Get("https://randomuser.me/api/")
		if err == nil {
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				a := AutoGenerated{}
				err := json.Unmarshal(body, &a)
				if err == nil {
					return a
				}
			}
		}
	}
}
