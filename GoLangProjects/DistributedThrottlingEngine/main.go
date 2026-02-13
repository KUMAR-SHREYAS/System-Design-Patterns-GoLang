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

func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			return
		}
		userIp, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			userIp = r.RemoteAddr
		}
		count, _ := rdb.Incr(ctx, userIp).Result()

		if count == 1 {
			rdb.Expire(ctx, userIp, 10*time.Second)
		}
		if count > 5 {
			http.Error(w, "Too many requests. Slow down!", http.StatusTooManyRequests)
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

	http.HandleFunc("/", rateLimitMiddleware(homePage))
	http.HandleFunc("/login", rateLimitMiddleware(loginPage))

	fmt.Println("Server starting at :8080...")
	http.ListenAndServe(":8080", nil) // spin up the server.

}
