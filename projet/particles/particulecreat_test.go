package particles

import (
	"testing"
)

//Test si la fonction ParticuleCr() crée bien une particule

func TestPartCreat(t *testing.T) {
	var listeParticule []Particle
	listeParticule = append(listeParticule, *ParticuleCr())
	if len(listeParticule) != 1 {
		t.Error()
	}
}
