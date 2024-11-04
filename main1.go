package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var keyId = "f66a6a0e49ec1dcb0c8ce69a8ea4ead5"

func fetchWather(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	// fmt.Println(url)
	var data struct {
		Main struct {
			Temp float64
		}
		TimeZone int64
		Name     string
		Code     int32
	}

	defer wg.Done()
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, keyId)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error in fetching data", err.Error())
		return err
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(data)
	ch <- fmt.Sprintln("this is city tempreture data of ", city, data)
	return data
}

func main() {
	// fetchWather("mumbai")
	startTime := time.Now()
	cities := []string{"mumbai", "delhi", "hyderabad", "chennai", "ahmedabad", "jaipur"}

	var ch = make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWather(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("Time taken:- ", time.Since(startTime))

}
