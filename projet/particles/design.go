package particles

func (particle *Particle) DesignParticle() {
	particle.ColorBlue = 0
	particle.ScaleX = 1.5
}

func (particle *Particle) IncreaseOpacity() {
	particle.Opacity = particle.Opacity + 0.1
}

func (particle *Particle) DecreaseOpacity() {
	particle.Opacity = particle.Opacity - 0.1
}
