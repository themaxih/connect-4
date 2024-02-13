package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

var (
	writer *bufio.Writer
	reader *bufio.Reader
)

// Envoie de la couleur séléctionnée au serveur
func (g *game) sendColor(color int) {
	writer.WriteString(fmt.Sprintf("Player %d chose color : %d\n", g.playerNumber, color))
	if err := writer.Flush(); err != nil {
		log.Printf("Erreur lors de l'envoi de la couleur : %v", err)
	}
}

// Envoie de la position du selecteur
func (g *game) sendColorPosition() {
	writer.WriteString()
}

// Envoie de la position jouée au serveur
func (g *game) sendPosition() {
	writer.WriteString(fmt.Sprintf("Player %d played at : %d\n", g.playerNumber, g.tokenPosition))
	if err := writer.Flush(); err != nil {
		log.Printf("Erreur lors de l'envoi du mouvement : %v", err)
	}
}

// Incrémenter le timer toutes les secondes
func (g *game) startTimer() {
	for {
		select {
		case <-time.Tick(time.Second): 
			if g.gameState == playState {
				g.timer++
			} else {
				g.timer = 0
			}
		}
	}
}

// Ecoute les messages du serveur
func (g *game) startListening(conn net.Conn) {
	writer = bufio.NewWriter(conn)
	reader = bufio.NewReader(conn)

	go g.startTimer()

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Print(err)
			return
		}

		log.Print(msg)
		switch msg {
		case "Je suis le joueur 1\n":
			g.turn = p1Turn
			g.playerNumber = 1
		case "Je suis le joueur 2\n":
			g.turn = p2Turn
			g.playerNumber = 2
		default:
			g.serverMsgChan <- msg
		}
	}
}
