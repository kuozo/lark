package pkg

const (
	// Github API Base url
	GITHUB_API_BASE_URL = "https://api.github.com/repo/"
	//Dingding reboot url
	DINDING_ROBOT_URL = "https://oapi.dingtalk.com/robot/send?access_token="
	// freshness
	FRESHNESS_DAYS = 1
)

// Support projects, key as repo, value as project name.
var projects = map[string]string{
	"kubernetes": "kubernetes",
	"istio": "istio",
}