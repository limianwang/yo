package accessor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const accessTokenAPI = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

type Access struct {
	AppID      string
	AppSecret  string
	httpClient *http.Client
	reset      chan time.Duration
}

// Error is a generic struct that implements error interface
type wError struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

type accessTokenInfo struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// NewAccessWorker ...
func NewAccessWorker(appID, secret string, duration int) {
	log.Println("Starting Accessor Worker...")
	a := &Access{
		AppID:      appID,
		AppSecret:  secret,
		httpClient: http.DefaultClient,
		reset:      make(chan time.Duration),
	}

	a.start(time.Duration(duration) * time.Minute)
}

func (a *Access) start(duration time.Duration) {
NEW_TICKER:
	ticker := time.NewTicker(duration)

	for {
		select {
		case duration = <-a.reset:
			ticker.Stop()
			goto NEW_TICKER
		case <-ticker.C:
			if token, err := a.getToken(); err != nil {
				break
			} else {
				fmt.Println(token)
			}
		}
	}
}

func (a *Access) getToken() (string, error) {
	_url := fmt.Sprintf(accessTokenAPI, a.AppID, a.AppSecret)
	resp, err := a.httpClient.Get(_url)

	if err != nil {
		fmt.Println(err)
	}

	var result struct {
		wError
		accessTokenInfo
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Token, nil

}
