package particles

import (
	"math"
	"project-particles/config"
)

func (particle *Particle) Cercle() {
	particle.PositionX = config.General.CercleRadius * math.Cos(math.Pi/config.General.Angle)
	particle.PositionY = config.General.CercleRadius * math.Cos(math.Pi/config.General.Angle)
}
