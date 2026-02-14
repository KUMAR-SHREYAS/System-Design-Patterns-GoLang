package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/go-redis/redis"
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

func main() {
	// Initialize Redis Client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	lockKey := "lock:payment_processor"
	lockTimeout := 10 * time.Second

	fmt.Println("Attempting to acquire lock...")

	// Try to grab the lock
	token, success := AcquireLock(rdb, lockKey, lockTimeout)

	if !success {
		fmt.Println("❌ Could not get lock. Another instance is already working.")
		return
	}

	// SUCCESS: We own the lock now
	fmt.Printf("✅ Lock Acquired! Token: %s\n", token)
	fmt.Println("🚧 Performing critical business logic (e.g., processing payment)...")

	// Simulate a long task (5 seconds)
	time.Sleep(5 * time.Second)

	fmt.Println("🎉 Task finished!")

	// For Step 1, we manually delete it, but Step 2 will make this safer!
	rdb.Del(lockKey)
	fmt.Println("🔓 Lock released.")
}