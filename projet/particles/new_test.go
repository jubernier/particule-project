package particles

import (
	"project-particles/config"
	"testing"
)

func TestNbParticule(t *testing.T) {
	if config.General.InitNumParticles != l.Len() {
		t.Error()
	}
}
