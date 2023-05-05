package fetcher

import (
	"fmt"
	"net/http"

	rpc "github.com/ATMackay/go-quantifier/rpc"
)

type Fetcher interface {
	GetData(request string) (interface{}, error)
	makeRequest(requestParams map[string]string) (string, error)
	processResponse(interface{}) (interface{}, error)
}

func GetData(f Fetcher, params map[string]string) (interface{}, error) {
	endpoint, err := f.makeRequest(params)
	if err != nil {
		return nil, err
	}
	resp, err := f.GetData(endpoint)
	if err != nil {
		return nil, err
	}
	return f.processResponse(resp)
}

// AplhpaVantage API

const alphaVantageBaseURL = "https://www.alphavantage.co/query?"

type StockDataAlpha struct {
	MetaData struct {
		Information string `json:"1. Information"`
		Symbol      string `json:"2. Symbol"`
		LastRefresh string `json:"3. Last Refreshed"`
		OutputSize  string `json:"4. Output Size"`
		TimeZone    string `json:"5. Time Zone"`
	} `json:"Meta Data"`
	TimeSeriesDaily map[string]struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	} `json:"Time Series (Daily)"`
}

type AlphaFetcher struct {
	baseUrl string
	rpc     rpc.Client
}

func NewAlphaFetcher(r rpc.Client) *AlphaFetcher {
	return &AlphaFetcher{baseUrl: alphaVantageBaseURL, rpc: r}
}

func (p *AlphaFetcher) GetData(request string) (interface{}, error) {
	req, err := http.NewRequest("GET", request, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.rpc.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (p *AlphaFetcher) makeRequest(requestParams map[string]string) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (p *AlphaFetcher) processResponse(interface{}) (interface{}, error) {
	return "", fmt.Errorf("not implemented")
}
