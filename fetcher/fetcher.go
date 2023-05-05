package fetcher

// Fetcher is the interface for obtaining market data via external API call
type Fetcher interface {
	GetData(request string) (interface{}, error)
	MakeRequest(requestParams map[string]string) string
	processResponse(interface{}) (interface{}, error)
}

// GetData returns market data for the supplied Fetcher and request params
func GetData(f Fetcher, params map[string]string) (interface{}, error) {
	resp, err := f.GetData(f.MakeRequest(params))
	if err != nil {
		return nil, err
	}
	return f.processResponse(resp)
}
