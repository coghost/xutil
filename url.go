package xutil

import (
	"net/url"
	"path"

	"github.com/gookit/goutil/dump"
	"github.com/spf13/cast"
)

func DecodeUrl(baseurl string) url.Values {
	// Refer: https://groups.google.com/g/golang-nuts/c/hbNCHMIA05g?pli=1
	pu, _ := url.Parse(baseurl)
	return pu.Query()
}

func EncodeUrl(baseurl string, kv map[string]interface{}, pathArgs ...string) *url.URL {
	pth := FirstOrDefaultArgs("", pathArgs...)

	pu, _ := url.Parse(baseurl)
	values := pu.Query()

	for k, v := range kv {
		values[k] = cast.ToStringSlice(v)
	}

	pu.RawQuery = values.Encode()
	pu.Path = path.Join(pu.Path, pth)

	dump.P(pu.String())
	return pu
}

func DecodeUrlToMap(baseurl string) (got map[string]string) {
	dat := DecodeUrl(baseurl)

	got = make(map[string]string)

	for k, v := range dat {
		got[k] = v[0]
	}
	return got
}

func DecodeUrlToMapIface(baseurl string) (got map[string]interface{}) {
	dat := DecodeUrl(baseurl)
	got = make(map[string]interface{})
	for k, v := range dat {
		got[k] = v[0]
	}
	return got
}

func GetUrlParam(baseurl string, key string) (bool, string) {
	dat := DecodeUrl(baseurl)
	if v, b := dat[key]; b {
		return true, v[0]
	}
	return false, ""
}
