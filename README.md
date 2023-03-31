# REDIS CLI for CRUD Operations and PUBSUB
## A go-redis project for the internship at "QZlab".

This is a simple Redis CLI (Command-Line Interface) program written in Golang that allows you to perform basic CRUD (Create, Read, Update, Delete) operations on a Redis database, as well as publish and subscribe to channels using the Redis Pub/Sub messaging system. This program uses the redis-go package for interacting with Redis.

### Prerequisites
* Golang 1.16 or higher installed on your system
* Redis server running on your machine or on a remote server with the necessary authentication information if needed.

### Installation
1. Clone this repository or download the source code as a ZIP file.
2. Install the required dependencies by running go get github.com/go-redis/redis/v9 in your terminal or command prompt.

### Usage
1. Start your Redis server.
2. Open a terminal or command prompt and navigate to the directory where you saved the program.
3. Input the **"connect"** keyword into the CLI to connect to the server.
4. Input the target server's address, database number and the supposed connection lifetime. For example:
    ```
    Enter command: connect
    Connecting...
    Enter the host address: localhost:6379
    Enter the db: 0
    Enter the expiry duration: 30m
    PONG
    Connection successful.
    ```
5. Input commands as needed.
6. To disconnect input the `disconnect` command. To exit the application at any time input the `exit` command.

#### Basic application flow
![Redis main white bg](https://user-images.githubusercontent.com/85491176/229050698-04c0c800-3cef-40be-ac3c-1d2bb9ccfe1b.png)

### CRUD Operations
The following CRUD operations are supported by this program (followed by their equivalents in the redis-cli):

* `set`: SET key value
* `get`: GET key
* `update`: SET key value
* `delete`: DEL key

To perform any of these operations, simply type the corresponding command and follow the instructions. 

For example for the `set` command:
```
Enter command: set    
Setting a key-value pair...
Enter the key: key1
Enter the value: value1
Enter the expiry duration: 1h
```

For the `get` command:
```
Enter command: get
Getting a key-value...
Enter the key: key1 
Key: key1
Value: value1
```

For the `update` command:
```
Enter command: update    
Updating an existing key-value...
Enter the key: key1
Enter the value: value2
Key: key1
Value: value1 --> value2
```

For the `delete` command:
```
Enter command: delete
Deleting an existing key-value...
Enter the key: key1
Deleted key-value pair.
Key: key1
Value: value2
```

### Pub/Sub
#### General
This program also supports the Redis Pub/Sub messaging system. Follow these steps in redis-cli to see how the feature works:
1. In one terminal window, run the subscribe command followed by the channel name to subscribe to:
    ```
    subscribe mychannel
    ```
2. In another terminal window, run the `publish` command followed by the channel name and the message to send:
    ```
    publish mychannel "Hello, world!"
    ```
3. In the first terminal window, you should see the message "Hello, world!" displayed in the console.

#### Usage
1. In this program the client can enter the 'listening mode', which will restrict the user from executing any commands and    only receive messages from the specified channels. To do that you must input the `subscribe` command. Example:
    ```
    Enter command: subscribe
    Entering listening mode.
    Enter the channels: channel1 channel2
    Use the "quit" command to quit to exit the listening mode. No other commands are allowed.
    Subscribed current connection to channels: PubSub(channel1 channel2)
    ```
2. To unsubscribe the client from the channels and exit the listening mode you input the `quit` command like so:
    ```
    quit
    Exiting the listening mode.
    ```
3. To publish a message to specified channels input the `publish` command. Example:
    ```
    Enter command: publish
    Publishing a message...
    A list of currently active channels:
    [channel2 channel1]
    Enter the channels: channel1 channel2
    Enter the message: hello!
    Channel: channel1
    Receivers: 1
    Channel: channel2
    Receivers: 1
    ```
### Credits
This program uses the redis-go package, which is a Redis client for Golang developed by the Go Redis team.
