package pkg

import (
	"log"
	"errors"
	"io/ioutil"
	"net/http"
	"encoding/json"
)


// Release struct
type Release struct {
	Name string `json:"name"`
	PreRelease bool `json:"prerelease"`
	PublishedAt  string `json:"published_at"`
	HtmlUrl string `json:"html_url"`
}

// ReleaseList struct
type ReleaseList struct {
	Releases []Release `json:"releases"`
}

// newReleaseList: New ReleaseList from Github API
func newReleaseList(url string)(releases []Release, err error){

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Get URL(%s), failured.", url)
		return 
	}
	defer resp.Body.Close()

	// Read Res ponse Data
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 
	}
	
	// Parese json data to ReleaseList
	err = json.Unmarshal(content, &releases)
	if err != nil{
		return 
	}
	return
}

// getLastRelease: Get Last Release from list
func getLastRelease(rels []Release)(rel Release, err error){
	
	var found = false
	for _, item := range rels{
		fresh, err := compareFresh(item.PublishedAt)
		if err != nil {
			continue
		}
		if fresh{
			rel = item
			found = true
			break
		}
	}
	if !found{
		err = errors.New("Not found fresh Release")
	}
	
	return 
}

func GitHub()(nRels []Release, err error){
	
	for repo, po := range projects{
		url := unionGitHubAPI(repo, po)
		rels, err := newReleaseList(url)
		if err != nil {
			continue
		}
		rel, err := getLastRelease(rels)
		if err != nil {
			continue
		}
		rel.Name =  po + " - " + rel.Name
		nRels = append(nRels, rel)
	}
	return 
}