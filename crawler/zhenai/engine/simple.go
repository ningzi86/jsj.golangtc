package engine

import (
	"jsj.golangtc/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (SimpleEngine) Run(seeds ...Request) {

	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {

		r := requests[0]
		requests = requests[1:]

		result, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, result.Request...)

		for _, item := range result.Items {
			log.Printf("Got item %s", item)
		}

	}
}

func worker(r Request) (ParseResult, error) {

	var bytes []byte
	var err error

	//if strings.HasPrefix(r.Url, "http://album.zhenai.com/u/") {
	//	bytes, err = ioutil.ReadFile(`crawler/zhenai/parser/user.in`)
	//	if err != nil {
	//		return ParseResult{}, err
	//	}
	//} else {
	bytes, err = fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	//}

	parseResult := r.ParserFunc(bytes)
	return parseResult, nil

}
