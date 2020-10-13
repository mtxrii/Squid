package main

import (
	"net/http"
	"time"
)

func Ping(url string) (speed int, result error) {

	now := time.Now().UnixNano() / 1000000
	resp, err := http.Head(url)
	speedResult := (time.Now().UnixNano() / 1000000) - now

	if err != nil {
		return 0, err
	}

	resp.Body.Close()
	return int(speedResult), nil
}
