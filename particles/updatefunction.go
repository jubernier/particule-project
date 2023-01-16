package particles

import(
	"project-particles/config"
)

func (particle *Particle) UpdatePosition() {
	if config.General.Gravity {
		particle.SpeedY +=config.General.GravityCoefficient
	}
	particle.PositionX += particle.SpeedX
	particle.PositionY += particle.SpeedY
	particle.LifeRate++
}

