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
		// 1. Get current time in Microseconds (high precision)
		now := time.Now().UnixMicro()
		windowStartTime := now -duration.Microseconds()
		// 2. Remove all hits that are now outside the "sliding" window
		rdb.ZRemRangeByScore(ctx, limitKey, "0", fmt.Sprintf("%d",windowStartTime))

		// 3. Count how many hits are left in the last X seconds
		count, _ := rdb.ZCard(ctx, limitKey).Result()

		if count >= limit{
			http.Error(w, "Sliding window limit reached!", 429)
			return
		}

		rdb.ZAdd(ctx, limitKey, redis.Z{
			Score: float64(now),
			Member: now,
		})
		rdb.Expire(ctx, limitKey, duration)
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
