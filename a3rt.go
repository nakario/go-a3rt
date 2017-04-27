package a3rt

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const apiBase string = "https://api.a3rt.recruit-tech.co.jp/"

type Client struct{
	key	string
}

func NewClient(key string) Client {
	return Client{key }
}

func (cli Client) get(endpoint string, values url.Values, resp interface{}) error {
	getResp, err := http.Get(endpoint + "?" + values.Encode())
	if err != nil {
		return err
	}
	defer getResp.Body.Close()

	bytes, err := ioutil.ReadAll(getResp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, resp)
	if err != nil {
		return err
	}

	return nil
}

func (cli Client) post(endpoint string, values url.Values, resp interface{}) error {
	postResp, err := http.PostForm(endpoint, values)
	if err != nil {
		return err
	}
	defer postResp.Body.Close()

	bytes, err := ioutil.ReadAll(postResp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, resp)
	if err != nil {
		return err
	}

	return nil
}
