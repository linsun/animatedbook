package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	// "log"
	"net/http"
	"strings"

	"github.com/ivolo/go-giphy"
)

func main() {
	fmt.Println("Hello, playground")

	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://169.60.115.58:8080/lrange/guestbook", nil)
	if err != nil {
		fmt.Printf("err" + err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("unable to do the req " + err.Error())
	}
	defer res.Body.Close()

	body := []string{}
	json.NewDecoder(res.Body).Decode(&body)
	fmt.Printf("body is %s\n", body)

	// obtain the last entry value
	if len(body) > 0 {
		lastEntry := body[len(body)-1]
		fmt.Printf("last entry is %s\n", lastEntry)
		// extract the tone
		start := strings.Index(lastEntry, " : ")
		end := strings.Index(lastEntry, " (")
		if end > start {
			tone := lastEntry[start+3 : end]
			fmt.Printf("tone is %s\n", tone)

			// get giphy url
			url, err := getGiphyU(tone)
			if err == nil {
				fmt.Printf("giphy url is %s\n", url)
			} else {
				fmt.Printf("unable to get the giphy url " + err.Error())
			}
		} else {
			// unable to find tone
		}
	}

}

// generate giphy url based on tone
func getGiphyU(tone string) (string, error) {
	url := "https://giphy.com/embed/3oKIPs1EVbbNZYq7EA"

	// generate giphy based on the value
	if strings.Contains(tone, "Error - unable to detect Tone from the Analyzer service") {
		// return an error giphy
		url = "https://giphy.com/embed/3oKIPs1EVbbNZYq7EA"
	} else {
		c := giphy.New("jyNPNgvgcIYoNWku6wC171WP24Sc5Xl3")
		gifs, err := c.Search("Joy")
		if err == nil {
			fmt.Printf("gif is %s", gifs[0])
			url = gifs[0].EmbedURL
			fmt.Printf("gif embed_url is %s", gifs[0].EmbedURL)
		}
		/*body := GiphyResponse{}
		err = json.NewDecoder(res.Body).Decode(&body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("embed url is %s", body.Datas)
		fmt.Printf("meta is %s", body.Metas)*/
	}
	return url, nil
}
