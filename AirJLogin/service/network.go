package service

import (
	"fmt"
	"net"
	"strings"
)

func GetMyIP() (string, error) {
	conn, err := net.Dial("udp", "210.28.18.6:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

func GetMyMAC(ip string) (string, error) {
	if len(ip) == 0 {
		return "", fmt.Errorf("IP addr invalid")
	}
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			if strings.Contains(addr.String(), ip) {
				//log.Printf("mac: %v -- addr: %v", i.HardwareAddr, addr)
				mac := strings.Replace(i.HardwareAddr.String(), ":", "-", -1)
				return mac, nil
			}
		}
	}
	return "", fmt.Errorf("未找到此IP对应的MAC地址")
}
