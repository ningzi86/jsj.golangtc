package core

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	lq "github.com/ahmetalpbalkan/go-linq"
	"crypto/md5"
	"encoding/hex"
	"crypto/rand"
	"encoding/base64"
)

func init() {


}

func GetHeader(headers string) http.Header {

	header := make(http.Header)
	temp := strings.Split(headers, "\n")

	for _, v := range temp {

		kv := strings.TrimSpace(v)
		if len(kv) == 0 {
			continue
		}

		temp2 := strings.Split(kv, ":")
		header.Add(temp2[0], temp2[1])
	}

	return header

}

func PrintCookies(cookies []*http.Cookie) {

	for _, c := range cookies {
		fmt.Printf("%s:%s\n", c.Name, c.Value)
	}

}

func AddCookies(cookies []*http.Cookie, name string, value string, domain string) []*http.Cookie {

	exists := lq.From(cookies).Where(func(c interface{}) bool {
		return c.(*http.Cookie).Name == name
	}).Select(func(c interface{}) interface{} {
		return c
	}).First()

	if exists == nil {
		cookie := &http.Cookie{
			Name:   name,
			Value:  value,
			Domain: domain, //"nike.com"
		}
		cookies = append(cookies, cookie)
	}

	return cookies

}

func MergeCookies(oldCookies []*http.Cookie, newCookies []*http.Cookie) []*http.Cookie {

	if len(newCookies) == 0 {
		return oldCookies
	}

	for _, new := range newCookies {

		exists := false
		for _, old := range oldCookies {
			if old.Name == new.Name {
				exists = true
				old.Value = new.Value
				old.Expires = new.Expires
			}
		}

		if !exists {
			oldCookies = append(oldCookies, new)
		}

	}

	return oldCookies

}

func JoinCookies(cookies []*http.Cookie, str []string) []*http.Cookie {

	if len(str) == 0 {
		return cookies
	}

	for _, c := range str {

		if len(c) == 0 {
			continue
		}
		key := c[0:strings.Index(c, "=")]
		value := c[strings.Index(c, "=")+1:]

		if strings.Index(value, ";") > -1 {
			value = value[0:strings.Index(value, ";")]
		}
		exists := false
		for _, old := range cookies {
			if old.Name == key {
				exists = true
				old.Value = value
				break
			}
		}

		if !exists {
			cookies = append(cookies, &http.Cookie{Name: key, Value: value})
		}

	}
	return cookies
}

func WriteCookie(cookies []*http.Cookie) string {

	bytes, err := json.Marshal(cookies)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func ReadCookie(str string) []*http.Cookie {
	var cookies []*http.Cookie
	json.Unmarshal([]byte(str), &cookies)
	return cookies
}

func ToCookies(cookies []*http.Cookie) string {

	cookie := ""
	for _, c := range cookies {
		cookie += fmt.Sprintf("%s=%s;", c.Name, c.Value)
	}

	return cookie
}

func WriteFile(fileName string, wireteString string) string {

	if len(fileName) == 0 {
		fileName = GetGuid() + ".txt"
	}

	pwd, _ := os.Getwd()
	fileName = pwd + "/cookie/" + fileName
	f, _ := os.Create(fileName)
	io.WriteString(f, wireteString)

	return fileName

}

func ReadFile(fileName string) string {

	pwd, _ := os.Getwd()
	fileName = pwd + "/cookie/" + fileName

	bytes, _ := ioutil.ReadFile(fileName)
	return string(bytes)

}

func WriteCookie2(fileName string, cookies []*http.Cookie) {
	str := ToCookies(cookies)
	WriteFile(fileName, str)
}

type Citys struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}


//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}


