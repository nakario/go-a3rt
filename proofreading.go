package a3rt

import (
	"fmt"
	"net/url"
)

const maxSentenceSize = 500

type Sensitivity string
const (
	Low	Sensitivity	= "low"
	Medium	Sensitivity	= "medium"
	High	Sensitivity	= "high"
)

type TypoResponse struct{
	Status	int	`json:"status"`
	Message	string	`json:"message"`
	Alerts	[]Alert	`json:"alerts"`
}

type Alert struct{
	CheckedSentence	string	`json:"checkedSentence"`
	AlertCode	int	`json:"alertCode"`
	AlertDetail	string	`json:"alertDetail"`
	Word		string	`json:"word"`
	RankingScore	int	`json:"RankingScore"`
}

func (cli Client) Typo(sentence string, sensitivity Sensitivity) ([]Alert, error) {
	if len(sentence) > maxSentenceSize {
		return nil, fmt.Errorf("request entity too long: sentence must not be more than %d characters.", maxSentenceSize)
	}

	values := url.Values{}
	values.Set("apikey", cli.key)
	values.Add("sentence", sentence)
	values.Set("sensitivity", string(sensitivity))
	var typo TypoResponse
	err := cli.get("https://api.a3rt.recruit-tech.co.jp/proofreading/v1/typo", values, &typo)
	if err != nil {
		return nil, err
	}

	switch typo.Status {
	case 0:
		return typo.Alerts, nil
	default:
		return nil, fmt.Errorf(typo.Message)
	}
}
