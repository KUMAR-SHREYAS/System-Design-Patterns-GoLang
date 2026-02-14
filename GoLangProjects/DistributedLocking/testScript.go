// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"

// 	"github.com/go-redis/redis"
// )

// func RunChaosTest() {
// 	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
// 	lockKey := "chaos_test_lock"
// 	ctx := context.Background()

// 	// Cleanup any old locks
// 	rdb.Del(lockKey)

// 	fmt.Println("🧪 Starting Chaos Test...")

// 	// TEST 1: The "Slow Worker" & The "Expiration"
// 	fmt.Println("\n--- Test 1: Expiration Safety ---")
// 	token1, _ := AcquireLock(rdb, lockKey, 2*time.Second)
// 	fmt.Println("Worker 1 acquired lock for 2s...")

// 	time.Sleep(3 * time.Second) // Wait for lock to expire

// 	// Worker 2 should be able to grab it now
// 	token2, success2 := AcquireLock(rdb, lockKey, 5*time.Second)
// 	if success2 {
// 		fmt.Println("Worker 2 successfully grabbed the expired lock.")
// 	}

// 	// TEST 2: The "Illegal Release"
// 	fmt.Println("\n--- Test 2: Token Mismatch Protection ---")
// 	fmt.Println("Worker 1 tries to release the lock it no longer owns...")
// 	released := ReleaseLock(rdb, lockKey, token1) // Using Worker 1's old token
// 	if !released {
// 		fmt.Println("✅ Success: Lua script blocked Worker 1 from stealing Worker 2's lock!")
// 	}

// 	// TEST 3: The "Correct Release"
// 	fmt.Println("\n--- Test 3: Valid Release ---")
// 	released2 := ReleaseLock(rdb, lockKey, token2)
// 	if released2 {
// 		fmt.Println("✅ Success: Worker 2 released its own lock correctly.")
// 	}

// 	rdb.Del(ctx, lockKey)
// }

