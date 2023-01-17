package tests

import (
	"container/list"
	"project-particles/config"
	"project-particles/particles"
	"testing"
)

func TestRegenerateParticle(t *testing.T) {
	config.NumberDeath = 1
	var l *list.List = list.New()
	l.PushFront(particles.CreatParticle())
	var SystemBis *particles.System = &particles.System{Content: l}
	var particule *particles.Particle = l.Front().Value.(*particles.Particle)
	particule.RecycleParticule()
	if particule.LifeRate == 0 {
		t.Error()
	}
}
