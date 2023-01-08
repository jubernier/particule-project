package particles

import (
	"container/list"
	"project-particles/config"
	"testing"
)

func TestUpdate(t *testing.T) {
	//Test par exemple si position de x a bougé
	//là je ne test pas que la particule a bougé par rapport a la position initiale
	//sauf que rand des fois a 0
	config.General.InitNumParticles = 1
	config.General.TypeSpeed = 1
	config.General.MaxSpeed = 10
	var l *list.List = list.New()
	l.PushFront(ParticuleCr())
	var SystemParticle *System = &System{Content: l}
	var particule *Particle = l.Front().Value.(*Particle)
	var posX float64 = particule.PositionX
	var posY float64 = particule.PositionY

	SystemParticle.Update()

	if posX == SystemParticle.Content.Front().Value.(*Particle).PositionX || posY == SystemParticle.Content.Front().Value.(*Particle).PositionY {
		t.Error("Les particules ne se déplace pas et pourtant elle devrait.")
	}
}
