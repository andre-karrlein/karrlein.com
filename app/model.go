package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type experience struct {
	Role      string `json:"role"`
	Company   string `json:"company"`
	City      string `json:"city"`
	Timeframe string `json:"timeframe"`
}

func getData() []experience {
	resp, err := http.Get("https://api.karrlein.com/resume/v1/experience/")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	var experiences []experience
	json.Unmarshal([]byte(sb), &experiences)

	return experiences
}