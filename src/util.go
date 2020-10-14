package main

import (
	"net/http"
	"time"
)

// Pings an http address, returning delay in ms and size in bytes
func Ping(url string) (speed, size int, result error) {
	now := time.Now().UnixNano() / 1000000
	resp, err := http.Get(url)
	speedResult := (time.Now().UnixNano() / 1000000) - now

	if err != nil {
		return 0, 0, err
	}

	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return resp.StatusCode, 0, err
	}

	return int(speedResult), int(resp.ContentLength), nil
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
