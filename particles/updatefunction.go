package particles

import (
	"math"
	"math/rand"
	"project-particles/config"
)

func (particle *Particle) UpdatePosition() {
	if config.General.Gravity {
		particle.SpeedY += config.General.GravityCoefficient
	}
	particle.PositionX += particle.SpeedX
	particle.PositionY += particle.SpeedY
}

func (particle *Particle) IncreaseOpacity() {
	particle.Opacity = particle.Opacity + 0.1
}

func (particle *Particle) DecreaseOpacity() {
	particle.Opacity = particle.Opacity - 0.1
}

func (particle *Particle) SizeModification() {
	if config.General.RandomSize {
		particle.ScaleX = particle.ScaleX * rand.Float64()
		particle.ScaleY = particle.ScaleX * rand.Float64()
	}
	particle.ScaleX -= 0.1
	particle.ScaleY -= 0.1
}

func (particle *Particle) Multicolor() {
	particle.ColorBlue = rand.Float64()
	particle.ColorGreen = rand.Float64()
	particle.ColorRed = rand.Float64()
}

func (particle *Particle) Velocity() {
	distanceX := float64(config.General.WindowSizeX/2) - (particle.PositionX)
	distanceY := float64(config.General.WindowSizeY/2) - (particle.PositionY)
	distance := math.Sqrt((distanceX * distanceX) + (distanceY * distanceY))
	if distanceX/distance < 100 {
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
