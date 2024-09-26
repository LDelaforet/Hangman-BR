package Hangman

import (
	"net"
)

type Player struct {
	id         int
	name       string
	score      int
	word       string
	lives      int
	connection net.Conn
}

var ClientList []net.Conn

// Données recues
var ReceivedData []byte
var ReceivedDataPtr *[]byte = &ReceivedData
