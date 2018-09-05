package validate

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type FacebookConfig struct {
	Token        string
	ClientID     string
	ClientSecret string
	GrantType    string
	RedirectURI  string
	Scope        string
}

type fbQuery struct {
	InputToken  string `url:"input_token"`
	AccessToken string `url:"access_token"`
}

const graphURL = "https://graph.facebook.com"
const debugTokenPath = "debug_token"
const appTokenPath = "oauth/access_token"

//GenerateURL takes in the url path and query string
//parameters and generates a URL out of it.
func generateURL(path string, params interface{}) string {
	v, _ := query.Values(params)
	url := []string{path, v.Encode()}

	return strings.Join(url, "?")
}

func httpGet(url string) (interface{}, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}

	return data, err
}

func appAccessTokenURL() string {
	s := []string{graphURL, appTokenPath}
	return strings.Join(s, "/")
}

func obtainAppAccessToken(config *FacebookConfig) (string, error) {
	url := generateURL(appAccessTokenURL(), config)
	data, err := httpGet(url)
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%s", data), nil
}

func inspectTokenURL(params fbQuery) string {
	s := []string{graphURL, debugTokenPath}
	path := strings.Join(s, "/")
	return generateURL(path, params)
}

func Facebook(config *FacebookConfig) (interface{}, error) {
	appAccessToken, err := obtainAppAccessToken(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(appAccessToken)
	q := fbQuery{
		InputToken:  config.Token,
		AccessToken: appAccessToken,
	}

	inspectURL := inspectTokenURL(q)
	data, err := httpGet(inspectURL)

	return data, err
}
