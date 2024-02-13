package main

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Affichage des graphismes à l'écran selon l'état actuel du jeu.
func (g *game) Draw(screen *ebiten.Image) {

	screen.Fill(globalBackgroundColor)

	switch g.gameState {
	case titleState:
		g.titleDraw(screen)
	case colorSelectState:
		g.colorSelectDraw(screen)
	case playState:
		g.playDraw(screen)
	case resultState:
		g.resultDraw(screen)
	}

}

// Affichage des graphismes de l'écran titre.
func (g game) titleDraw(screen *ebiten.Image) {
	text.Draw(screen, "Puissance 4 en réseau", largeFont, 90, 150, globalTextColor)
	text.Draw(screen, "Projet de programmation système", smallFont, 105, 190, globalTextColor)
	text.Draw(screen, "Année 2023-2024", smallFont, 210, 230, globalTextColor)

	if g.stateFrame >= globalBlinkDuration/3 {
		if g.otherConnected {
			text.Draw(screen, "Appuyez sur entrée", smallFont, 210, 500, globalTextColor)
		} else {
			text.Draw(screen, "En attente de l'adversaire...", smallFont, 150, 500, globalTextColor)
		}
	}
}

// Affichage des graphismes de l'écran de sélection des couleurs des joueurs.
func (g game) colorSelectDraw(screen *ebiten.Image) {
	text.Draw(screen, "Quelle couleur pour vos pions ?", smallFont, 110, 130, globalTextColor)

	line := 0
	col := 0
	for numColor := 0; numColor < globalNumColor; numColor++ {
		
		xPos := (globalNumTilesX-globalNumColorCol)/2 + col
		yPos := (globalNumTilesY-globalNumColorLine)/2 + line

		if numColor == g.p1Color {
			vector.DrawFilledCircle(screen, float32(globalTileSize/2+xPos*globalTileSize), float32(globalTileSize+globalTileSize/2+yPos*globalTileSize+80), globalTileSize/2, globalSelectColor, true)
		}

		vector.DrawFilledCircle(screen, float32(globalTileSize/2+xPos*globalTileSize), float32(globalTileSize+globalTileSize/2+yPos*globalTileSize+80), globalTileSize/2-globalCircleMargin, globalTokenColors[numColor], true)

		// Dessin d'une croix sur la couleur de l'adversaire
		if numColor == g.p2Color {
			// Dessiner une croix sur la couleur du joueur 2
			xCenter := float32(globalTileSize/2 + xPos*globalTileSize)
			yCenter := float32(globalTileSize + globalTileSize/2 + yPos*globalTileSize + 80)
			crossSize := float32(33)

			// Dessiner deux lignes pour former la croix
			vector.StrokeLine(screen, xCenter-crossSize, yCenter-crossSize, xCenter+crossSize, yCenter+crossSize, 3.0, globalCrossColor, true)
			vector.StrokeLine(screen, xCenter+crossSize, yCenter-crossSize, xCenter-crossSize, yCenter+crossSize, 3.0, globalCrossColor, true)
		}

		col++
		if col >= globalNumColorCol {
			col = 0
			line++
		}
	}

	if g.colorSelected {
		text.Draw(screen, "En attente de l'adversaire...", smallFont, 160, 650, globalTextColor)
	}
}


// Affichage des graphismes durant le jeu.
func (g game) playDraw(screen *ebiten.Image) {
	g.drawGrid(screen)

	text.Draw(screen, strconv.Itoa(g.timer), largeFont, 10, 50, globalTextColor)
	text.Draw(screen, "Au tour de ", largeFont, 370, 50, globalTextColor)
	if (g.turn == p1Turn) {
		vector.DrawFilledCircle(screen, 662, 37, 27, globalTokenColors[g.p1Color], true)
	} else {
		vector.DrawFilledCircle(screen, 662, 37, 27, globalTokenColors[g.p2Color], true)
	}


	vector.DrawFilledCircle(screen, float32(globalTileSize/2+g.tokenPosition*globalTileSize), float32(globalTileSize/2+80), globalTileSize/2-globalCircleMargin, globalTokenColors[g.p1Color], true)
}

// Affichage des graphismes à l'écran des résultats.
func (g game) resultDraw(screen *ebiten.Image) {
	g.drawGrid(offScreenImage)

	options := &ebiten.DrawImageOptions{}
	options.ColorScale.ScaleAlpha(0.2)
	screen.DrawImage(offScreenImage, options)

	message := "Égalité"
	if g.result == p1wins {
		message = "Gagné !"
	} else if g.result == p2wins {
		message = "Perdu…"
	}
	text.Draw(screen, message, smallFont, 300, 400, globalTextColor)
}

// Affichage de la grille de puissance 4, incluant les pions déjà joués.
func (g game) drawGrid(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, globalTileSize+80, globalTileSize*globalNumTilesX, (globalTileSize+80)*globalNumTilesY, globalGridColor, true)

	for x := 0; x < globalNumTilesX; x++ {
		for y := 0; y < globalNumTilesY; y++ {

			var tileColor color.Color
			switch g.grid[x][y] {
			case p1Token:
				tileColor = globalTokenColors[g.p1Color]
			case p2Token:
				tileColor = globalTokenColors[g.p2Color]
			default:
				tileColor = globalBackgroundColor
			}

			vector.DrawFilledCircle(screen, float32(globalTileSize/2+x*globalTileSize), float32(globalTileSize+globalTileSize/2+y*globalTileSize+80), globalTileSize/2-globalCircleMargin, tileColor, true)
		}
	}
}
