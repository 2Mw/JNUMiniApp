package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

var successCount = 0
var failCount = 0

var xh string

type RetStruct struct {
	Code      int         `json:"code"` // -1 失败， 1 成功
	Msg       string      `json:"msg"`
	Transcode string      `json:"transcode"`
	Dataset   interface{} `json:"dataset"`
}

func examineNetworkEnv() bool {
	log.Println("Starting examine your network environment...")
	conn, err := net.Dial("udp", "210.28.18.6:80")
	if err != nil {
		log.Println("You may not in school inner network, please check it out.")
		return false
	}
	defer conn.Close()
	return true
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newReq() *http.Request {
	now := rand.Int63n(1000) + +time.Now().Unix()*1000

	uri, _ := url.Parse("https://e.jiangnan.edu.cn/jnapp/action/invokeMobile/invoke")
	params := url.Values{
		"inStrParams": {"{\"e_account\":\"" + xh + "\",\"serviceId\":\"1100015\"}"},
		"_":           {fmt.Sprintf("%v", now)},
	}
	uri.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", uri.String(), nil)
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"
	req.Header.Set("User-Agent", ua)
	return req
}

func getOnce(req *http.Request) {
	// Construct parameters
	client := &http.Client{}
	// Start a request
	//log.Println(uri.String())
	rsp, err := client.Do(req)
	if err != nil {
		log.Println("Error occurs:" + err.Error())
		wg.Done()
		return
	}

	// Process request
	if rsp.StatusCode == 200 {
		body, _ := io.ReadAll(rsp.Body)
		//fmt.Println(string(body))
		ret := &RetStruct{}
		err := json.Unmarshal(body, &ret)
		if err != nil {
			log.Printf("%v\n", "JSON format conversion failed: "+string(body))
			wg.Done()
			return
		}
		if ret.Code == 1 {
			mutex.Lock()
			successCount += 1
			mutex.Unlock()
			fmt.Print("1 ")
		} else {
			mutex.Lock()
			failCount += 1
			mutex.Unlock()
			fmt.Print("0 ")
		}
	}

	wg.Done()
}

func StartThreads(threads int) {
	b := examineNetworkEnv()
	if !b {
		return
	} else {
		log.Printf("%v\n", "Network OK")
	}

	fmt.Print("Please Input Your Student Number: ")
	_, _ = fmt.Scanf("%s\n", &xh)
	log.Printf("Starting working for %v, Threads: %v\n", xh, threads)
	if len(xh) != 10 {
		log.Println("Your student number is invalid.")
		_, _ = fmt.Scanf("%s")
		return
	}

	wg.Add(threads)
	// create req group
	log.Println("Creating requests group.")
	reqs := make([]*http.Request, threads)
	for i := 0; i < threads; i++ {
		reqs[i] = newReq()
	}

	log.Printf("Starting threads %v.\n", threads)
	for i := 0; i < threads; i++ {
		go getOnce(reqs[i])
	}
	wg.Wait()
	fmt.Println()
	log.Printf("You have succeed for %.2f GB currency. S/F(%v/%v)\n", float32(successCount)*54, successCount, failCount)
	_, _ = fmt.Scanf("%s")
}
