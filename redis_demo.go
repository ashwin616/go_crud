package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// Global context (used for Redis operations)
var ctx = context.Background()

func main() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Default DB
	})

	// Test Redis connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("✅ Connected to Redis:", pong)

	// Set a key with a 10-minute expiration
	err = rdb.Set(ctx, "username", "ashwin123", 10*time.Minute).Err()
	if err != nil {
		log.Fatalf("❌ Failed to set key: %v", err)
	}
	fmt.Println("✅ Key 'username' set with value 'ashwin123'")

	// Get the key back from Redis
	val, err := rdb.Get(ctx, "username").Result()
	if err != nil {
		log.Fatalf("❌ Failed to get key: %v", err)
	}
	fmt.Println("🔁 Retrieved from Redis - 'username':", val)
}
