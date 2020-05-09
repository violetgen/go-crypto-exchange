package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/metarsit/exchange"
	"gopkg.in/yaml.v2"
)

func main() {
	var (
		key, secret                string
		symbol, startDate, endDate string
		page, pageSize             int
	)

	flag.StringVar(&key, "api", "", "API Key Generated by Crypto Exchange")
	flag.StringVar(&secret, "secret", "", "Secret Key Generated by Crypto Exchange")
	flag.StringVar(&symbol, "symbol", "", "Market mark, ethbtc, See below for details")
	flag.StringVar(&startDate, "startDate", "", "Start date of the trade history")
	flag.StringVar(&endDate, "endDate", "", "End date of the trade history")
	flag.IntVar(&page, "page", 0, "Page number")
	flag.IntVar(&pageSize, "pageSize", 0, "Page Size")
	flag.Parse()

	api, err := exchange.NewUserAPI(key, secret)
	if err != nil {
		log.Fatalf("Unable to create UserAPI Instance: %s", err.Error())
	}

	resp, err := api.AllOrders(symbol, startDate, endDate, page, pageSize)
	if err != nil {
		log.Fatalf("Unable to retrieve Orders: %s", err.Error())
	}

	if resp.Code != "0" {
		log.Fatalf("[%s] API Error %s", resp.Code, resp.Message)
	}

	var data exchange.AllOrders
	json.Unmarshal(*resp.Data, &data)
	yamlFormat, err := yaml.Marshal(data)
	if err != nil {
		log.Fatalf("Enable to parse into YAML format: %s | %v", err.Error(), data)
	}
	fmt.Print(string(yamlFormat))
}