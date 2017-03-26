package worker

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

// AUTHKEY the redis key for the cache
const AUTHKEY = "auth-key"

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

func makeHTTPRequest() string {
	fmt.Println("making request")

	return "test"
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
			token := makeHTTPRequest()
			saveTokenToRedis(token)
		}
	}
}

// GetToken ...
func GetToken() string {
	auth, err := client.Get(AUTHKEY).Result()

	if err != nil || auth == "" {
		return makeHTTPRequest()
	}

	return auth
}

func saveTokenToRedis(token string) {
	fmt.Printf("saving token %s \n", token)

	if err := client.Set(AUTHKEY, "test", 0).Err(); err != nil {
		fmt.Printf("error saving %s", err)
	}
}
