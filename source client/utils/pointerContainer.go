package hangmanClient

import (
	"net"
)

// Connexion actuelle
var CurrentConnection net.Conn
var CurrentConnectionPtr *net.Conn = &CurrentConnection

// Données recues
var ReceivedData []byte
var ReceivedDataPtr *[]byte = &ReceivedData
