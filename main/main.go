package main

import (
	"context"
	"fmt"
	"github.com/Skaifai/RedisProject/util"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

// const version = "1.0"

var ctx = context.Background()

// Connection variable which will store the current connection to the database
var currentConnection redis.Client

// Global status flag
var isConnectedToDB = false

var connectionTimer *time.Timer

func main() {
	// Input variable declared in advance, because we use
	// input value as the stop condition for the loop
	var input string

	for input != "exit" {
		fmt.Print("Enter command: ")
		input = util.ReadAndCleanString()
		if !isConnectedToDB {
			switch input {
			case "connect":
				fmt.Println("Connecting...")
				currentConnection = Connect()
			case "exit":
				break
			default:
				fmt.Println("Incorrect input.")
			}
		} else {
			switch input {
			case "disconnect":
				fmt.Println("Disconnecting...")
				err := currentConnection.Close()
				if err != nil {
					log.Fatal(err)
				}
				isConnectedToDB = false
			case "set":
				fmt.Println("Setting a key-value pair...")
				SetString()
			case "get":
				fmt.Println("Getting a key-value...")
				GetString()
			case "update":
				fmt.Println("Updating an existing key-value...")
				UpdateString()
			case "delete":
				fmt.Println("Deleting an existing key-value...")
				DeleteString()
			case "subscribe":
				fmt.Println("Entering listening mode...")
				SubscribeAndListen()
			case "publish":
				fmt.Println("Publishing a message...")
				PublishMessage()
			case "exit":
				break
			default:
				fmt.Println("Incorrect input.")
			}
		}
	}
}
