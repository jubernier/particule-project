package main

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {
	// Permet de passer d'une config à une autre .
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		config.Get("config/config4.json")
	} else if ebiten.IsKeyPressed(ebiten.KeyU) {
		config.Get("config/config5.json")
	} else if ebiten.IsKeyPressed(ebiten.KeyX) {
		config.Get("config/config3.json")
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		config.Get("config/config2.json")
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		config.Get("config/config1.json")
	}
	// Permet de modifier les couleurs en fonctions des touches
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
	g.system.Update()
	return nil
}
