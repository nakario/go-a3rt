package a3rt

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strings"
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
		return nil, fmt.Errorf("request entity too long: query must not be more than 2048 bytes.")
	}
	values := url.Values{}
	values.Set("apikey", cli.key)
	values.Add("query", query)
	resp, err := http.Post(apiBase + "talk/v1/smalltalk", "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var smalltalk SmalltalkResponse
	err = json.Unmarshal(bytes, &smalltalk)
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
