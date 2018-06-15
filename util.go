package main

import (
	"strconv"
	"time"
	"net/http"
)

func GetTimeStamp() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}
func InitHttpClientNoKeepAlive() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableKeepAlives :true,
	}
	return &http.Client{Transport: tr}
}

func InitHttpClientKeepAlive() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableKeepAlives :false,
	}
	return &http.Client{Transport: tr}
}

type HttpResponse struct {
	Data string `json:"data"`
	Error string `json:"error"`
	Status string `json:"status"`
}



