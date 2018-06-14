package pkg

import (
	"net/http"
)


// Release struct
type Release struct {
	Name string `json:"name"`
	PreRelease string `json:"prerelease"`
	PublishedAt  string `json:"published_at"`
	Body string `json:"body"`
}

// ReleaseList struct
type ReleaseList struct {
	Releases []Release `json:"releases"`
}

func get(url string)error{

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	

	return nil
}