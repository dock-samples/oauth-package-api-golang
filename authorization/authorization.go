package authorization

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	client       ClientHTTPInterface = &http.Client{Timeout: 60 * time.Second}
	Homologation                     = "homologation"
	Production                       = "production"
)

func New(username, password, environment string) *Authorization {
	url := "https://auth.hml.caradhras.io/oauth2/token?grant_type=client_credentials"
	if environment == Production {
		url = "https://auth.caradhras.io/oauth2/token?grant_type=client_credentials"
	}
	return &Authorization{
		username: username,
		password: password,
		url:      url,
	}
}

func (a *Authorization) GetAccessToken() (string, error) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.IsExpired() {
		oauth, err := a.oauth()
		if err != nil {
			return "", err
		}
		a.accessToken = oauth.AccessToken
		a.expiresIn = time.Now().Add(time.Duration(oauth.ExpiresIn-10) * time.Second)
	}
	return a.accessToken, nil
}

func (a *Authorization) oauth() (oauthResponse, error) {

	var response oauthResponse

	basicAuth := base64.StdEncoding.EncodeToString([]byte(a.username + ":" + a.password))

	request, err := http.NewRequest(http.MethodPost, a.url, nil)
	if err != nil {
		return response, err
	}
	request.Header.Add("Authorization", "Basic "+basicAuth)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(request)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != http.StatusOK {
		return response, errors.New(response.Error)
	}

	return response, nil
}

func (a *Authorization) IsExpired() bool {
	return time.Now().After(a.expiresIn)
}

func (a *Authorization) ExpireAccessToken() {
	a.expiresIn = time.Time{}
}

type ClientHTTPInterface interface {
	Do(req *http.Request) (*http.Response, error)
}

type Authorization struct {
	username    string
	password    string
	url         string
	accessToken string
	expiresIn   time.Time
	lock        sync.Mutex
}

type oauthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Error       string `json:"error"`
}
