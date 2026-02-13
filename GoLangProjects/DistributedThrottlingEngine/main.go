package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func rateLimitMiddleware(next http.HandlerFunc, limit int64, duration time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			return
		}
		userIp, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			userIp = r.RemoteAddr
		}
		limitKey := fmt.Sprintf("%s:%s", userIp, r.URL.Path)
		count, _ := rdb.Incr(ctx, limitKey).Result()

		if count == 1 {
			rdb.Expire(ctx, limitKey, duration)
		}
		if count > limit {
			http.Error(w, fmt.Sprintf("Limit reached! Max %d hits per %v", limit, duration), 429)
			return
		}
		next(w, r)
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HOME page!.")
}
func loginPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the LOGIN page!.")
}
func main() {
	// 1. Establishing connection to local redis server using rdb client.
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	http.HandleFunc("/", rateLimitMiddleware(homePage, 10, 30*time.Second))
	http.HandleFunc("/login", rateLimitMiddleware(loginPage, 3, 1*time.Minute))

	fmt.Println("Server starting at :8080...")
	http.ListenAndServe(":8080", nil) // spin up the server.

}
