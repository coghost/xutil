package xutil

import (
	"github.com/golang-module/dongle"
)

// chaos
type chaos struct {
	cipher *dongle.Cipher
}

func NewChaos(key, iv string) *chaos {
	cp := dongle.NewCipher()
	cp.SetMode(dongle.CBC)
	cp.SetPadding(dongle.PKCS7)
	cp.SetKey(key)
	cp.SetIV(iv)

	return &chaos{cipher: cp}
}

func (c *chaos) Encrypt(plain string) string {
	return dongle.Encrypt.FromString(plain).ByAes(c.cipher).ToHexString()
}

func (c *chaos) Decrypt(sec string) string {
	return dongle.Decrypt.FromHexString(sec).ByAes(c.cipher).ToString()
}
