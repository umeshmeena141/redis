package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"redis/commands"
	"redis/enums"
	"redis/storage"
	"strings"
)

func splitArgs(input string) []string {
	var args []string
	var current strings.Builder
	i := 0
	for i < len(input) {
		r := input[i]
		switch r {
		case '\'':
			i += 1
			for i < len(input) && input[i] != '\'' {
				if i < len(input)-1 && input[i:i+2] == "\\'" {
					i += 1
				}
				current.WriteByte(input[i])
				i += 1
			}
		case ' ':
			if current.Len() > 0 {
				args = append(args, current.String())
			}
			current.Reset()
		default:
			current.WriteByte(r)
		}
		i += 1
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}

func handleCommand(conn net.Conn, command string, dataStore *storage.Storage, tableName *string) string {
	log.Println("Recieved command:", command)
	args := splitArgs(command)
	if len(args) < 1 {
		log.Printf("No Command Found")
		return enums.GetResponseMessage(enums.UNKNOWN_COMMAND)
	}
	log.Println("Args: ", strings.Join(args, ","), "Length of Args", len(args))
	cmd := commands.GetCMD(conn, dataStore, tableName, args)
	return cmd.HandleCommand()

}

func handleTCPConnection(conn net.Conn, dataStore *storage.Storage) {
	defer func() {
		log.Println("Closing connection")
		conn.Close()
	}()
	log.Println("Established new TCP connection from:", conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)
	tableName := "default"
	for scanner.Scan() {
		command := scanner.Text()
		response := handleCommand(conn, command, dataStore, &tableName)
		log.Println("Sending response back to :", conn.RemoteAddr())
		fmt.Fprintln(conn, response)
	}
}

func listenConnections(listener net.Listener, dataStore *storage.Storage) {

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleTCPConnection(conn, dataStore)
	}
}
