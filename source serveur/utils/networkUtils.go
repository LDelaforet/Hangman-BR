package Hangman

import (
	"log"
	"net"
)

func ListenSocket(connection net.Conn) {
	for {
		buffer := make([]byte, 1024)
		_, err := connection.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("Connexion fermée par le client")
			} else {
				log.Printf("Erreur lors de la lecture des données : %v", err)
			}
		} else {
			ReceivedData = buffer
		}
	}
}

func WriteSocket(connection net.Conn, message string) {
	connection.Write([]byte(message + "$"))
}

func ConnectSocket(ip string, port string) net.Conn {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func CloseSocket(connection net.Conn) {
	connection.Close()
}
