package utils

import (
	"net"
	"net/http"
	"time"
)

// see https://stackoverflow.com/questions/26567529/go-to-test-website-status-ping
var client = http.Client{
	Transport: &http.Transport{
		Dial: (&net.Dialer{Timeout: 2 * time.Second}).Dial,
	},
}

/*
Returns the status code of a domain. 
*/
func Ping(domain string) (int, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}