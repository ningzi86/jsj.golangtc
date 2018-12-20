package core

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"time"

	"golang.org/x/net/html/charset"
)

type NetClient struct {
	Url    string
	Method string
	Body   string

	Request *http.Request
	Cookies []*http.Cookie

	Response     *http.Response
	ResponseBody string

	TimeOut time.Duration
}

func NewNetClient(url string, method string, headers string, body string, cookies []*http.Cookie, proxy *url.URL, timeout time.Duration) *NetClient {

	sss := &NetClient{Url: url, Method: method, Body: body, TimeOut: timeout}

	_body := strings.NewReader(body)
	req, err := http.NewRequest(sss.Method, sss.Url, _body)

	if err != nil {
		panic(err)
	}

	req.Header = GetHeader(headers)

	if len(cookies) > 0 {
		sss.Cookies = cookies
		req.Header.Add("Cookie", ToCookies(cookies))
	}

	sss.Request = req
	return sss
}

func (sss *NetClient) Do() error {

	client := &http.Client{}

	if sss.TimeOut > 0 {
		client.Timeout = sss.TimeOut
	}

	resp, err := client.Do(sss.Request)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var bodyStr string
	if resp.Header.Get("Content-Encoding") == "gzip" {
		bodyStr = sss.changeCharsetEncodingAutoGzipSupport(resp.Header.Get("Content-Type"), resp.Body)
	} else {
		bodyStr = sss.changeCharsetEncodingAuto(resp.Header.Get("Content-Type"), resp.Body)
	}
	defer resp.Body.Close()

	cookies := resp.Header["Set-Cookie"]
	sss.Cookies = JoinCookies(sss.Cookies, cookies)

	sss.Response = resp
	sss.ResponseBody = bodyStr

	return nil

}

// Charset auto determine. Use golang.org/x/net/html/charset. Get page body and change it to utf-8
func (sss *NetClient) changeCharsetEncodingAuto(contentTypeStr string, sor io.ReadCloser) string {
	var err error
	destReader, err := charset.NewReader(sor, contentTypeStr)

	if err != nil {
		destReader = sor
	}

	var sorbody []byte
	if sorbody, err = ioutil.ReadAll(destReader); err != nil {
		// For gb2312, an error will be returned.
		// Error like: simplifiedchinese: invalid GBK encoding
		// return ""
	}
	//e,name,certain := charset.DetermineEncoding(sorbody,contentTypeStr)
	bodystr := string(sorbody)

	return bodystr
}

func (sss *NetClient) changeCharsetEncodingAutoGzipSupport(contentTypeStr string, sor io.ReadCloser) string {
	var err error
	gzipReader, err := gzip.NewReader(sor)
	if err != nil {
		return ""
	}
	defer gzipReader.Close()
	destReader, err := charset.NewReader(gzipReader, contentTypeStr)

	if err != nil {
		destReader = sor
	}

	var sorbody []byte
	if sorbody, err = ioutil.ReadAll(destReader); err != nil {
		// For gb2312, an error will be returned.
		// Error like: simplifiedchinese: invalid GBK encoding
		// return ""
	}
	//e,name,certain := charset.DetermineEncoding(sorbody,contentTypeStr)
	bodystr := string(sorbody)

	return bodystr
}
