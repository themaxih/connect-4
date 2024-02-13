package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

// Constantes définissant les paramètres généraux du programme.
const (
	globalWidth         = globalNumTilesX * globalTileSize
	globalHeight        = (globalNumTilesY + 1) * globalTileSize + 80
	globalTileSize      = 100
	globalNumTilesX     = 7
	globalNumTilesY     = 6
	globalCircleMargin  = 5
	globalBlinkDuration = 60
	globalNumColorLine  = 3
	globalNumColorCol   = 3
	globalNumColor      = globalNumColorLine * globalNumColorCol
)

// Variables définissant les paramètres généraux du programme.
var (
	globalBackgroundColor color.Color = color.NRGBA{R: 41, G: 43, B: 47, A: 255}
	globalGridColor       color.Color = color.NRGBA{R: 64, G: 68, B: 75, A: 255}
	globalTextColor       color.Color = color.NRGBA{R: 250, G: 250, B: 250, A: 255}
	globalSelectColor     color.Color = color.NRGBA{R: 150, G: 150, B: 150, A: 255}
	globalCrossColor	  color.Color = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	smallFont             font.Face
	largeFont             font.Face
	globalTokenColors     [globalNumColor]color.Color = [globalNumColor]color.Color{
		color.NRGBA{R: 255, G: 0, B: 0, A: 255},
		color.NRGBA{R: 255, G: 128, B: 0, A: 255},
		color.NRGBA{R: 255, G: 160, B: 193, A: 255},
		color.NRGBA{R: 124, G: 252, B: 0, A: 255},
		color.NRGBA{R: 30, G: 144, B: 255, A: 255},
		color.NRGBA{R: 169, G: 67, B: 15, A: 255},
		color.NRGBA{R: 255, G: 238, B: 0, A: 255},
		color.NRGBA{R: 153, G: 50, B: 204, A: 255},
		color.NRGBA{R: 255, G: 255, B: 255, A: 255},
	}
	offScreenImage *ebiten.Image
)
