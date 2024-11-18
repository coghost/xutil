package xutil

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"
	"os/exec"
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

func GetenvIntWithDft(key string, dft int) int {
	val, err := GetenvInt(key)
	if err != nil {
		return dft
	}

	if val == 0 {
		return dft
	}

	return val
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

func GetenvBoolWithDft(key string, dft bool) bool {
	val, err := GetenvBool(key)
	if err != nil {
		return dft
	}

	return val
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

const ShellToUse = "bash"

func RunCommandInBash(command string) (string, string, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}
