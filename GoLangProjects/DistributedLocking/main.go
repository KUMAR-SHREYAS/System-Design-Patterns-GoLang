package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

var ctx = context.Background()

// AcquireLock tries to grab a unique lock in Redis
func AcquireLock(rdb *redis.Client, lockKey string, expiration time.Duration) (string, bool) {
	// 1. Create a unique token so we know we own this lock
	token := uuid.New().String()

	// 2. SET if Not Exists (NX) with an Expiration (EX)
	// This is atomic - only one server can succeed
	success, err := rdb.SetNX(lockKey, token, expiration).Result()
	if err != nil {
		log.Printf("Redis Error: %v", err)
		return "", false
	}

	return token, success
}

// Lua script: Returns 1 if deleted, 0 if token mismatch or key doesn't exist
var unlockLua = redis.NewScript(`
	if redis.call("get", KEYS[1])==ARGV[1] then
		return redis.call("del", KEYS[1])
	else 
		return 0
	end
`)

// ReleaseLock uses Lua to ensure we only delete our OWN lock
func ReleaseLock(rdb *redis.Client, lockKey string, token string) bool {
	result, err := unlockLua.Run(rdb, []string{lockKey}, token).Int()
	if err != nil {
		log.Printf("unlock Error: %v", err)
		return false
	}
	return result == 1
}
func main() {
	// Initialize Redis Client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	lockKey := "lock:payment_processor"
	// We'll set a short timeout to test the safety logic
	lockTimeout := 4 * time.Second

	fmt.Println("Attempting to acquire lock...")

	// Try to grab the lock
	token, success := AcquireLock(rdb, lockKey, lockTimeout)

	if !success {
		fmt.Println("❌ Busy Try again later.")
		return
	}

	// SUCCESS: We own the lock now
	fmt.Printf("✅ Lock Acquired! Token: %s\n", token)
	fmt.Println("🚧 Processing... (This will outlast the lock)")

	// Simulate a long task (5 seconds), wil outlast the lock
	time.Sleep(5 * time.Second)

	// Attempt to release
	release := ReleaseLock(rdb, lockKey, token)
	if release {
		fmt.Println("🔓 Lock released successfully.")
	} else {
		fmt.Println("⚠️ Failed to release: Lock already expired or owned by someone else!")
	}
}
