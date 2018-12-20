package parser

import (
	"regexp"
	"jsj.golangtc/crawler/zhenai/engine"
)

var cityListRex = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte) (engine.ParseResult) {

	result := engine.ParseResult{}
	matches := cityListRex.FindAllSubmatch(contents, -1)

	index := 0;
	for _, m := range matches {
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Request = append(result.Request, engine.Request{Url: string(m[1]),
			ParserFunc: ParseCity,
		})

		if index >= 1 {
			//break
		}
		index++
	}
	return result

}
