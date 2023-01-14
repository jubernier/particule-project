package particles

import (
	"math"
	"math/rand"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

func (particle *Particle) DesignParticle() {
	particle.ColorBlue = 0
	particle.ScaleX = 1.5
	particle.Rotation = 1.5
}

func (particle *Particle) IncreaseOpacity() {
	particle.Opacity = particle.Opacity + 0.1
}

func (particle *Particle) DecreaseOpacity() {
	particle.Opacity = particle.Opacity - 0.1
}

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
	particle.PositionX = particle.PositionX + float64(axeX)
	particle.PositionY = particle.PositionY + float64(axeY)
}

func (particle *Particle) Velocity() {
	distanceX := float64(config.General.WindowSizeX/2) - (particle.PositionX)
	distanceY := float64(config.General.WindowSizeY/2) - (particle.PositionY)
	distance := math.Sqrt((distanceX * distanceX) + (distanceY * distanceY))
	if distanceX/distance < 100.0 {
		particle.SpeedX += distanceX / distance
	}
	if distanceY/distance < 100 {
		particle.SpeedY += distanceY / distance
	}
	if particle.ExDistance > distance {
		particle.SpeedX -= distanceX / distance
		particle.SpeedY -= distanceY / distance
	}
	particle.ExDistance = distance
}
