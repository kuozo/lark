package pkg

import (
	"fmt"
	"time"
	"strings"
)

func compareFresh(publish string)(fresh bool, err error){

	publishTime, err := time.Parse(time.RFC3339, publish)
	if err != nil {
		return
	}
	nowTime := time.Now()
	freshTime := nowTime.AddDate(0, 0, -FRESHNESS_DAYS)
	if freshTime.Before(publishTime) {
		fresh = true
	}
	return 
}

func unionGitHubAPI(repo, po string)(url string){
	var b strings.Builder
	fmt.Fprint(&b, GITHUB_API_BASE_URL, repo, "/", po, "/releases")
	return b.String()
}

func unionDingDingAPI(token string)(url string){
	var b strings.Builder
	fmt.Fprint(&b, DINDING_ROBOT_URL, token)
	return b.String()
}