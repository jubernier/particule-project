package particles

import (
	"math"
	"math/rand"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

func (particle *Particle) CircleShape() {

	var axeX = config.General.WindowSizeX / 2
	var axeY = config.General.WindowSizeY / 2

	if config.General.CursorCercle {
		axeX, axeY = ebiten.CursorPosition()
	}
	var i float64 = rand.Float64()
	var e float64 = rand.Float64()
	particle.PositionX = config.General.CercleRadius * math.Cos(e*math.Pi/i)
	particle.PositionY = config.General.CercleRadius * math.Sin(e*math.Pi/i)
	particle.PositionX += float64(axeX)
	particle.PositionY += float64(axeY)
}
