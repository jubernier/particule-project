package tests

import (
	"container/list"
	"project-particles/config"
	"project-particles/particles"
	"testing"
)

// Test qui permet de vérifier si la position de la particule change en un appel de la fonction Update().
func TestUpdate(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.TypeSpeed = 1
	config.General.MaxSpeed = 10
	var l *list.List = list.New()
	l.PushFront(particles.CreatParticle())
	var SystemParticle *particles.System = &particles.System{Content: l}
	var particule *particles.Particle = l.Front().Value.(*particles.Particle)
	var posX float64 = particule.PositionX
	var posY float64 = particule.PositionY

	SystemParticle.Update()

	if posX == SystemParticle.Content.Front().Value.(*particles.Particle).PositionX || posY == SystemParticle.Content.Front().Value.(*particles.Particle).PositionY {
		t.Error("Les particules ne se déplace pas et pourtant elle devrait.")
	}
}

// Test qui permet de vérifier si il y a bien création de particules comme le définis le paramêtre SpawnRate.
func TestRandomSpawnUpdate(t *testing.T) {
	config.General.SpawnRate = 1
	var l *list.List = list.New()
	l.PushFront(particles.CreatParticle())
	var systema *particles.System = &particles.System{Content: l}
	var systemb *particles.System = &particles.System{Content: list.New()}
	systema.Update()
	if systema.Content.Len() == systemb.Content.Len() {
		t.Error("Les particules ne sont pas crées au cours du temps alors que cela devraient être possible.")
	}
}
