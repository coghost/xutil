package xutil

import (
	"io"
	"net/http"
)

func GetPublicIp(args ...string) (ip string) {
	/**
		"https://www.ipify.org",
		"http://myexternalip.com",
		"http://api.ident.me",
		"http://whatismyipaddress.com/api",
	**/
	servers := append([]string{
		"https://api.ipify.org?format=text",
		"http://ip.42.pl/raw",
	}, args...)

	for _, svr := range servers {
		v := GetHostPublicIp(svr)
		if v != "" {
			ip = v
			return ip
		}
	}
	return ip
}

func GetHostPublicIp(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(ip)
}
