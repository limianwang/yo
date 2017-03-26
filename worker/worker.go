package worker

import "github.com/go-redis/redis"

// AUTHKEY the redis key for the cache
const AUTHKEY = "auth-key"

var c *redis.Client

// Start initializes and starts the worker
func Start(uri, password string) {
	c = redis.NewClient(&redis.Options{

		Addr:     uri,
		Password: password,
		DB:       0,
	})

	c.Ping()
}

func makeHttpRequest() string {
	return "test"
}

func GetToken() string {
	auth, err := c.Get(AUTHKEY).Result()

	if err != nil {
		return makeHttpRequest()
	}

	return auth
}
