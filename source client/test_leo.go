package main

import (
	"fmt"
	client "hangmanClient/utils"
)

func main() {
	conn := client.ConnectSocket("127.0.0.1", "1597")
	go client.ListenSocket(conn)
	for {
		// Get user input
		input := client.GetInput()
		client.WriteSocket(conn, input)
		fmt.Println("Message: ")
		fmt.Println(client.ReceivedData)
	}
}
