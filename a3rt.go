package a3rt

const apiBase string = "https://api.a3rt.recruit-tech.co.jp/"

type Client struct{
	key	string
}

func NewClient(key sting) Client {
	return Client{key }
}
