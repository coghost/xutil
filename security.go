package xutil

import (
	"github.com/golang-module/dongle"
)

// chaos
type chaos struct {
	cipher *dongle.Cipher
}

// NewChaos creates cipher with mode CBC, padding PKCS7
//
//	@return *chaos
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

func (c *chaos) EncryptToRawBytes(plain []byte) []byte {
	return dongle.Encrypt.FromBytes(plain).ByAes(c.cipher).ToRawBytes()
}

func (c *chaos) EncryptToHexBytes(plain []byte) []byte {
	return dongle.Encrypt.FromBytes(plain).ByAes(c.cipher).ToHexBytes()
}

func (c *chaos) DecryptBytes(sec []byte) []byte {
	return dongle.Decrypt.FromHexBytes(sec).ByAes(c.cipher).ToBytes()
}
