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
	request.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729; InfoPath.3)")
	//User-Agent: Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729; InfoPath.3)

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
