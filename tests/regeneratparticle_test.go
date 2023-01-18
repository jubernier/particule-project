package tests

import (
	"container/list"
	"project-particles/config"
	"project-particles/particles"
	"testing"
)

// Test si la particule regénéré apparait bien a la position définis par SpawnX et SpawnY quand les paramêtres qui definissent les positions ne sont pas activés.
func TestRegenerateParticleSpawn(t *testing.T) {
	config.General.RandomSpawn = false
	config.General.CercleShape = false
	var l *list.List = list.New()
	l.PushFront(particles.CreatParticle())
	var SystemBis *particles.System = &particles.System{Content: l}
	var particule *particles.Particle = l.Front().Value.(*particles.Particle)
	particles.RecycleParticule(SystemBis)
	if particule.PositionX != float64(config.General.SpawnX) || particles.CreatParticle().PositionY != float64(config.General.SpawnY) {
		t.Error()
	}
}

/*
func TestRegenerateParticleCercle(t *testing.T) {
	config.General.CursorCercle = true
	axeX, axeY = ebiten.CursorPosition()

	config.General.CercleShape = true
	var l *list.List = list.New()
	l.PushFront(particles.CreatParticle())
	var SystemBis *particles.System = &particles.System{Content: l}
	var particule *particles.Particle = l.Front().Value.(*particles.Particle)
	particles.RecycleParticule(SystemBis)
	if particule.PositionX != particule.PositionX + regenerateparticles.CircleShape().axeX {
		t.Error()
	}
}
*/

// Test qui permet de vérifier si les particules ne dépassent pas le paramétre MaxSpeed
func TestRegenerateParticleSpeed(t *testing.T) {
	config.General.TypeSpeed = 2
	config.General.MaxSpeed = 10
	var l *list.List = list.New()
	l.PushFront(particles.CreatParticle())
	var SystemTest *particles.System = &particles.System{Content: l}
	var testParticle *particles.Particle = l.Front().Value.(*particles.Particle)
	particles.RecycleParticule(SystemTest)
	if testParticle.SpeedX > config.General.MaxSpeed || testParticle.SpeedY > config.General.MaxSpeed {
		t.Error("La vitesse des particules ne devrait pas dépasser le paramêtre MaxSpeed.")
	}
}

// Test qui permet de vérifier que la vitesse a une valeur négative lorsque le paramêtre TypeSpeed = 1.
func TestRegenerateParticleTypeSpeed(t *testing.T) {
	config.General.TypeSpeed = 1
	var l *list.List = list.New()
	l.PushFront(particles.CreatParticle())
	var SystemTestTypeSpeed *particles.System = &particles.System{Content: l}
	var testParticleTypeSpeed *particles.Particle = l.Front().Value.(*particles.Particle)
	particles.RecycleParticule(SystemTestTypeSpeed)
	if testParticleTypeSpeed.SpeedX > 0 || testParticleTypeSpeed.SpeedY > 0 {
		t.Error("La vitesse devraient avoir une valeur négative.")
	}
}
