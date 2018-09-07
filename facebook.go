package validate

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type FacebookConfig struct {
	Token        string `url:"token"`
	ClientID     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
	GrantType    string `url:"grant_type"`
	RedirectURI  string `url:"redirect_uri"`
	Scope        string `url:"scope"`
}

type fbQuery struct {
	InputToken  string `url:"input_token"`
	AccessToken string `url:"access_token"`
}

type accessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type InspectTokenResponse struct {
	Data FacebookAuthenticateSuccessResponse `json:"data"`
}

type FacebookAuthenticateSuccessResponse struct {
	Token       string   `json:"token"`
	AppID       string   `json:"app_id"`
	Type        string   `json:"type"`
	Application string   `json:"application"`
	ExpiresAt   int64    `json:"expires_at"`
	IsValid     bool     `json:"is_valid"`
	Scopes      []string `json:"scopes"`
	UserID      string   `json:"user_id"`
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

func obtainAppAccessToken(config *FacebookConfig) (string, error) {

	url := generateURL(appTokenPath, config)

	r, err := http.Get(url)
	if err != nil {
		return "", nil
	}
	defer r.Body.Close()

	var atr accessTokenResponse
	err = json.NewDecoder(r.Body).Decode(&atr)
	if err != nil {
		return "", nil
	}

	return atr.AccessToken, nil
}

func Facebook(config *FacebookConfig) (FacebookAuthenticateSuccessResponse, error) {
	appAccessToken, err := obtainAppAccessToken(config)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(appAccessToken)
	q := fbQuery{
		InputToken:  config.Token,
		AccessToken: appAccessToken,
	}

	inspectURL := generateURL(debugTokenPath, q)
	res, err := http.Get(inspectURL)
	if err != nil {
		return FacebookAuthenticateSuccessResponse{}, err
	}
	defer res.Body.Close()

	var data InspectTokenResponse
	err = json.NewDecoder(res.Body).Decode(&data)

	return data.Data, err
}
