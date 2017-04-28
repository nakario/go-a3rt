package a3rt

import (
	"fmt"
	"net/url"
)

const maxPreviousDescriptionSize = 500

type TextSuggestClient struct{
	client
}

func NewTextSuggestClient(key string) TextSuggestClient {
	return TextSuggestClient{newClient(key)}
}

type Style int
const (
	Gendaibun Style = iota
	Waka
	Go
)

type Separation int
const (
	Word Separation = iota
	Phrase
	Sentence
)

type PredictResponse struct{
	Status		int		`json:"status"`
	Message		string		`json:"message"`
	Suggestion	[]string	`json:"suggestion"`
}

func (cli TextSuggestClient) Predict(previousDescription string, style Style, separation Separation) ([]string, error) {
	if len(previousDescription) > maxPreviousDescriptionSize {
		return nil, fmt.Errorf("request entity too long: previous description must not be more than %d characters.", maxPreviousDescriptionSize)
	}

	values := url.Values{}
	values.Set("apikey", cli.key)
	values.Add("previous_description", previousDescription)
	values.Set("style", string(style))
	values.Set("separation", string(separation))
	var predict PredictResponse
	err := cli.get("https://api.a3rt.recruit-tech.co.jp/text_suggest/v1/predict", values, &predict)
	if err != nil {
		return nil, err
	}

	switch predict.Status {
	case 0:
		return predict.Suggestion, nil
	default:
		return nil, fmt.Errorf(predict.Message)
	}
}
