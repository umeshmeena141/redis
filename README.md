# Redis Clone

A clone of Redis written in Go. This application provides in-memory caching, data manipulation, and publish/subscribe functionality.

## Features

- In-memory data storage with support for multiple data types (e.g., strings, integers, lists, sets, objects).
- Publish/subscribe mechanism for channels.
- Support for multiple tables to organize data.
- Command-based interaction over TCP.

## How to Run the Application

1. **Compile the Application**:
   Use the provided `compile.sh` script to build the application:

   ```sh
   ./compile.sh
   ```

   This will generate the executable at `build-redis-go`.

2. **Run the Application**:
   Start the server by running the compiled binary:

   ```sh
   build-redis-go
   ```

   The server will start listening for TCP connections on port `6379`.

3. **Connect to the Server**:
   Use a TCP client (e.g., `telnet` or `nc`) to connect to the server:
   ```sh
   telnet 127.0.0.1 6379
   ```

## Available Commands

Below is a list of supported commands, their arguments, and descriptions:

### General Commands

- **PING**  
  Usage: `PING`  
  Response: `PONG`.

### Data Manipulation Commands

- **GET**  
  Usage: `GET <key>`  
  Description: Retrieves the value associated with the given key.

- **SET**  
  Usage: `SET <key> <value> <type>`  
  Description: Sets the value for the given key with the specified type (`int`, `float`, `string`, etc.).

- **DEL**  
  Usage: `DEL <key>`  
  Description: Deletes the value associated with the given key.

### List Commands

- **LPUSH**  
  Usage: `LPUSH <key> <json_value>`  
  Description: Pushes a value to the end of the list.

- **LPOP**  
  Usage: `LPOP <key>`  
  Description: Pops a value from the end of the list.

- **LGET**  
  Usage: `LGET <key> <index>`  
  Description: Retrieves the value at the specified index in the list.

### Set Commands

- **SADD**  
  Usage: `SADD <key> <value>`  
  Description: Adds a value to the set.

- **SPOP**  
  Usage: `SPOP <key> <value>`  
  Description: Removes a value from the set.

- **SFIND**  
  Usage: `SFIND <key> <value>`  
  Description: Checks if a value exists in the set. Returns `TRUE` or `FALSE`.

### Numeric Commands

- **INC**  
  Usage: `INC <key>`  
  Description: Increments the integer value by 1.

- **INC_N**  
  Usage: `INC_N <key> <n>`  
  Description: Increments the integer value by `n`.

- **DEC**  
  Usage: `DEC <key>`  
  Description: Decrements the integer value by 1.

- **DEC_N**  
  Usage: `DEC_N <key> <n>`  
  Description: Decrements the integer value by `n`.

### Table Commands

- **SET_TABLE**  
  Usage: `SET_TABLE <table_name>`  
  Description: Switches the current table to the specified table.

- **CREATE_TABLE**  
  Usage: `CREATE_TABLE <table_name>`  
  Description: Creates a new table with the specified name.

### Publish/Subscribe Commands

- **CREATE_CHAN**  
  Usage: `CREATE_CHAN <channel_name>`  
  Description: Creates a new channel.

- **SUB_CHAN**  
  Usage: `SUB_CHAN <channel_name>`  
  Description: Subscribes the current connection to the specified channel.

- **PUB_CHAN**  
  Usage: `PUB_CHAN <channel_name> <message>`  
  Description: Publishes a message to the specified channel.

## Example Usage

1. **Set and Get a Value**:

   ```sh
   SET mykey 42 int
   GET mykey
   ```

2. **Work with Lists**:

   ```sh
   LPUSH mylist {"val":"item1","type":"string"}
   LGET mylist 0
   ```

3. **Publish/Subscribe**:
   ```sh
   CREATE_CHAN mychannel
   SUB_CHAN mychannel
   PUB_CHAN mychannel "Hello, World!"
   ```

## Notes

- Ensure that the server is running before issuing commands.
- Commands are case-sensitive.
- Use valid JSON for list and object values where applicable.
