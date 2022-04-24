package otp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	DefaultExpiredSecond = 30
	DefaultDigits        = 6
)

func NewOTPAuth() *OTPAuth {
	return &OTPAuth{
		ExpiredSecond: DefaultExpiredSecond,
		Digits:        DefaultDigits,
	}
}

func (g *OTPAuth) un() int64 {
	return time.Now().UnixNano() / 1000 / 30
}

func (g *OTPAuth) hmacSha1(key, data []byte) []byte {
	h := hmac.New(sha1.New, key)
	if total := len(data); total > 0 {
		h.Write(data)
	}
	return h.Sum(nil)
}

func (g *OTPAuth) base32encode(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (g *OTPAuth) base32decode(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}

func (g *OTPAuth) toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func (g *OTPAuth) toUint32(bts []byte) uint32 {
	return (uint32(bts[0]) << 24) + (uint32(bts[1]) << 16) +
		(uint32(bts[2]) << 8) + uint32(bts[3])
}

func (g *OTPAuth) oneTimePassword(key []byte, data []byte) uint32 {
	hash := g.hmacSha1(key, data)
	offset := hash[len(hash)-1] & 0x0F
	hashParts := hash[offset : offset+4]
	hashParts[0] = hashParts[0] & 0x7F
	number := g.toUint32(hashParts)
	return number % 1000000
}

// 获取秘钥
func (g *OTPAuth) GenSecret() {
	var buf bytes.Buffer
	binary.Write(&buf, binary.BigEndian, g.un())
	secret := strings.ToUpper(g.base32encode(g.hmacSha1(buf.Bytes(), nil)))
	g.SecretKey = secret
}

// 获取动态码
func (g *OTPAuth) GenCode(secret string) (string, error) {
	secretUpper := strings.ToUpper(secret)
	secretKey, err := g.base32decode(secretUpper)
	if err != nil {
		return "", err
	}
	number := g.oneTimePassword(secretKey, g.toBytes(time.Now().Unix()/30))
	return fmt.Sprintf("%06d", number), nil
}

// 获取动态码二维码内容
func (g *OTPAuth) GenOtpCode(user, secret string) {
	otpcode := fmt.Sprintf("otpauth://totp/%s?secret=%s", user, secret)
	g.OtpCode = otpcode
}

// 获取动态码二维码图片地址,这里是第三方二维码api
func (g *OTPAuth) GenQrcodeUrl(user, secret string) {
	width := "200"
	height := "200"
	data := url.Values{}
	data.Set("data", g.OtpCode)
	codeurl := "https://api.qrserver.com/v1/create-qr-code/?" + data.Encode() + "&size=" + width + "x" + height + "&ecc=M"
	g.OtpUrl = codeurl
}
