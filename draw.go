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
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Nombre de particules actuels :", g.system.Content.Len()), 0, 30)
	}
	//Permet d'afficher sur l'écran sur qu'elle config on ce situe
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		ebitenutil.DebugPrintAt(screen, ("< Config4 >"), (config.General.WindowSizeX/2)-15, 10)
	} else if ebiten.IsKeyPressed(ebiten.KeyU) {
		ebitenutil.DebugPrintAt(screen, ("< Config5 >"), (config.General.WindowSizeX/2)-15, 10)
	} else if ebiten.IsKeyPressed(ebiten.KeyX) {
		ebitenutil.DebugPrintAt(screen, ("< Config3 >"), (config.General.WindowSizeX/2)-15, 10)
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		ebitenutil.DebugPrintAt(screen, ("< Config2 >"), (config.General.WindowSizeX/2)-15, 10)
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		ebitenutil.DebugPrintAt(screen, ("< Config1 >"), (config.General.WindowSizeX/2)-15, 10)
	}
}
