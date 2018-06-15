package pkg

import (
	"errors"
	"io/ioutil"
	"net/http"
	"encoding/json"
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

// newReleaseList: New ReleaseList from Github API
func newReleaseList(url string)(releaseList *ReleaseList, err error){

	resp, err := http.Get(url)
	if err != nil {
		return 
	}
	defer resp.Body.Close()

	// Read Res ponse Data
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 
	}
	
	// Parese json data to ReleaseList
	err = json.Unmarshal(content, &releaseList)
	if err != nil{
		return 
	}
	return
}

// getLastRelease: Get Last Release from list
func (r *ReleaseList) getLastRelease()(release *Release, err error){
	var found = false
	for _, item := range r.Releases {
		fresh, err := compareFresh(release.PublishedAt)
		if err != nil{
			continue
		}
		if fresh{
			release = &item
			found = true
			break
		}
	}
	if !found {
		err = errors.New("Not found Fresh Release.")
	}
	return 
}

func GitHub()(rels []Release, err error){
	
	for repo, po := range projects{
		url := unionGitHubAPI(repo, po)
		relList, err := newReleaseList(url)
		if err != nil {
			continue
		}
		rel, err := relList.getLastRelease()
		if err!=nil {
			continue
		}
		rels = append(rels, *rel)
	}
	return 
}