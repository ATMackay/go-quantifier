package volatility

import (
	"fmt"
	"net/http"

	"github.com/ATMackay/go-quantifier/rpc"
)

const ServiceName = "volatility"

type Config struct {
	Port      int    `yaml:"port"`
	Path      string `yaml:"path"`
	Loglevel  string `yaml:"loglevel"`
	LogFormat string `yaml:"logformat"`
	ApiKey    string `yaml:"apikey"`
}

// TODO implement business logic

type HelloResp struct {
}

func Hello(w http.ResponseWriter, req *http.Request) {

	rpc.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("not supported"))

	/*
		resp, err := hello()
		if err != nil {

			return
		}
		rpc.RespondWithJSON(w, http.StatusOK, hello)
	*/
}

/*
func hello() {

}
*/
