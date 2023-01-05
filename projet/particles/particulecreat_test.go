package particles

import (
	"fmt"
	"testing"
)

//Test si la fonction ParticuleCr() crée bien une particule

func TestPartCreat(t *testing.T) {
	var listeParticule []Particle
	listeParticule = append(listeParticule, *ParticuleCr())
	fmt.Println(listeParticule)
	if len(listeParticule) != 1 {
		t.Error()
	}
}
