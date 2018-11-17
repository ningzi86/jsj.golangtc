package parser

import (
	"regexp"
	"jsj.golangtc/crawler/zhenai/engine"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

func ParseCity(contents []byte) (engine.ParseResult) {

	result := engine.ParseResult{}
	matches := cityRe.FindAllSubmatch(contents, -1)

	index := 0

	for _, m := range matches {
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Request = append(result.Request, engine.Request{Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			},
		})

		if index > 5 {
			//break
		}

		index++
	}
	return result

}
