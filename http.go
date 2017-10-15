package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Ip string `json:"origin"`
}

func main() {
	myClient := &http.Client{Timeout: 10 * time.Second}
	res, _ := myClient.Get("https://httpbin.org/ip")

	defer res.Body.Close()

	var data Response

	err := json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		log.Fatal(err2)
	}

	fmt.Println(data.Ip)
}
