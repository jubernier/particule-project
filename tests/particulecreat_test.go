package tests

import (
	"project-particles/config"
	"project-particles/particles"
	"testing"
)

// Test pour vérifier que la particule a bien été créer.
func TestPartCreat(t *testing.T) {
	var listeParticule []particles.Particle
	listeParticule = append(listeParticule, *particles.CreatParticle())
	if len(listeParticule) != 1 {
		t.Error("La fonction ne crée pas une particule et pourtant elle devrait.")
	}
}

// Test si il y a bien des particules qui spawn dans la fenêtre.
func TestScreenRandomSpawn(t *testing.T) {
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.RandomSpawn = true
	for i := 0; i < 10000; i++ {
		var particule *particles.Particle = particles.CreatParticle()
		if particule.PositionX > float64(config.General.WindowSizeX) || particule.PositionX < 0 || particule.PositionY > float64(config.General.WindowSizeY) || particule.PositionY < 0 {
			t.Error("Les particules ne doivent normallement pas apparaitre or de la fenêtre.")
		}
	}
}

// Test qui permet de verifier si les particules apparaissent dans les positions définis par les paramétres SpawnX et SpawnY
func TestSpawnParticule(t *testing.T) {
	config.General.RandomSpawn = false
	var particule *particles.Particle = particles.CreatParticle()
	if particule.PositionX != float64(config.General.SpawnX) || particles.CreatParticle().PositionY != float64(config.General.SpawnY) {
		t.Error("Les particules doivent toutes apparaitre dans les positions définis par les paramêtres SpawnX et SpawnY, ce qui n'est pas le cas ici.")
	}
}

// Test qui permet de vérifier si les particules ne dépassent pas le paramétre MaxSpeed
func TestSpeed(t *testing.T) {
	config.General.TypeSpeed = 2
	config.General.MaxSpeed = 10

	var testParticle *particles.Particle = particles.CreatParticle()

	if testParticle.SpeedX > config.General.MaxSpeed || testParticle.SpeedY > config.General.MaxSpeed {
		t.Error("La vitesse des particules ne devrait pas dépasser le paramêtre MaxSpeed.")
	}
}

// Test qui permet de vérifier que la vitesse a une valeur négative lorsque le paramêtre TypeSpeed = 1
func TestTypeSpeed(t *testing.T) {
	config.General.TypeSpeed = 1

	var testparticle *particles.Particle = particles.CreatParticle()

	if testparticle.SpeedX > 0 || testparticle.SpeedY > 0 {
		t.Error("La vitesse devraient avoir une valeur négative.")
	}
}
