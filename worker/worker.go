package worker

import (
	"fmt"

	"github.com/go-redis/redis"
)

// AUTHKEY the redis key for the cache
const AUTHKEY = "auth-key"

var c *redis.Client

// Start initializes and starts the worker
func Start(uri, password string, db int) {
	c = redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: password,
		DB:       db,
	})

	v, _ := c.Ping().Result()
	fmt.Println("result:", v)
}

func makeHTTPRequest() string {
	fmt.Println("making request")
	return "test"
}

func GetToken() string {
	auth, err := c.Get(AUTHKEY).Result()

	if err != nil {
		return makeHTTPRequest()
	}

	return auth
}
