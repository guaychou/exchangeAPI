package exchangeAPI

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"errors"
)

type Rates map[string]interface{}

type Result struct {
	Rates Rates `json:"rates"`
	Base string `json:"base"`
	Date string `json:"date"`
}

func GetExchange(from string , to string)Result{
	var response Result
	url:="https://api.exchangeratesapi.io/latest?base="+from+"&symbols="+to
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err!=nil{
		   log.Fatal(err)
		}
		jsonErr := json.Unmarshal(body, &response)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		response.Rates[to],err=convertToInt(response, to)
		if err!=nil{
			log.Fatal(err)
		}
		return response
	}
	return response
}

func convertToInt(param Result, to string) (int,error) {
	if param.Rates[to]==0{
		return 0,errors.New("Value not found")
	}
	param.Rates[to]=int(param.Rates[to].(float64))
	return param.Rates[to].(int),nil
}