package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

// Pings an http address, returning delay in ms and size in bytes
func Ping(url string) (speed, size int, result error) {
	now := time.Now().UnixNano() / 1000000
	resp, err := http.Get(url)
	speedResult := (time.Now().UnixNano() / 1000000) - now

	if err != nil {
		return 404, 0, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return resp.StatusCode, 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		body = []byte{}
	}

	resp.Body.Close()
	return int(speedResult), len(body), nil
}

// Takes an int slice and returns the median (as a single element slice)
func Median(slc []int) (mdn []int) {
	n := len(slc)
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return slc
	}
	if n == 2 {
		return []int{(slc[0] + slc[1]) / 2}
	}

	return Median(slc[1 : n-1])
}
