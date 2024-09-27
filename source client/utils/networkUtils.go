package hangmanClient

import (
	"bytes"
	"log"
	"net"
	"strings"
)

// Envoie un message au serveur
func SendMessageToServer(serv Server, message string) {
	serv.Connection.Write([]byte(message + "$"))
}

func ListenServer(serv *Server) {
	connection := serv.Connection
	for {
		// Véfifie si conn est toujours écouté
		if !serv.IsListened {
			PrintDebug("Le serveur n'est plus écouté, j'me casse")
			break
		}
		PrintDebug("En attente de données de la part du client")
		// Crée un buffer pour les données reçues
		buffer := make([]byte, 1024)

		// Lit les données reçues
		_, err := connection.Read(buffer)

		// Gère les erreurs
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("Connexion fermée par le serveur")
			} else {
				log.Printf("Erreur lors de la lecture des données : %v", err)
			}
			// Si aucune erreur, on traite les données
		} else {
			PrintDebug("Données reçues : ", string(buffer))
			// Remplace l'attribut ReceivedBuffer de conn par le buffer
			serv.ReceivedBuffer = buffer
		}
	}
}

// Lis en permanance le buffer du joueur et réponds automatiquement
func AutoResponse(serv *Server) {
	// connection := conn.Connection
	buffer := []byte{}
	lastBuffer := []byte{}
	for {
		// Véfifie si conn est toujours écouté
		if !serv.IsListened {
			PrintDebug("Le serveur n'est plus écouté, j'me casse")
			break
		}
		// lis le buffer de réponse de conn
		buffer = serv.ReceivedBuffer

		// Le buffer n'a pas changé, on skip
		if AreByteSliceEqual(lastBuffer, buffer) {
			continue
		}

		// Si le buffer est vide, on skip
		if len(buffer) == 0 {
			continue
		}

		// Copie le buffer dans lastBuffer avant de tronquer le buffer
		lastBuffer = make([]byte, len(buffer))
		copy(lastBuffer, buffer)

		buffer = buffer[:bytes.IndexByte(buffer, 0)]

		// Retire les caractères: $, \r et \n
		buffer = RemoveFromByteSlice(buffer, []byte{36, 13, 10})

		// On met le buffer en string et la vraie verficiation commence
		bufferString := string(buffer)
		request := strings.Split(bufferString, " ")[0]
		PrintDebug("Requête reçue : ", request)
		switch request {
		// Gère les demandes du serveur
		case "PING":
			PrintDebug("Requête reçue : PING")
			WriteSocket(serv.Connection, "PONG")
		case "GETNAME":
			PrintDebug("Requête reçue : GETNAME")
			WriteSocket(serv.Connection, "SETNAME "+PlayerName)

		// Gère les réponses aux requêtes
		case "PONG":
			PrintDebug("Réponse reçue : PONG")
			// On arrête la fonction LateSet
			serv.LateSetBypass = true
			// On attends que le LateSet ait bien fini
			WaitMilliseconds(500)
			serv.RequestResponded = true
			serv.LastRequest = "PONG"
			go LateSet(&serv.RequestResponded, false, 5000, &serv.LateSetBypass)
		case "SETNAME":
			PrintDebug("Réponse reçue : SETNAME")
			serv.Name = GetArgsFromRequest(bufferString, "SETNAME")[0]
			// On arrête la fonction LateSet si elle est en cours
			serv.LateSetBypass = true
			// On attends que le LateSet ait bien fini de se fermer
			WaitMilliseconds(500)
			serv.RequestResponded = true
			serv.LastRequest = "SETNAME"
			go LateSet(&serv.RequestResponded, false, 5000, &serv.LateSetBypass)
		default:
			PrintDebug("Requête ", request, " non reconnue")
		}

	}
}

// Ecoute ce qu'il se passe sur le socket
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

// Se connecte au socket donné
func ConnectSocket(ip string, port string) net.Conn {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

// Ferme la connexion donnée
func CloseSocket(connection net.Conn) {
	connection.Close()
}

// Envoie un message sur la connexion donnée
func WriteSocket(connection net.Conn, message string) {
	connection.Write([]byte(message + "$"))
}
