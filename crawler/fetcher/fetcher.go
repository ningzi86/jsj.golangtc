package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"bufio"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
)

var tm = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {

	<-tm

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Host", "album.zhenai.com")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error:status code %d", resp.StatusCode)
	}
	newReader := bufio.NewReader(resp.Body)
	decoder := determineEncodeing(newReader).NewDecoder()
	reader := transform.NewReader(newReader, decoder)
	return ioutil.ReadAll(reader)

}

func determineEncodeing(reader *bufio.Reader) encoding.Encoding {

	bytes, err := reader.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
