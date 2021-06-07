package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

type Params struct {
	issuer  string
	account string
	key     string
}

func main() {
	params := Params{
		issuer:  "Example",
		account: "go@example.com",
		key:     "JBSWY3DPEHPK3PXP",
	}
	uri := Uri(params)

	err := qrcode.WriteFile(uri, qrcode.Medium, 256, "./qrcode/qr.png")
	if err != nil {
		fmt.Println("Error: Failed to generate QR code")
	} else {
		Authentication(params.key)
	}
}

func Totp(k string) int {
	var (
		t0 uint64 = 0
		x  uint64 = 30 // default: 30 seconds
	)
	key, err := base32.StdEncoding.DecodeString(k)
	if err != nil {
		return 0
	}

	return hotp(key, timeCounter(t0, x))
}

func timeCounter(t0, x uint64) uint64 {
	return (uint64(time.Now().Unix()) - t0) / x
}

func hotp(k []byte, c uint64) int {
	return truncate(hmacSha1(k, c))
}

func hmacSha1(k []byte, c uint64) []byte {
	cb := make([]byte, 8)
	binary.BigEndian.PutUint64(cb, c)

	mac := hmac.New(sha1.New, k)
	mac.Write(cb)

	return mac.Sum(nil)
}

func truncate(hs []byte) int {
	offset := int(hs[len(hs)-1] & 0x0F)
	p := hs[offset : offset+4]

	return (int(binary.BigEndian.Uint32(p)) & 0x7FFFFFFF) % 1000000
}

func Uri(params Params) string {
	var uri strings.Builder
	uri.WriteString("otpauth://totp/")
	uri.WriteString(params.issuer)
	uri.WriteString(":")
	uri.WriteString(params.account)
	uri.WriteString("?secret=")
	uri.WriteString(params.key)
	uri.WriteString("&issuer=")
	uri.WriteString(params.issuer)

	return uri.String()
}

func Authentication(k string) {
	fmt.Println("Scan the generated QR code")
	for i := 0; ; i++ {
		var input int
		fmt.Print("Enter your One Time Password: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error: Invalid input")
		} else {
			if input == Totp(k) {
				fmt.Println("Authentication successful!")
				break
			} else {
				fmt.Print("Password does not match, ")
				// Limit trials to three
				if i < 2 {
					fmt.Println("please try again")
					continue
				} else {
					fmt.Println("failed to authenticate")
					break
				}
			}
		}
	}
}
