package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	// Lancement du serveur
	log.Println("Launching server...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()
	log.Println("Server started")
	log.Println("Waiting for connections...")

	// Attente de la connexion du joueur 1
	conn1, err := listener.Accept()
	if err != nil {
		log.Print(err)
		return
	}
	defer conn1.Close()
	log.Println("Joueur 1 connecté !")
	reader1 := bufio.NewReader(conn1)
	writer1 := bufio.NewWriter(conn1)
	writer1.WriteString("Je suis le joueur 1\n")
	writer1.Flush()

	// Attente de la connexion du joueur 2
	conn2, err := listener.Accept()
	if err != nil {
		log.Print(err)
		return
	}
	defer conn2.Close()
	log.Println("Joueur 2 connecté !")
	reader2 := bufio.NewReader(conn2)
	writer2 := bufio.NewWriter(conn2)
	writer2.WriteString("Je suis le joueur 2\n")
	writer2.Flush()

	chaine := "Tous les joueurs sont connectés !\n"
	writer1.WriteString(chaine)
	writer2.WriteString(chaine)
	writer1.Flush()
	writer2.Flush()
	log.Print(chaine)

	player1Chan := make(chan string)
	player2Chan := make(chan string)

	go startListening(1, reader1, &player1Chan)
	go startListening(2, reader2, &player2Chan)

	for {
		select {
		case msg := <-player1Chan:
			msg = strings.TrimSuffix(msg, "\n")
			s := strings.Split(msg, " : ")
			msgType, data := s[0], s[1]
			switch msgType {
			case "Player 1 chose color":
				msg = "Opponent chose color : " + data + "\n"

			case "Player 1 played at":
				msg = "Opponent played at : " + data + "\n"
			}
			writer2.WriteString(msg)
			writer2.Flush()

		case msg := <-player2Chan:
			msg = strings.TrimSuffix(msg, "\n")
			s := strings.Split(msg, " : ")
			msgType, data := s[0], s[1]
			switch msgType {
			case "Player 2 chose color":
				msg = "Opponent chose color : " + data + "\n"
			case "Player 2 played at":
				msg = "Opponent played at : " + data + "\n"
			}
			writer1.WriteString(msg)
			writer1.Flush()
		}
	}
}

func startListening(nJoueur int, reader *bufio.Reader, channel *chan string) {
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				log.Fatal(fmt.Sprintf("Le joueur %d s'est déconnecté ! Arrêt du serveur...", nJoueur))
			}
			log.Fatal(err)
		}
		*channel <- msg
	}
}
