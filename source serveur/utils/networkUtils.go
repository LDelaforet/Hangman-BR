package hangmanServer

import (
	"bytes"
	"log"
	"net"
	"strings"
)

// Crée un serveur sur le port donné
func CreateServer(port string) net.Listener {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Erreur lors de la création du serveur : %v", err)
	}
	return listener
}

// Accepte toutes les connexions entrantes
func AcceptConnection(listener net.Listener) {
	for {
		if !*IsListeningPtr {
			break
		}
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalf("Erreur lors de l'acceptation de la connexion : %v", err)
		}
		// Crée un joueur avec la connexion
		go NewPlayer(connection)
	}
}

// Lis en permanance le buffer du joueur et réponds automatiquement
func AutoResponse(conn *Player) {
	// connection := conn.Connection
	buffer := []byte{}
	lastBuffer := []byte{}
	for {
		// Véfifie si conn est toujours écouté
		if !conn.IsListened {
			PrintDebug("Le joueur a l'ID ", conn.Id, " n'est plus écouté, j'me tire")
			break
		}
		// lis le buffer de réponse de conn
		buffer = conn.ReceivedBuffer

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

		switch request {
		case "PING":
			PrintDebug("J'ai reçu un PING")
			WriteSocket(conn.Connection, "PONG")
		case "GETNAME":
			PrintDebug("Nom demandé")
			WriteSocket(conn.Connection, "SETNAME "+ServerName)
		case "SETNAME":
			PrintDebug("Nouveau nom saisi")
			conn.Name = GetArgsFromRequest(bufferString, "SETNAME")[0]
		}
	}
}

// Ecoute ce que dis le client et fous la réponse dans le buffer, je veut que le buffer soit modifié pour toutes les fonctions donc faut utiliser un pointeur, on prends en argument le joueur pour pouvoir fermer la connexion si il y a une erreur
func ListenToClient(conn *Player) {
	connection := conn.Connection
	for {
		// Véfifie si conn est toujours écouté
		if !conn.IsListened {
			PrintDebug("Le joueur a l'ID ", conn.Id, " n'est plus écouté, j'me tire")
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
				log.Println("Connexion fermée par le client")
			} else {
				if strings.Contains(err.Error(), "Une connexion existante a dû être fermée par l’hôte distant") || strings.Contains(err.Error(), "Connexion fermée par le client") {
					conn.IsListened = false
					conn.IsDisconnected = true
					CloseSocket(connection)
					CheckIfDisconnectedPlayers()
					break
				} else {
					log.Printf("Erreur lors de la lecture des données : %v", err)
				}
			}
			// Si aucune erreur, on traite les données
		} else {
			PrintDebug("Données reçues : ", string(buffer))
			// Remplace l'attribut ReceivedBuffer de conn par le buffer
			conn.ReceivedBuffer = buffer
		}
	}
}

// Envoie un message sur la connexion donnée
func WriteSocket(connection net.Conn, message string) {
	connection.Write([]byte(message + "$"))
}

// Ferme la connexion donnée
func CloseSocket(connection net.Conn) {
	connection.Close()
}

// Envoi un message a un joueur précis
func SendMessageToPlayer(player Player, message string) {
	player.Connection.Write([]byte(message + "$"))
}

func SendMessageToAllPlayers(message string) {
	for _, player := range *PlayerListPtr {
		player.Connection.Write([]byte(message + "$"))
	}
}

/*
Liste des requêtes pouvant être envoyées au client :

PING: Demande une réponse
PONG: Réponse a PING
SETNAME name: Envoi son nom au Serveur
GETNAME: Demande le nom du joueur
AREYOUREADY: Demande si le joueur est prêt
NEWVOTE choices: Demande le vote du joueur
GETLETTER: Demande une lettre
*/

/*
Liste des requêtes pouvant être envoyées au serveur par le client :

PING: Demande une réponse
PONG: Réponse a PING
SETNAME name: Envoi son nom au client
LISTPLAYERS: Récupère la liste des joueurs
VOTE vote: Donne son vote
LETTER letter: Envoi une lettre
*/
