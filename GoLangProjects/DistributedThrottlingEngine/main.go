package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// 1. Establishing connection to local redis server using rdb client.
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" { // Disable the tab icon request.
			return
		}
		userIp := r.RemoteAddr //  Identifying the visitor with IP Address

		count, err := rdb.Incr(ctx, userIp).Result() // redis INCR counts the userIp hits
		if err != nil {
			http.Error(w, "Redis Error", 500)
			return
		}

		if count == 1 { // first count sets an expiration to the counter
			rdb.Expire(ctx, userIp, 10*time.Second)
		}

		if count > 5 {// rate limiter
			fmt.Fprintf(w, "Slow down!, You've visited %d times in 10 seconds. Blocked!", count)
			return
		}
		fmt.Fprintf(w, "Welcome! You have visited %d times. Refresh to test the limit.", count)
	})

	fmt.Println("Server starting at :8080...")
	http.ListenAndServe(":8080", nil) // spin up the server.

}
