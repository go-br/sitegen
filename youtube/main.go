package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/avelino/slugify"
)

type youtube struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
				Standard struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"standard"`
				Maxres struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"maxres"`
			} `json:"thumbnails"`
			ChannelTitle string `json:"channelTitle"`
			PlaylistID   string `json:"playlistId"`
			Position     int    `json:"position"`
			ResourceID   struct {
				Kind    string `json:"kind"`
				VideoID string `json:"videoId"`
			} `json:"resourceId"`
		} `json:"snippet"`
		ContentDetails struct {
			VideoID          string    `json:"videoId"`
			VideoPublishedAt time.Time `json:"videoPublishedAt"`
		} `json:"contentDetails"`
	} `json:"items"`
}

func main() {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	yt := youtube{}
	err = json.Unmarshal(data, &yt)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, v := range yt.Items {

		title := splitTitle(v.Snippet.Title)
		outputFileName := "data/" + slugify.Slugify("hangout "+title) + ".md"
		metadata := fmt.Sprintf("+++\ntitle = \"%s\"\ndescription = \"%s\"\ntags = [\"Golang\"]\ndate = \"%s\"\n+++\n",
			title,
			v.Snippet.Title,
			(v.ContentDetails.VideoPublishedAt.Format(time.RFC3339)))

		//fmt.Println(k, v.Snippet.Title)
		//fmt.Println(k, v.Snippet.Description)
		//fmt.Println(k, v.Snippet.PlaylistID)
		//fmt.Println(k, v.ContentDetails.VideoID)
		//fmt.Println(k, v.ContentDetails.VideoPublishedAt)

		metadata += fmt.Sprintf("\n{{< youtube %s >}}\n\n", v.ContentDetails.VideoID)
		metadata += v.Snippet.Description
		//fmt.Println(metadata)

		metadata = strings.Replace(metadata,
			"https://github.com/crgimenes/Go-Hands-On",
			"https://github.com/go-br/estudos", -1)

		err = ioutil.WriteFile(outputFileName, []byte(metadata), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func splitTitle(title string) string {
	s := strings.Split(title, "Hangout")
	title = strings.Replace(s[0], ";", " ", -1)
	title = strings.Replace(title, ",", " ", -1)
	title = strings.Replace(title, "  ", " ", -1)
	title = strings.TrimSpace(title)
	return title
}

func formatDate(date string) string {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatal(err)
	}
	return t.String()
}
