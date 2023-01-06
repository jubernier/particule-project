package particles

import (
	"fmt"
	"project-particles/config"
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

//Test si le spawn n'est pas en dehors de l'écran
func TestScreenRandomSpawn(t *testing.T) {
	var particule *Particle = ParticuleCr()
	fmt.Println(particule.PositionX)
	fmt.Println(particule.PositionY)
	if particule.PositionX > float64(config.General.WindowSizeX) || particule.PositionX < 0 || particule.PositionY > float64(config.General.WindowSizeY) || particule.PositionY < 0 {
		t.Error("Les particules ne doivent normallement pas apparaitre or de la fenêtre.")
	}
}
