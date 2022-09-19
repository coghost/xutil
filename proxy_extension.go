package xutil

import (
	"fmt"
	"path/filepath"
	"strings"
)

// NewChromeExtension will create two files from line, and save to savePath
//   - background.js
//   - manifest.json
//
// line format is:
//   "host:port:username:password:<OTHER>"
func NewChromeExtension(line, savePath string) (string, string, string) {
	proxy_js := `var config = {
  mode: 'fixed_servers',
  rules: {
    singleProxy: {
      scheme: 'http',
      host: '%s',
      port: parseInt(%s),
    },
    bypassList: ['foobar.com'],
  },
}
chrome.proxy.settings.set({ value: config, scope: 'regular' }, function () {})
function callbackFn(details) {
  return {
    authCredentials: {
      username: '%s',
      password: '%s',
    },
  }
}
chrome.webRequest.onAuthRequired.addListener(callbackFn, { urls: ['<all_urls>'] }, ['blocking'])`
	manifest := `{
    "version": "1.0.0",
    "manifest_version": 2,
    "name": "Chrome Proxy",
    "permissions": ["proxy", "tabs", "unlimitedStorage", "storage", "<all_urls>", "webRequest", "webRequestBlocking"],
    "background": {
        "scripts": ["background.js"]
    },
    "minimum_chrome_version": "22.0.0"
}`
	arr := strings.Split(line, ":")
	proxy_js = fmt.Sprintf(proxy_js, arr[0], arr[1], arr[2], arr[3])

	ip := arr[0]
	if strings.Contains(arr[0], "superproxy") {
		ipArr := strings.Split(arr[2], "-")
		ip = ipArr[len(ipArr)-1]
	}
	ip = strings.ReplaceAll(ip, ".", "_")

	baseDir := filepath.Join(savePath, ip)

	bg := filepath.Join(baseDir, "background.js")
	mf := filepath.Join(baseDir, "manifest.json")

	MustWriteFile(bg, proxy_js)
	MustWriteFile(mf, manifest)

	return baseDir, bg, mf
}
