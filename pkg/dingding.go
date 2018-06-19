package pkg

import (
	"bytes"
	"strings"
	"net/http"
	"encoding/json"
	"text/template"
)


type Message struct {
	Type string `json:"msgtype"`
	Markdown string `json:"markdown"`
}


func ding(token string, rels []Release)(err error){
	var b bytes.Buffer
	var relList ReleaseList
	relList.Releases = rels

	var letter = template.New("letter")
	letter, err= letter.Parse(`
		## Open Source Release Letter \n
		{{ range .Releases }}
		### {{ .Name }}\n
		> PreRelease: {{ if .PreRelease }}True{{ else }}False{{ end }} |  {{ .PublishedAt }} \n
		{{ .Body }}\n
		{{ end }}
		`)
	
	if err != nil {return}
	
	err = letter.Execute(&b, relList)
	if err != nil {return}
	
	msg := &Message{
		Type: "markdown",
		Markdown: b.String(),
	}
	body, err := json.Marshal(msg)
	if err != nil{return}

	br := strings.NewReader(string(body))
	url := unionDingDingAPI(token)
	_, err = http.Post(url, "Content-Type: application/json", br)	
	
	return 
}
