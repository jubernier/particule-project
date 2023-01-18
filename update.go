package main

import (
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {
	g.system.Update()
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
	}
	config.NumberDeath++
	switch ConfigSelector {
	case 0:
		config.Get("config/config1.json")
		g.system = particles.NewSystem()
	case 1:
		config.Get("config/config2.json")
	case 2:
		config.Get("config/config3.json")
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
	return nil
}
