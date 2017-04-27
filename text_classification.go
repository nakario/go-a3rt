package a3rt

import (
	"fmt"
	"net/url"
)

const maxTextSize = 1000

type ClassifyResponse struct{
	Status	int	`json:"status"`
	Message	string	`json:"message"`
	Classes	[]Class	`json:"suggestion"`
}

type Class struct{
	Label		string	`json:"label"`
	Probability	float64	`json:"probability"`
}

func (cli Client) Classify(text string, modelID string) ([]Class, error) {
	if len(text) > maxTextSize {
		return nil, fmt.Errorf("request entity too long: text must not be more than %d characters.", maxTextSize)
	}

	values := url.Values{}
	values.Set("apikey", cli.key)
	values.Set("model_id", modelID)
	values.Add("text", text)
	var classify ClassifyResponse
	err := cli.get("https://api.a3rt.recruit-tech.co.jp/text_classification/v1/classify", values, &classify)
	if err != nil {
		return nil, err
	}

	switch classify.Status {
	case 0:
		return classify.Classes, nil
	default:
		return nil, fmt.Errorf(classify.Message)
	}
}
