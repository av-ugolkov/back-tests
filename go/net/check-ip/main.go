package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
)

var privateCIDRs []*net.IPNet

func init() {
	cirds := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}
	for _, cidr := range cirds {
		_, block, _ := net.ParseCIDR(cidr)
		privateCIDRs = append(privateCIDRs, block)
	}
}

func fromPrivateIP(r *http.Request) (bool, error) {
	remoreIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return false, err
	}

	ip := net.ParseIP(remoreIP)
	if ip == nil {
		return false, errors.New("couldn't parse ip")
	}
	if ip.IsLoopback() {
		return true, nil
	}

	for _, block := range privateCIDRs {
		if block.Contains(ip) {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	fmt.Println(fromPrivateIP(&http.Request{
		RemoteAddr: "192.168.0.10:5000",
	}))
	fmt.Println(fromPrivateIP(&http.Request{
		RemoteAddr: "127.0.0.1:8080",
	}))
	fmt.Println(fromPrivateIP(&http.Request{
		RemoteAddr: "242.168.20.10:443",
	}))
	fmt.Println(fromPrivateIP(&http.Request{
		RemoteAddr: "10.168.20.10:443",
	}))
}
