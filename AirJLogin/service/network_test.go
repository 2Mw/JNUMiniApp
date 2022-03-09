package service

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"testing"
)

func TestGetMyMAC(t *testing.T) {
	ip, _ := GetMyIP()
	fmt.Printf("%v\n", ip)
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		addrs, _ := i.Addrs()

		for _, addr := range addrs {
			if strings.Contains(addr.String(), ip) {
				log.Printf("%v", strings.Replace(i.HardwareAddr.String(), ":", "-", -1))
			}
			//fmt.Printf("%v\n", addr)
		}
		//fmt.Println()
	}
}

func TestGetUser(t *testing.T) {
	ip, _ := GetMyIP()
	mac, _ := GetMyMAC(ip)
	user, err := GetCurrentUser(ip, mac)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", user)
}

func TestUnicode(t *testing.T) {
	s := `{"result":"fail","msg":"\u9519\u8bef\u65e0\u6cd5\u5b9a\u4e49"}`

	//tq := strconv.QuoteToASCII(s)
	//s3, _ := strconv.Unquote(strings.Replace(strconv.Quote(s), `\\u`, `\u`, -1))
	fmt.Printf("%v->%v\n", s, UnicodeStrToEntity(s))
	s1 := "hee 你好"
	fmt.Printf("%v\n", strconv.QuoteToASCII(s1))
	s2 := "\"hee 你好\""
	fmt.Println(strconv.Unquote(s2))
}

func TestMarshal(t *testing.T) {
	//s := `{"result":"fail","msg":"错误无法定义"}`

}
