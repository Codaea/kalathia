package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/exp/rand"
)

type xkcdJson struct {
Month string `json:"month"`
Link string    `json:"link"`
Year string `json:"year"`
News string    `json:"news"`
SafeTitle string `json:"safe_title"`
Transcript string `json:"transcript"`
Alt string      `json:"alt"`
Img string      `json:"img"`
Title string    `json:"title"`
Day string `json:"day"`
}

// getXkcd returns the link to the xkcd comic with the given number. latest is also acceptable. none will give random
func GetXkcd(s string) (string, error) {

	var xkcd xkcdJson
	var resp *http.Response
	var err error
	
	if s == "latest" { // latest xkcd
		resp, err = http.Get("https://xkcd.com/info.0.json")

	} else if s == "" { // random xkcd
		// get days since xkcd epoch
		xkcdEpoch := time.Since(time.Date(2006, 1, 1, 0, 0, 0, 0, time.UTC))
		daysSinceEpoch := xkcdEpoch.Hours() / 24 // days since xkcd epoch (~7000)
		totalishXkcds := daysSinceEpoch * 3/7 // xkcd posts 3 times a week (m,w,f) (3/7)
		
		randomDay := rand.Intn(int(totalishXkcds)) // its called totalish becasue its off by about 57.

		resp, err = http.Get("https://xkcd.com/" + fmt.Sprint(randomDay)+ "/info.0.json")

	} else { // specific xkcd
		resp, err = http.Get("https://xkcd.com/" + s + "/info.0.json")
		
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &xkcd)
	if err != nil {
		return "", err
	}

	img := xkcd.Img

	return img, nil
}
