package main

import (
	"fmt"
	"github.com/Skaifai/RedisProject/util"
	"github.com/redis/go-redis/v9"
	"log"
	"strings"
)

var pubsub *redis.PubSub

// SubscribeAndListen asks the user to input the channels to subscribe to, subscribes to these channels, and listens for messages on these
// channels until the user enters the "quit" command. While listening, it prints out any received messages.
func SubscribeAndListen() {
	fmt.Print("Enter the channels: ")
	channels := util.ReadAndCleanString()
	pubsub = currentConnection.Subscribe(ctx, channels)
	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(pubsub)

	result, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Use the \"quit\" command to exit the listening mode. No other commands are allowed.")

	// Should be *Subscription, but others are possible if other actions have been
	// taken on pubsub since it was created.
	switch result.(type) {
	case *redis.Subscription:
		// subscribe succeeded
		fmt.Println("Subscribed current connection to channels: " + pubsub.String())
	case *redis.Message:
		// Assert type
		m, ok := result.(redis.Message)
		if ok {
			// received first message
			fmt.Println("Received first message.")
			fmt.Println("Channel: " + m.Channel)
			fmt.Println("Message: " + m.Payload)
		}
	case *redis.Pong:
		// pong received
		fmt.Println("Pong received.")
	default:
		log.Fatal("Error in redis has occurred.")
	}

	// Make a channel that will track if the exit command was issued
	quit := make(chan string)
	go func() {
		for {
			input := util.ReadAndCleanString()
			if input == "quit" {
				quit <- "quit"
				close(quit)
				return
			}
		}
	}()

	// Make a channel that will receive messages
	ch := pubsub.Channel()
	for {
		select {
		case message := <-ch:
			fmt.Println("Channel: " + message.Channel + " \nMessage: " + message.Payload)
		case <-quit:
			fmt.Println("Exiting the listening mode.")
			return
		}
	}
}

// PublishMessage function lists the currently active channels, asks the user to enter the channels and message
// to publish, and publishes the message to the specified channels. For each channel, it also prints out the number of
// subscribers that received the message.
func PublishMessage() {
	fmt.Println("A list of currently active channels:")
	listOfChannels, err := currentConnection.PubSubChannels(ctx, "*").Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(listOfChannels)

	fmt.Print("Enter the channels: ")
	channelsString := util.ReadAndCleanString()
	fmt.Print("Enter the message: ")
	message := util.ReadAndCleanString()

	channels := strings.Split(channelsString, " ")
	for _, channel := range channels {
		currentConnection.Publish(ctx, channel, message)
		channelSubscribers, err := currentConnection.PubSubNumSub(ctx, channel).Result()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Channel: " + channel)
		fmt.Println("Receivers:", channelSubscribers[channel])
	}
}
