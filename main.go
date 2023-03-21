package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Skaifai/RedisProject/util"
	"github.com/redis/go-redis/v9"
	"os"
)

const version = "1.0"

var ctx = context.Background()

// Connection variable which will store the current connection to the database
var currentConnection redis.Client

// Global status flag
var isConnectedToDB bool = false

// Instance of a reader class, which will handle the input from the console
var reader = bufio.NewReader(os.Stdin)

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
				currentConnection.Close()
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
			case "exit":
				break
			default:
				fmt.Println("Incorrect input.")
			}
		}
	}
}
