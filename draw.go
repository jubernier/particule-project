package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules g.system. Elle est appelée automatiquement environ 60 fois par seconde par la bibliothèque Ebiten.
var ConfigSelector int

func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}

	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
		ebitenutil.DebugPrintAt(screen, fmt.Sprint(g.system.Content.Len()), 0, 30)
	}
	var MaxNbConfig = 2
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if ConfigSelector == MaxNbConfig {
			ConfigSelector = 0
		} else {
			ConfigSelector++
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if ConfigSelector == 0 {
			ConfigSelector = MaxNbConfig
		} else {
			ConfigSelector--
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		ConfigSelector = -1
	}
	switch ConfigSelector {
	case 0:
		config.Get("config/config1.json")
		ebitenutil.DebugPrintAt(screen, ("< Config1 >"), (config.General.WindowSizeX/2)-15, 10)
	case 1:
		config.Get("config/config2.json")
		ebitenutil.DebugPrintAt(screen, ("< Config2 >"), (config.General.WindowSizeX/2)-15, 10)
	case 2:
		config.Get("config/config3.json")
		ebitenutil.DebugPrintAt(screen, ("< Config3 >"), (config.General.WindowSizeX/2)-15, 10)
	case -1:
		config.Get("config/config4.json")
		ebitenutil.DebugPrintAt(screen, ("< Config4 >"), (config.General.WindowSizeX/2)-15, 10)
	}

	if ebiten.IsKeyPressed(ebiten.KeyP) {
		if config.General.ColorRandom {
			config.General.ColorRandom = false
			config.General.ColorRed, config.General.ColorBlue, config.General.ColorGreen = 1, 1, 0 //Color purple
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyA) { //La touche A représente la touche Q en systeme européen
		if config.General.ColorRandom {
			config.General.ColorRandom = false
			config.General.ColorRed, config.General.ColorBlue, config.General.ColorGreen = 0, 1, 1 //Color Aqua
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyY) {
		if config.General.ColorRandom {
			config.General.ColorRandom = false
			config.General.ColorRed, config.General.ColorBlue, config.General.ColorGreen = 1, 0, 1 //Color Yellow
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyR) {
		if config.General.ColorRandom {
			config.General.ColorRandom = false
			config.General.ColorRed, config.General.ColorBlue, config.General.ColorGreen = 1, 0, 0 //Color Red
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyB) {
		if config.General.ColorRandom {
			config.General.ColorRandom = false
			config.General.ColorRed, config.General.ColorBlue, config.General.ColorGreen = 0, 1, 0 //Color Blue
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyG) {
		if config.General.ColorRandom {
			config.General.ColorRandom = false
			config.General.ColorRed, config.General.ColorBlue, config.General.ColorGreen = 0, 0, 1 //Color Green
		}
	}
}

// func Mouvement() {
// 	if config.General.MouseMouvement {
// 		// Déplacement à la souris
// 		mx, my := ebiten.CursorPosition()

// 		config.General.SpawnX, config.General.SpawnY = float64(mx), float64(my)

// 	} else {
// 		// Déplacement au clavier
// 		if ebiten.IsKeyPressed(ebiten.KeyU) {
// 			config.General.SpawnY--
// 		}
// 		if ebiten.IsKeyPressed(ebiten.KeyO) {
// 			config.General.SpawnY++
// 		}
// 		if ebiten.IsKeyPressed(ebiten.KeyL) {
// 			config.General.SpawnX--
// 		}
// 		if ebiten.IsKeyPressed(ebiten.KeyX) {
// 			config.General.SpawnX++
// 		}
// 	}
// }

//Essayer de fixer le beug de tourbillon avec les touches dynamiques  chose que j'essaye de faire il y'a 3heures quelle honte !!
// Remettre le cercle au centre de l'ecran a chaque fois
// Essayer de refaire ca sur Update
// pourquoi ma fonction mouvement ne marche pas  + regler ca tomorrow
// Faire le Redmi et terminer les tests
