package main

import (
	"fmt"
	"github.com/Skaifai/RedisProject/util"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// SetString function prompts the user to enter a key, a value, and an expiry duration. It then sets the specified
// key-value pair with the given expiry time.
func SetString() {
	fmt.Print("Enter the key: ")
	key := util.ReadAndCleanString()

	fmt.Print("Enter the value: ")
	val := util.ReadAndCleanString()

	fmt.Print("Enter the expiry duration: ")
	input := util.ReadAndCleanString()

	exp, err := time.ParseDuration(input)
	if err != nil {
		log.Fatal(err)
	}

	currentConnection.Set(ctx, key, val, exp)
}

// GetString function prompts the user to enter a key and retrieves the corresponding value from the Redis store.
func GetString() {
	fmt.Print("Enter the key: ")
	key := util.ReadAndCleanString()
	val, err := currentConnection.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Key: " + key)
	fmt.Println("Value: " + val)
}

// UpdateString function prompts the user to enter a key and a new value, and attempts to update the key-value pair
// with the new value. If the key does not exist, it prints an error message.
func UpdateString() {
	fmt.Print("Enter the key: ")
	key := util.ReadAndCleanString()

	fmt.Print("Enter the value: ")
	val := util.ReadAndCleanString()

	oldValue, _ := currentConnection.Get(ctx, key).Result()

	success := currentConnection.SetXX(ctx, key, val, redis.KeepTTL).Val()
	if !success {
		fmt.Println("Key-value with the specified key does not exist.")
		return
	}
	fmt.Println("Key: " + key)
	fmt.Println("Value: " + oldValue + " --> " + val)
}

// DeleteString function prompts the user to enter a key and deletes the corresponding key-value pair
// from the Redis store. It also prints the key and value of the deleted pair.
func DeleteString() {
	fmt.Print("Enter the key: ")
	key := util.ReadAndCleanString()

	val, err := currentConnection.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	currentConnection.Del(ctx, key)
	fmt.Println("Deleted key-value pair.")
	fmt.Println("Key: " + key)
	fmt.Println("Value: " + val)
}
