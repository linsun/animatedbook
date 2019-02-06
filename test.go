package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
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
			url, err := getGiphyURL(tone)
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
func getGiphyURL(tone string) (string, error) {
	url := "https://giphy.com/embed/3oKIPs1EVbbNZYq7EA"
	apiKey := "jyNPNgvgcIYoNWku6wC171WP24Sc5Xl3"

	// generate a seed for the random number
	rand.Seed(time.Now().UnixNano())
	offset := strconv.Itoa(rand.Intn(25))
	fmt.Printf("offset is %s", offset)
	// generate giphy based on the value
	if strings.Contains(tone, "Error - unable to detect Tone from the Analyzer service") {
		// return an error giphy
		url = "https://giphy.com/embed/3oKIPs1EVbbNZYq7EA"
	} else {
		client := &http.Client{}
		reqURL := "http://api.giphy.com/v1/gifs/search?q=" + tone + "&api_key=" + apiKey + "&limit=1&offset=" + offset
		fmt.Printf("request url is %s", reqURL)
		req, err := http.NewRequest("GET", reqURL, nil)
		if err != nil {
			fmt.Printf("create new request err" + err.Error())
		}

		res, err := client.Do(req)
		if err != nil {
			fmt.Printf("unable to do the req " + err.Error())
		}
		defer res.Body.Close()

		fmt.Printf("decoding response \n")
		bodyBytes, err2 := ioutil.ReadAll(res.Body)
		fmt.Printf("readall" + "\n")

		if err2 == nil {
			bodyString := string(bodyBytes)
			fmt.Printf(bodyString + "\n")
			// bodyString = `{"data":[{"type":"gif","id":"3o85xmfmpiRpVYLrb2","slug":"20thcenturyfox-jennifer-lawrence-joy-movie-3o85xmfmpiRpVYLrb2","url":"https:\/\/giphy.com\/gifs\/20thcenturyfox-jennifer-lawrence-joy-movie-3o85xmfmpiRpVYLrb2","bitly_gif_url":"https:\/\/gph.is\/1Hv2HCs","bitly_url":"https:\/\/gph.is\/1Hv2HCs","embed_url":"https:\/\/giphy.com\/embed\/3o85xmfmpiRpVYLrb2","username":"20thcenturyfox","source":"","rating":"pg","content_url":"","source_tld":"","source_post_url":"","is_sticker":0,"import_datetime":"2015-11-10 17:57:28","trending_datetime":"2016-02-29 09:39:14","user":{"avatar_url":"https:\/\/media4.giphy.com\/avatars\/20thcenturyfox\/ABLNdS6yfqe6.gif","banner_url":"https:\/\/media4.giphy.com\/avatars\/20thcenturyfox\/dT8B9SfZkXxp.gif","banner_image":"https:\/\/media4.giphy.com\/avatars\/20thcenturyfox\/dT8B9SfZkXxp.gif","profile_url":"https:\/\/giphy.com\/20thcenturyfox\/","username":"20thcenturyfox","display_name":"20th Century Fox","is_verified":true},"images":{"fixed_height_still":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200_s.gif","width":"379","height":"200"},"original_still":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy_s.gif","width":"500","height":"264"},"fixed_width":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200w.gif","width":"200","height":"106","size":"189425","mp4":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200w.mp4","mp4_size":"18910","webp":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200w.webp","webp_size":"75170"},"fixed_height_small_still":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/100_s.gif","width":"189","height":"100"},"fixed_height_downsampled":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200_d.gif","width":"379","height":"200","size":"331922","webp":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200_d.webp","webp_size":"97746"},"preview":{"width":"348","height":"182","mp4":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy-preview.mp4","mp4_size":"44987"},"fixed_height_small":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/100.gif","width":"189","height":"100","size":"170329","mp4":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/100.mp4","mp4_size":"17215","webp":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/100.webp","webp_size":"67758"},"downsized_still":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy-downsized_s.gif","width":"500","height":"264","size":"104081"},"downsized":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy-downsized.gif","width":"500","height":"264","size":"1251696"},"downsized_large":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy.gif","width":"500","height":"264","size":"1251696"},"fixed_width_small_still":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/100w_s.gif","width":"100","height":"53"},"preview_webp":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy-preview.webp","width":"176","height":"93","size":"49596"},"fixed_width_still":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200w_s.gif","width":"200","height":"106"},"fixed_width_small":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/100w.gif","width":"100","height":"53","size":"55163","mp4":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/100w.mp4","mp4_size":"8731","webp":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/100w.webp","webp_size":"28020"},"downsized_small":{"width":"500","height":"264","mp4":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy-downsized-small.mp4","mp4_size":"102093"},"fixed_width_downsampled":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200w_d.gif","width":"200","height":"106","size":"96427","webp":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200w_d.webp","webp_size":"38282"},"downsized_medium":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy.gif","width":"500","height":"264","size":"1251696"},"original":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy.gif","width":"500","height":"264","size":"1251696","frames":"12","mp4":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy.mp4","mp4_size":"61440","webp":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy.webp","webp_size":"352770"},"fixed_height":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200.gif","width":"379","height":"200","size":"657901","mp4":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200.mp4","mp4_size":"41496","webp":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/200.webp","webp_size":"196024"},"looping":{"mp4":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy-loop.mp4","mp4_size":"814585"},"original_mp4":{"width":"480","height":"252","mp4":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy.mp4","mp4_size":"61440"},"preview_gif":{"url":"https:\/\/media1.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/giphy-preview.gif","width":"110","height":"58","size":"47105"},"480w_still":{"url":"https:\/\/media4.giphy.com\/media\/3o85xmfmpiRpVYLrb2\/480w_s.jpg","width":"480","height":"253"}},"title":"jennifer lawrence coffee GIF by 20th Century Fox","_score":2300013.2,"analytics":{"onload":{"url":"https:\/\/giphy_analytics.giphy.com\/simple_analytics?response_id=5c59a527304d455245164708&event_type=GIF_SEARCH&gif_id=3o85xmfmpiRpVYLrb2&action_type=SEEN"},"onclick":{"url":"https:\/\/giphy_analytics.giphy.com\/simple_analytics?response_id=5c59a527304d455245164708&event_type=GIF_SEARCH&gif_id=3o85xmfmpiRpVYLrb2&action_type=CLICK"},"onsent":{"url":"https:\/\/giphy_analytics.giphy.com\/simple_analytics?response_id=5c59a527304d455245164708&event_type=GIF_SEARCH&gif_id=3o85xmfmpiRpVYLrb2&action_type=SENT"}}}],"pagination":{"total_count":7681,"count":1,"offset":1},"meta":{"status":200,"msg":"OK","response_id":"5c59a527304d455245164708"}}`
			res := &SearchResponse{}
			err := json.Unmarshal([]byte(bodyString), res)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v\n", res.Data[0].EmbedURL)

			url = res.Data[0].EmbedURL
		} else {
			fmt.Printf("unable to read the response " + err.Error())
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

// SearchResponse is the response on the Giphy API search
type SearchResponse struct {
	Data       []Gif      `json:"data"`
	Meta       Meta       `json:"meta"`
	Pagination Pagination `json:"pagination"`
}

// Meta represents the API responds
type Meta struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// Pagination allows you to paginate
type Pagination struct {
	TotalCount int `json:"total_count"`
	Count      int `json:"count"`
	Offset     int `json:"offset"`
}

// Gif is the standard Giphy Gif object
type Gif struct {
	Type             string           `json:"type"`
	ID               string           `json:"id"`
	URL              string           `json:"url"`
	BitlyGifURL      string           `json:"bitly_gif_url"`
	BitlyURL         string           `json:"bitly_url"`
	EmbedURL         string           `json:"embed_url"`
	Username         string           `json:"username"`
	Source           string           `json:"source"`
	Rating           string           `json:"rating"`
	Caption          string           `json:"caption"`
	ContentURL       string           `json:"content_url"`
	ImportDatetime   string           `json:"import_datetime"`
	TrendingDatetime string           `json:"trending_datetime"`
	Images           map[string]Image `json:"images"`
}

// Image is a specifically sized gif
type Image struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Width    string `json:"width"`
	Height   string `json:"height"`
	Size     string `json:"size"`
	Mp4      string `json:"mp4"`
	Mp4Size  string `json:"mp4_size"`
	Webp     string `json:"webp"`
	WebpSize string `json:"webp_size"`
}
