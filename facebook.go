package validate

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type FacebookConfig struct {
	Token        string `url:"token"`
	ClientID     string `url:"client_id""`
	ClientSecret string `url:"client_secret"`
	GrantType    string `url:"grant_type"`
	RedirectURI  string `url:"redirect_uri"`
	Scope        string `url:"scope"`
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

	s := []string{graphURL, path}
	r := strings.Join(s, "/")
	v, _ := query.Values(params)

	url := []string{r, v.Encode()}

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

func obtainAppAccessToken(config *FacebookConfig) (string, error) {

	url := generateURL(appTokenPath, config)

	data, err := httpGet(url)
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%s", data), nil
}

func Facebook(config *FacebookConfig) (interface{}, error) {
	appAccessToken, err := obtainAppAccessToken(config)

	if err != nil {
		return nil, err
	}

	fmt.Println(appAccessToken)
	q := fbQuery{
		InputToken:  config.Token,
		AccessToken: appAccessToken,
	}

	inspectURL := generateURL(debugTokenPath, q)
	data, err := httpGet(inspectURL)

	return data, err
}
