package particles

import (
	"container/list"
	"project-particles/config"

	"testing"
)

// Test si la particule regénéré apparait bien a la position définis par SpawnX et SpawnY quand les paramêtres qui definissent les positions ne sont pas activés.
func TestRegenerateParticleSpawn(t *testing.T) {
	config.General.RandomSpawn = false
	config.General.CercleShape = false
	var l *list.List = list.New()
	l.PushFront(CreatParticle())
	var SystemBis *System = &System{Content: l}
	var particule *Particle = l.Front().Value.(*Particle)
	RecycleParticule(SystemBis)
	if particule.PositionX != float64(config.General.SpawnX) || CreatParticle().PositionY != float64(config.General.SpawnY) {
		t.Error()
	}
}

// Test qui permet de vérifier si les particules ne dépassent pas le paramétre MaxSpeed
func TestRegenerateParticleSpeed(t *testing.T) {
	config.General.TypeSpeed = 2
	config.General.MaxSpeed = 10
	var l *list.List = list.New()
	l.PushFront(CreatParticle())
	var SystemTest *System = &System{Content: l}
	var testParticle *Particle = l.Front().Value.(*Particle)
	RecycleParticule(SystemTest)
	if testParticle.SpeedX > config.General.MaxSpeed || testParticle.SpeedY > config.General.MaxSpeed {
		t.Error("La vitesse des particules ne devrait pas dépasser le paramêtre MaxSpeed.")
	}
}

// Test qui permet de vérifier que la vitesse a une valeur négative lorsque le paramêtre TypeSpeed = 1.
func TestRegenerateParticleTypeSpeed(t *testing.T) {
	config.General.TypeSpeed = 1
	var l *list.List = list.New()
	l.PushFront(CreatParticle())
	var SystemTestTypeSpeed *System = &System{Content: l}
	var testParticleTypeSpeed *Particle = l.Front().Value.(*Particle)
	RecycleParticule(SystemTestTypeSpeed)
	if testParticleTypeSpeed.SpeedX > 0 || testParticleTypeSpeed.SpeedY > 0 {
		t.Error("La vitesse devraient avoir une valeur négative.")
	}
}
