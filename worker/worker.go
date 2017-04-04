package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

// AUTHKEY the redis key for the cache
const AUTHKEY = "auth-key"
const accessTokenAPI = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"

// Worker ...
type Worker struct {
	AppID     string
	Secret    string
	Frequency int
}

var client *redis.Client

// Init initializes and returns the worker
func Init(uri, password string, db int, appID, secret string, frequency int) *Worker {
	w := &Worker{
		AppID:     appID,
		Secret:    secret,
		Frequency: frequency,
	}

	client = redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: password,
		DB:       db,
	})

	return w
}

// Start ...
func (w *Worker) Start(ch chan bool) {
	t := time.NewTicker(time.Duration(w.Frequency) * time.Second)
	defer close(ch)

	for {
		select {
		case <-ch:
			log.Println("Stopping ticker")
			t.Stop()
			ch <- true
		case <-t.C:
			token, err := makeHTTPRequest(w)
			if err != nil {
				log.Printf("Error in makeHTTPRequest, %s", err)
			} else {
				saveTokenToRedis(token)
			}
		}
	}
}

// GetToken ...
func (w *Worker) GetToken() (string, error) {
	auth, err := client.Get(AUTHKEY).Result()

	fmt.Println("HELLO")
	if err != nil || auth == "" {
		return makeHTTPRequest(w)
	}

	return auth, nil
}

func saveTokenToRedis(token string) {
	fmt.Printf("saving token %s \n", token)

	if err := client.Set(AUTHKEY, token, 0).Err(); err != nil {
		fmt.Printf("error saving %s", err)
	}
}

type accessError struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

type accessResponse struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

func makeHTTPRequest(w *Worker) (string, error) {
	fmt.Println("making request")

	url := fmt.Sprintf(accessTokenAPI, w.AppID, w.Secret)
	log.Println(url)

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", err
	}

	var result struct {
		accessError
		accessResponse
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	fmt.Println("YYEA", result.Token)

	return result.Token, nil
}
