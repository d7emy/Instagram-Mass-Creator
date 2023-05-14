package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func initalSession(proxy string) (sessions sessions, e error) {
	u, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{Proxy: http.ProxyURL(u)}

	req, err := http.NewRequest("GET", "https://z-p4.www.instagram.com/", nil)
	if err != nil {
		panic(err)
	}
	req.Header = initalHeader()
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return sessions, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sessions, err
	}

	libs := regex(string(body), "ConsumerLibCommons.js/(.*?).js")
	if !(len(libs) > 1 && libs[1] != "") {
		return sessions, fmt.Errorf("no lib")
	}

	libUrl := fmt.Sprintf("https://z-p4.www.instagram.com/static/bundles/es6/ConsumerLibCommons.js/%s.js", libs[1])
	req, err = http.NewRequest("GET", libUrl, nil)
	if err != nil {
		panic(err)
	}
	req.Header = initalHeader()
	resp, err = transport.RoundTrip(req)
	if err != nil {
		return sessions, err
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return sessions, err
	}
	s := regex(string(body), "ASBD_ID='(\\d+)'")
	if !(len(s) > 1 && s[1] != "") {
		return sessions, fmt.Errorf("no reg")
	}
	sessions.asbd_id = s[1]
	s = regex(string(body), "AppId='(\\d+)'")
	if !(len(s) > 1 && s[1] != "") {
		return sessions, fmt.Errorf("no reg")
	}
	sessions.fbapp_id = s[1]

	req, err = http.NewRequest("GET", "https://z-p4.www.instagram.com/", nil)
	if err != nil {
		panic(err)
	}
	req.Header = initalHeader()
	resp, err = transport.RoundTrip(req)
	if err != nil {
		return sessions, err
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return sessions, err
	}
	s = regex(string(body), `"device_id":"(.*?)"`)
	if !(len(s) > 1 && s[1] != "") {
		return sessions, fmt.Errorf("no reg")
	}
	sessions.device_id = s[1]
	sessions.ig_did = s[1]
	s = regex(string(body), `"csrf_token":"(.*?)"`)
	if !(len(s) > 1 && s[1] != "") {
		return sessions, fmt.Errorf("no reg")
	}
	sessions.csrf = s[1]
	s = regex(string(body), `"rollout_hash":"(.*?)"`)
	if !(len(s) > 1 && s[1] != "") {
		return sessions, fmt.Errorf("no reg")
	}
	sessions.rollout_hash = s[1]
	mid := ""
	for _, c := range resp.Cookies() {
		if c.Name == "mid" {
			mid = c.Value
		}
	}
	if mid == "" {
		return sessions, fmt.Errorf("no mid")
	}
	sessions.mid = mid
	return sessions, nil
}
