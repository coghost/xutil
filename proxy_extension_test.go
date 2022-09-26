package xutil_test

import (
	"testing"

	"github.com/coghost/xutil"

	"github.com/stretchr/testify/suite"
	"github.com/ungerik/go-dry"
)

type ProxySuite struct {
	suite.Suite
}

func TestProxy(t *testing.T) {
	suite.Run(t, new(ProxySuite))
}

func (s *ProxySuite) SetupSuite() {
}

func (s *ProxySuite) TearDownSuite() {
}

func (s *ProxySuite) Test_01_Luminati() {
	line := "zproxy.lum-superproxy.io:22225:lum-customer-fake-data-ip-111.11.1.1:fakepassword"
	saveTo := "/tmp/nbcc"

	want := `var config = {
  mode: 'fixed_servers',
  rules: {
    singleProxy: {
      scheme: 'http',
      host: 'zproxy.lum-superproxy.io',
      port: parseInt(22225),
    },
    bypassList: ['foobar.com'],
  },
}
chrome.proxy.settings.set({ value: config, scope: 'regular' }, function () {})
function callbackFn(details) {
  return {
    authCredentials: {
      username: 'lum-customer-fake-data-ip-111.11.1.1',
      password: 'fakepassword',
    },
  }
}
chrome.webRequest.onAuthRequired.addListener(callbackFn, { urls: ['<all_urls>'] }, ['blocking'])`

	want_mf := `{
    "version": "1.0.0",
    "manifest_version": 2,
    "name": "Chrome Proxy",
    "permissions": ["proxy", "tabs", "unlimitedStorage", "storage", "<all_urls>", "webRequest", "webRequestBlocking"],
    "background": {
        "scripts": ["background.js"]
    },
    "minimum_chrome_version": "22.0.0"
}`

	_, bg, mf := xutil.NewChromeExtension(line, saveTo)

	got, _ := dry.FileGetString(bg)
	got_mf, _ := dry.FileGetString(mf)

	s.Equal(want, got)
	s.Equal(want_mf, got_mf)
}

func (s *ProxySuite) Test_02_Normal() {
	saveTo := "/tmp/nbcc"

	want := `var config = {
  mode: 'fixed_servers',
  rules: {
    singleProxy: {
      scheme: 'http',
      host: '179.61.134.75',
      port: parseInt(4444),
    },
    bypassList: ['foobar.com'],
  },
}
chrome.proxy.settings.set({ value: config, scope: 'regular' }, function () {})
function callbackFn(details) {
  return {
    authCredentials: {
      username: 'fakeuser',
      password: 'password',
    },
  }
}
chrome.webRequest.onAuthRequired.addListener(callbackFn, { urls: ['<all_urls>'] }, ['blocking'])`

	want_mf := `{
    "version": "1.0.0",
    "manifest_version": 2,
    "name": "Chrome Proxy",
    "permissions": ["proxy", "tabs", "unlimitedStorage", "storage", "<all_urls>", "webRequest", "webRequestBlocking"],
    "background": {
        "scripts": ["background.js"]
    },
    "minimum_chrome_version": "22.0.0"
}`

	line := "179.61.134.75:4444:fakeuser:password"
	_, bg, mf := xutil.NewChromeExtension(line, saveTo)

	got, _ := dry.FileGetString(bg)
	got_mf, _ := dry.FileGetString(mf)

	s.Equal(want, got)
	s.Equal(want_mf, got_mf)
}
