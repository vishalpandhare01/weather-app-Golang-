package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var keyId = "f66a6a0e49ec1dcb0c8ce69a8ea4ead5"

func fetchWather(city string) interface{} {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, keyId)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error in fetching data", err.Error())
		return err
	}

	// fmt.Println(url)
	var data struct {
		Main struct {
			Temp float64
		}
		TimeZone int64
		Name     string
		Code     int32
	}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(data)
	return data
}

func main() {
	// fetchWather("mumbai")
	startTime := time.Now()
	cities := []string{"mumbai", "delhi", "hydrabad", "chennai", "ahemdabad", "jaipur"}

	for _, city := range cities {
		data := fetchWather(city)
		fmt.Println("This is tempreture data of", city, " : ", data)
	}

	fmt.Println("Time taken:- ", time.Since(startTime))

}
