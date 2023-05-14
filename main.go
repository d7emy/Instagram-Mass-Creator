package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

var (
	Proxy string
)

func main() {
	args := os.Args
	if len(args) < 1 {
		fmt.Println("Please, include a proxy in the arguments. Example: go run . http://127.0.0.1:8080")
		return
	}
	Proxy = args[1]
	if _, err := url.Parse(Proxy); err != nil {
		fmt.Println("Invalid proxy format. Please enter a valid proxy address in the format: http://ip:port or http://user:pass@ip:port. For example: http://127.0.0.1:8080")
		return
	}
	t := 0
	fmt.Print("Creator Bots Counts (default=10) ? ")
	fmt.Scan(&t)
	if t == 0 {
		t = 10
	}
	for i := 0; i < t; i++ {
		go createAccount()
	}
	select {}
}

func createAccount() {
again:
	email := getEseEmail()
	proxy := Proxy
	session, err := initalSession(proxy)
	fmt.Printf("%+v", session)
	if err == nil {
		person := getUser(proxy)
		username := <-rander
		password := fmt.Sprintf("%s@%s", <-rander, (<-rander)[:5])
		first_name := person.Results[0].Name.First + " " + person.Results[0].Name.Last
		if createAttempt(session, username, email, password, first_name, proxy) {
			if verifyEmail(session, email, proxy) {
				code := ""
				for {
					code = getEseMailCode(email)
					if code != "" && code != "err" {
						break
					}
				}
				fmt.Println(code)
				fmt.Println(username)
				if signUpCode := verifyCode(session, email, proxy, code); signUpCode != "" {
					resp, coo := (create(session, username, email, password, first_name, proxy, signUpCode))
					//checking the response from any  100% not create success response
					//cuz the else of the response will return from the creation endpoint, will be succes.
					//dont worry it the response is "Oops, an error occurred" sometime it work
					if resp != "err" && !strings.Contains(resp, "The IP address you are using has been flagged as an open proxy") {
						cookies := map[string]string{}
						for _, c := range coo {
							cookies[c.Name] = c.Value
						}
						result := map[string]interface{}{
							"username":   username,
							"password":   password,
							"first_name": first_name,
							"resp":       resp,
							"cookies":    cookies,
						}
						data, _ := json.Marshal(result)
						AppendText("created.txt", string(data)+"\n")
					}
				} else {
					goto again
				}
			} else {
				goto again
			}
		} else {
			goto again
		}
	} else {
		goto again
	}

}

func create(session sessions, username, email, password, first_name, proxy, signUpCode string) (string, []*http.Cookie) {
	u, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{Proxy: http.ProxyURL(u)}
	params := url.Values{
		"enc_password":           {"#PWD_INSTAGRAM_BROWSER:0:0:" + password},
		"email":                  {email},
		"username":               {username},
		"first_name":             {first_name},
		"month":                  {"3"},
		"day":                    {"4"},
		"year":                   {"1986"},
		"client_id":              {session.mid},
		"seamless_login_enabled": {"1"},
		"tos_version":            {"row"},
		"force_sign_up_code":     {signUpCode},
	}
	req, err := http.NewRequest("POST", "https://www.instagram.com/accounts/web_create_ajax/", strings.NewReader(params.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header = initalHeader(session)
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return "err", nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(body), resp.Cookies()
}

func verifyCode(session sessions, email, proxy, code string) string {
	u, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{Proxy: http.ProxyURL(u)}
	params := url.Values{
		"email":     {email},
		"device_id": {session.mid},
		"code":      {code},
	}
	req, err := http.NewRequest("POST", "https://i.instagram.com/api/v1/accounts/check_confirmation_code/", strings.NewReader(params.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header = initalHeader(session)
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return gjson.Get(string(body), "signup_code").String()
}

func verifyEmail(session sessions, email, proxy string) bool {
	u, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{Proxy: http.ProxyURL(u)}
	params := url.Values{
		"email":     {email},
		"device_id": {session.mid},
	}
	req, err := http.NewRequest("POST", "https://i.instagram.com/api/v1/accounts/send_verify_email/", strings.NewReader(params.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header = initalHeader(session)
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	return gjson.Get(string(body), "email_sent").Bool()
}

func createAttempt(session sessions, username, email, password, first_name, proxy string) bool {
	u, err := url.Parse(proxy)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{Proxy: http.ProxyURL(u)}
	params := url.Values{
		"enc_password":           {"#PWD_INSTAGRAM_BROWSER:0:0:" + password},
		"email":                  {email},
		"username":               {username},
		"first_name":             {first_name},
		"client_id":              {session.mid},
		"seamless_login_enabled": {"1"},
		"opt_into_one_tap":       {"false"},
	}
	req, err := http.NewRequest("POST", "https://z-p4.www.instagram.com/accounts/web_create_ajax/attempt/", strings.NewReader(params.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header = initalHeader(session)
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	fmt.Println(string(body))
	return gjson.Get(string(body), "dryrun_passed").Bool()
}
func AppendText(path, text string) {
	for {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err == nil {
			f.WriteString(text)
			defer f.Close()
			return
		}
	}

}
