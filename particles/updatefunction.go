package particles

import (
	"math"
	"math/rand"
	"project-particles/config"
)

// Définis la position des particules au cour du temps
func (particle *Particle) UpdatePosition() {
	// La gravité influx sur la postion Y de la particule
	if config.General.Gravity {
		particle.SpeedY += config.General.GravityCoefficient
	}
	particle.PositionX += particle.SpeedX
	particle.PositionY += particle.SpeedY
}

// Définis la position des particules
func (particle *Particle) IncreaseOpacity() {
	particle.Opacity = particle.Opacity + 0.05
}

func (particle *Particle) DecreaseOpacity() {
	particle.Opacity = particle.Opacity - 0.05
}

func (particle *Particle) SizeModification() {
	if config.General.RandomSize {
		particle.ScaleX = particle.ScaleX * rand.Float64()
		particle.ScaleY = particle.ScaleX * rand.Float64()
	}
	particle.ScaleX -= 0.1
	particle.ScaleY -= 0.1
}

// La couleur de chaque particule change en fonction du temps
func (particle *Particle) Multicolor() {
	particle.ColorBlue = rand.Float64()
	particle.ColorGreen = rand.Float64()
	particle.ColorRed = rand.Float64()
}

// Lorsque les particules spawn en formant un cercle, cette fonction définis leurs vitesses angulaires au sein du cercle.
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
