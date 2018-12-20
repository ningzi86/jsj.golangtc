package parser

import (
	"regexp"
	"jsj.golangtc/crawler/zhenai/engine"
)

var regex = regexp.MustCompile(`<div class="des f-cl" data-v-07a0138b>([^>]+)</div>`)

func ParseProfile(contents []byte, name, url, id string) (engine.ParseResult) {

	result := engine.ParseResult{}
	match := regex.FindSubmatch(contents)

	if match != nil {
		result.Items = append(result.Items, engine.Item{
			Payload: "UserName " + name + ";Infos " + string(match[1]),
			Url:     url,
			Id:      id,
			Type:    "profile",
		})
	}

	return result

}
