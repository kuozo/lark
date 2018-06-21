package pkg

import (
	"bytes"
	"errors"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"text/template"
)

type Content struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}


type Message struct {
	Type string `json:"msgtype"`
	Markdown Content `json:"markdown"`
}

type Response struct {
	Code int `json:"errcode"`
	Msg  string `json:"errmsg"`
}


func ding(token string, rels []Release)(err error){
	var b bytes.Buffer
	var relList ReleaseList
	relList.Releases = rels

	var letter = template.New("letter")
	letter, err= letter.Parse("## Open Resource Release Letters\n {{ range .Releases }} ### {{ .Name }}\n > PreRelease: {{ .PreRelease }} | {{ .PublishedAt }}\n\n ##### More information, [view detail]({{ .HtmlUrl }})\n{{ end }}")
	if err != nil {return}
	
	err = letter.Execute(&b, relList)
	if err != nil {return}
	
	msg := &Message{
		Type: "markdown",
		Markdown: Content{
			Title: "Open Resource Release Letters",
			Text: b.String(),
		},
	}
	body, err := json.Marshal(msg)
	if err != nil{return}

	br := strings.NewReader(string(body))
	url := unionDingDingAPI(token)
	resp, err := http.Post(url, "application/json", br)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 
	}
	var response Response
	err = json.Unmarshal(content, &response)
	if err != nil {return}
	if response.Code != 0 {
		return errors.New(response.Msg)
	}
	return 
}
