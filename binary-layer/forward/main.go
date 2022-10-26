package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func makeAPICall(body []byte) (int, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:6723/async", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	res, resErr := client.Do(req)
	if resErr != nil {
		fmt.Println("Failed to send data to internal API", res.StatusCode, res.Body, resErr)
	}
	return res.StatusCode, resErr
}

func main() {
	args := os.Args[1:]
	makeAPICall([]byte(args[0]))
}
