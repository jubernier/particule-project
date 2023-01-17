package main

import (
	"log"
	"math/rand"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// main est la fonction principale du projet.Elle initialise la fenêtre d'affichage, puis elle crée un système de particules encapsulé dans un "game" et appelle la fonction RunGame qui se charge de faire les mise-à-jour (Update) et affichages (Draw) de manière  régulière.
func main() {
	rand.Seed(time.Now().UnixNano())
	config.Get("config/config1.json")
	assets.Get()

	ebiten.SetWindowTitle(config.General.WindowTitle)
	ebiten.SetWindowSize(config.General.WindowSizeX, config.General.WindowSizeY)

	g := game{system: particles.NewSystem()}

	err := ebiten.RunGame(&g)
	if err != nil {
		log.Print(err)
	}
}
