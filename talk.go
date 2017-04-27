package a3rt

import (
	"fmt"
	"net/url"
)

const maxBodySize = 2048

type SmalltalkResponse struct{
	Status	int			`json:"status"`
	Message	string			`json:"message"`
	Results	[]SmalltalkResult	`json:"results"`
}

type SmalltalkResult struct{
	Perplexity	float32	`json:"perplexity"`
	Reply		string	`json:"reply"`
}

func (cli Client) SmallTalk(query string) ([]SmalltalkResult, error) {
	if len([]byte(query)) > maxBodySize {
		return nil, fmt.Errorf("request entity too long: query must not be more than %d bytes.", maxBodySize)
	}

	values := url.Values{}
	values.Set("apikey", cli.key)
	values.Add("query", query)
	var smalltalk SmalltalkResponse
	err := cli.post("https://api.a3rt.recruit-tech.co.jp/talk/v1/smalltalk", values, &smalltalk)
	if err != nil {
		return nil, err
	}

	switch smalltalk.Status {
	case 0:
		return smalltalk.Results, nil
	default:
		return nil, fmt.Errorf(smalltalk.Message)
	}
}
