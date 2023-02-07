package _stringUtils

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/jameskeane/bcrypt"
)

func TrimAll(str string) string {
	str = strings.Trim(str, "\n")
	str = strings.TrimSpace(str)

	return str
}

func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func StrInArr(str string, arr []string) bool {
	for _, s := range arr {
		if str == s {
			return true
		}
	}

	return false
}

func RandStr(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var src = rand.NewSource(time.Now().UnixNano())

	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func Base64Decode(str string) string {
	s, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}
	return string(s)
}

func HashPassword(pwd string) string {
	salt, err := bcrypt.Salt(10)
	if err != nil {
		return ""
	}
	hash, err := bcrypt.Hash(pwd, salt)
	if err != nil {
		return ""
	}

	return hash
}

func ParseInt(str string) (ret int) {
	ret, _ = strconv.Atoi(str)
	return
}
func ParseUint(str string) (ret uint) {
	i := ParseInt(str)
	ret = uint(i)
	return
}
