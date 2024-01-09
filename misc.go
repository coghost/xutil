package xutil

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

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

func GetenvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEnvVarEmpty
	}
	return v, nil
}

func GetenvInt(key string) (int, error) {
	s, err := GetenvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func GetenvBool(key string) (bool, error) {
	s, err := GetenvStr(key)
	if err != nil {
		return false, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// Duration
//
//	usage: `xutil.Duration(xutil.Track("doing sth"))`
func Duration(msg string, start time.Time) {
	log.Info().Msgf("%v: %v\n", msg, time.Since(start))
}
