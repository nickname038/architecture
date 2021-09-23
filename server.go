package main

import (
	"encoding/json"
	"fmt"
	. "net/http"
	"time"
)

type CurrentTimeResponse struct {
	Time string `json:"time"`
}

const layoutRFC3339 = "2006-01-02T15:04:05Z07:00"

func currentTime(response ResponseWriter, request *Request) {
	if request.Method == "GET" {
		currentTime := time.Now().Format(layoutRFC3339)
		body, _ := json.MarshalIndent(CurrentTimeResponse{
			Time: currentTime,
		}, "", "  ")
		response.Write(body)
	}
}

func main() {
	HandleFunc("/time", currentTime)
	if err := ListenAndServe(":8795", nil); err != nil {
		fmt.Println(err)
	}
}
