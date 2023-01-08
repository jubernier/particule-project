package particles

import (
	"project-particles/config"
	"testing"
)

// Le test ci-dessous sert a vérifier que new.go créer bien un nombre de particules équivalant a InitNumParticle
func TestNbParticule(t *testing.T) {
	if config.General.InitNumParticles != NewSystem().Content.Len() {
		t.Error("Le nombre de particules créés ne corresponds pas a la valeur du paramêtre InitNumParticle.")
	}
}
