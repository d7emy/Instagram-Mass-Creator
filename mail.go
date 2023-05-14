package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func getEseEmail() string {
	req, err := http.NewRequest("GET", "http://ese.kr/?pb=6549", nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")

	HTTPresp, err := (&http.Client{}).Do(req)
	if err == nil {
		defer HTTPresp.Body.Close()
		body, _ := ioutil.ReadAll(HTTPresp.Body)
		resp := string(body)
		return splitRegex(resp, `<input type="search" name="mailbox" value="`, `"`)
	}
	return "err"
}

func splitRegex(value, a, b string) string {
	ss := strings.Split(value, a)
	if len(ss) > 1 {
		ss2 := strings.Split(ss[1], b)
		if len(ss2) > 0 {
			return ss2[0]
		}
	}
	return ""
}

func getEseMailCode(email string) string {
	req, err := http.NewRequest("POST", "http://ese.kr/", strings.NewReader("mail_id=&mail_mode=text&lang=en&mailbox="+email))
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	HTTPresp, err := (&http.Client{}).Do(req)
	if err == nil {
		defer HTTPresp.Body.Close()
		body, _ := ioutil.ReadAll(HTTPresp.Body)
		resp := string(body)
		return splitRegex(resp, `no-reply@mail.instagram.com</td><td style="font-weight:bold;"><a href="#">`, ` is your Instagram code`)
	}
	return "err"
}
