package particles

import (
	"project-particles/config"
	"testing"
)

// Test si la fonction ParticuleCr() crée bien une particule.
func TestPartCreat(t *testing.T) {
	var listeParticule []Particle
	listeParticule = append(listeParticule, *ParticuleCr())
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
		var particule *Particle = ParticuleCr()
		if particule.PositionX > float64(config.General.WindowSizeX) || particule.PositionX < 0 || particule.PositionY > float64(config.General.WindowSizeY) || particule.PositionY < 0 {
			t.Error("Les particules ne doivent normallement pas apparaitre or de la fenêtre.")
		}
	}
}

var particule *Particle = ParticuleCr()

func TestSpawnParticule(t *testing.T) {
	config.General.RandomSpawn = false
	if particule.PositionX != float64(config.General.SpawnX) || ParticuleCr().PositionY != float64(config.General.SpawnY) {
		t.Error("Les particules doivent toutes apparaitre dans les positions définis par les paramêtres SpawnX et SpawnY, ce qui n'est pas le cas ici.")
	}
}

func TestSpeed(t *testing.T) {
	config.General.MaxSpeed = 10
	if particule.SpeedX > config.General.MaxSpeed {
		t.Error("La vitesse des particules ne devrait pas dépasser le paramêtre MaxSpeed.")
	}
}

func TestTypeSpeed(t *testing.T) {
	config.General.TypeSpeed = 1
	if particule.SpeedX < 0 || particule.SpeedY < 0 {
		t.Error("La vitesse devraient avoir une valeur négative.")
	}
}
