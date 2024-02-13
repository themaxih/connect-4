package main

import (
	"log"
	"net"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font/opentype"
)

// Mise en place des polices d'écritures utilisées pour l'affichage.
func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	smallFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 30,
		DPI:  72,
	})
	if err != nil {
		log.Fatal(err)
	}

	largeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 50,
		DPI:  72,
	})
	if err != nil {
		log.Fatal(err)
	}
}

// Création d'une image annexe pour l'affichage des résultats.
func init() {
	offScreenImage = ebiten.NewImage(globalWidth, globalHeight)
}

// Création, paramétrage et lancement du jeu.
func main() {
	// Vérifier si un argument d'adresse a été fourni
	if len(os.Args) != 2 {
		log.Fatal("Usage: ./[filename] [server_address]")
	}

	address := os.Args[1]

	// Connexion au serveur
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Erreur de connexion : ", err)
		return
	}
	defer conn.Close()

	g := game{
		p2Color:       -1, // -1 car elle n'est pas encore choisis
		serverMsgChan: make(chan string, 1),
	}

	// Écoute le serveur pour savoir si tout les joueurs sont connecté
	go g.startListening(conn)

	ebiten.SetWindowTitle("Puissance 4")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
