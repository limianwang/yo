package accessor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const accessTokenAPI = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

type Access struct {
	AppID      string
	AppSecret  string
	httpClient *http.Client
	reset      chan time.Duration

	LastToken struct {
		sync.Mutex
		tokenInfo accessTokenInfo
		timestamp int64
	}
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
func NewAccessWorker(appID, secret string, duration int) *Access {
	log.Println("Starting Accessor Worker...")
	a := &Access{
		AppID:      appID,
		AppSecret:  secret,
		httpClient: http.DefaultClient,
		reset:      make(chan time.Duration),
	}

	go a.start(time.Duration(duration) * time.Second)
	return a
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
	a.LastToken.Lock()
	defer a.LastToken.Unlock()

	/*
		now := time.Now().Unix()

		if n := a.LastToken.timestamp; n <= now && now < n+2 {
			fmt.Println("using cache")

			return a.LastToken.tokenInfo.Token, nil
		}
	*/

	ch := make(chan *httpResponse)

	go fetchToken(a.AppID, a.AppSecret, ch)

	var result struct {
		wError
		accessTokenInfo
	}

	for {
		select {
		case r := <-ch:
			if r.err != nil {
				fmt.Println(r.err)
			} else {
				if err := json.NewDecoder(r.response.Body).Decode(&result); err != nil {
					return "", nil
				}
				return result.Token, nil
			}
		case <-time.After(2 * time.Second):
			fmt.Println("timeout?")
		}
	}

}

type httpResponse struct {
	err      error
	response *http.Response
}

func fetchToken(appID, secret string, c chan *httpResponse) {
	fmt.Println("In go rountine..")
	_url := fmt.Sprintf(accessTokenAPI, appID, secret)
	resp, err := http.DefaultClient.Get(_url)

	c <- &httpResponse{err, resp}
}
