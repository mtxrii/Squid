package main

import (
	"net/http"
	"sort"
	"time"
)

func Ping(url string) (speed, size int, result error) {
	now := time.Now().UnixNano() / 1000000
	resp, err := http.Get(url)
	speedResult := (time.Now().UnixNano() / 1000000) - now

	if err != nil {
		resp.Body.Close()
		return 0, 0, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		resp.Body.Close()
		return resp.StatusCode, 0, err
	}

	resp.Body.Close()
	return int(speedResult), int(resp.ContentLength), nil
}

func Median(slc []int) (med int) {
	sort.Ints(slc)
	n := len(slc)
	if n%2 == 0 {
		return (slc[n/2] + slc[(n/2)-1]) / 2
	} else {
		return slc[(n-1)/2]
	}
}
