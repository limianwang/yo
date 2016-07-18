package accessor

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const accessTokenAPI = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

type Access struct {
	AppID      string
	AppSecret  string
	httpClient *http.Client
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
func NewAccessWorker(appID, secret string) *Access {
	a := &Access{AppID: appID, AppSecret: secret, httpClient: http.DefaultClient}
	a.getToken()

	return a
}

func (a *Access) getToken() {
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
		fmt.Println(err)
	}

	fmt.Println(result)
}
