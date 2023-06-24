package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func GetCatFact() {
	url := "https://catfact.ninja/fact"
	catFact := &CatFact{}

	err := GetJson(url, catFact)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Print(catFact.Fact)
	}
}

func GetJson(url string, target interface{}) error {
	response, err := client.Get(url)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}

func ConvertStructToJson(catFacStruct interface{}) (catFactJson []byte) {

	var err error
	catFactJson, err = json.Marshal(catFacStruct)
	if err != nil {
		fmt.Print("Error Marshaling", catFacStruct)
	}

	return
}

func ConvertJsonToStruct(carFactJson []byte, carFactType interface{}) interface{} {
	err := json.Unmarshal(carFactJson, carFactType)
	if err != nil {
		fmt.Print("Error Unmarshaling", string(carFactJson))
	} else {
		fmt.Print()
	}
	return carFactType
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	//GetCatFact()
	catFacStruct := CatFact{"The Maine Coone is the only native American long haired breed.", 10}
	catFactFromStructToJson := ConvertStructToJson(catFacStruct)
	fmt.Println(string(catFactFromStructToJson))

	catFactFromJsonToStruct := ConvertJsonToStruct(catFactFromStructToJson, &CatFact{})
	fmt.Println(catFactFromJsonToStruct)
}
