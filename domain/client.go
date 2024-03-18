package domain

import (
	"encoding/json"
	"github.com/mengfengkong/coin/model"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client interface {
	Get(coin string) (string, error)
}

type client struct {
	ApiKey string `json:"apiKey"`
}

var key string

func NewClient(apiKey string) *client {
	return &client{ApiKey: apiKey}
}

func (c *client) Get(coin string) (*model.CoinResp, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("convert", "USD")
	q.Add("symbol", coin)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", c.ApiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)

	res := &model.CoinResp{}
	err = json.Unmarshal(respBody, res)

	return res, err
}
