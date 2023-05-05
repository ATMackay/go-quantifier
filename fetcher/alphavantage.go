package fetcher

import (
	"fmt"
	"net/http"

	rpc "github.com/ATMackay/go-quantifier/rpc"
)

// AplhpaVantage API

const (

	// Base URL for all queries
	alphaVantageBaseURL = "https://www.alphavantage.co/query?"

	// request keys
	function = "function"
	symbol   = "symbol"
	interval = "interval"
	apiKey   = "apiKey"
)

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
	apiKey string
	rpc    rpc.Client
}

func NewAlphaFetcher(r rpc.Client, k string) *AlphaFetcher {
	return &AlphaFetcher{apiKey: k, rpc: r}
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

func (p *AlphaFetcher) MakeRequest(requestParams map[string]string) string {
	var request string = alphaVantageBaseURL

	if param, ok := requestParams[function]; ok {
		request += "=" + param
	}
	if param, ok := requestParams[symbol]; ok {
		request += "=" + param
	}
	if param, ok := requestParams[interval]; ok {
		request += "=" + param
	}
	if param, ok := requestParams[apiKey]; ok {
		request += "=" + param
	}
	return request
}

func (p *AlphaFetcher) processResponse(resp interface{}) (interface{}, error) {
	httpResp, ok := resp.(http.Response)
	if !ok {
		return nil, fmt.Errorf("unsupported response type")
	}
	if err := rpc.HandleResponseErr(&httpResp); err != nil {
		return nil, err
	}
	data := StockDataAlpha{}
	if err := rpc.DecodeJSON(httpResp.Body, &data); err != nil {
		return nil, err
	}
	return data, nil
}
