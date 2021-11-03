package service

import (
	"JNUMiniApp/AirJLogin/params"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func GetCurrency(id string) (cur string) {
	//link := "https://app.jiangnan.edu.cn/jnapp/action/invokeMobile/invoke?inStrParams={\"serviceId\": \"1100002\",\"userid\": \""+id+"\"}"
	//link := "https://httpbin.org/get?inStrParams={\"serviceId\": \"1100002\",\"userid\": \""+id+"\"}"
	uri, _ := url.Parse("https://app.jiangnan.edu.cn/jnapp/action/invokeMobile/invoke")
	//由于Go语言不会对URL自动进行编码，因此需要使用url.values进行编码，否则会出错
	params := url.Values{
		"inStrParams": {"{\"serviceId\": \"1100002\",\"userid\": \"" + id + "\"}"},
	}
	uri.RawQuery = params.Encode()
	client := &http.Client{}
	//fmt.Println(uri.String())
	req, _ := http.NewRequest("GET", uri.String(), nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	//req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	rsp, err := client.Do(req)
	if err != nil {
		log.Printf("Init client error : %v\n", err)
	}
	//fmt.Println(rsp.Status)
	data, _ := io.ReadAll(rsp.Body)
	dec := json.NewDecoder(strings.NewReader(string(data)))
	type T struct {
		Code    int `json:"code"`
		Dataset struct {
			Rows []struct {
				Shengyu string `json:"shengyu"`
				Userid  string `json:"userid"`
			} `json:"rows"`
		} `json:"dataset"`
	}

	var m T
	_ = dec.Decode(&m)
	//fmt.Println(m)
	curr, _ := strconv.ParseFloat(m.Dataset.Rows[0].Shengyu, 64)
	return fmt.Sprintf("%.2f GB", curr/1024)
}

func Login(acc string, pass string) {
	ip := GetMyIP()
	link := fmt.Sprintf("http://210.28.18.6:801/eportal/?c=ACSetting&a=Login&protocol=http:&hostname=210.28.18.6&iTermType=1&mac=00-00-00-00-00-00&ip=%s&enAdvert=0&queryACIP=0&loginMethod=1", ip)
	param := url.Values{ // data
		"DDDDD":  {",0," + acc + ""},
		"upass":  {pass},
		"R1":     {"0"},
		"R2":     {"0"},
		"R6":     {"0"},
		"para":   {"0"},
		"0MKKey": {"123456"},
	}

	client := &http.Client{}

	req, _ := http.NewRequest("POST", link, strings.NewReader(param.Encode()))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rsp, err := client.Do(req)
	if err != nil {
		log.Println("Not connect to inner network," + err.Error())
	}
	defer rsp.Body.Close()
	if rsp.StatusCode == 200 {
		body, _ := io.ReadAll(rsp.Body)
		if strings.Contains(string(body), "Dr.COMWebLoginID_2.htm") {
			if g := strings.Split(rsp.Request.URL.RequestURI(), "ErrorMsg="); len(g) > 1 {
				errMsg, _ := url.QueryUnescape(g[1])
				decodeString, _ := base64.StdEncoding.DecodeString(errMsg)
				for _, status := range params.Status {
					strs := strings.Split(status, "|")
					if strs[1] == string(decodeString) {
						log.Println("Login failed, error msg:" + strs[3])
					}
				}
			} else {
				log.Println("Login failed, Unknown error")
			}
			//_ = ioutil.WriteFile("./temp", body, 0644)
		} else if strings.Contains(string(body), "Dr.COMWebLoginID_3.htm") {
			log.Println("Login successfully")
			log.Println("Remain Currency：", GetCurrency(acc))
		}

	} else {
		log.Println("Network unknown error.")
	}
}

func Logout() {
	link := "http://210.28.18.3/F.htm"
	client := http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Println("Logout error:" + err.Error())
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	rsp, err := client.Do(req)
	if err != nil {
		log.Printf("Error %v\n", err.Error())
	} else {
		if rsp.StatusCode == 200 {
			log.Println("Logout Successfully!")
		} else {
			log.Println(fmt.Sprintf("Network invalid code:%v\n", rsp.StatusCode))
		}
	}

}

func GetMyIP() string {
	log.Println("Examine if in inner network...")
	conn, err := net.Dial("udp", "210.28.18.6:80")
	if err != nil {
		log.Printf("Get network err: %v\n", err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
