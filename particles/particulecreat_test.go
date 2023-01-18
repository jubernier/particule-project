package particles

import (
	"fmt"
	"math"
	"project-particles/config"

	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

// Test pour vérifier que la particule a bien été créer.
func TestPartCreat(t *testing.T) {
	var listeParticule []Particle
	listeParticule = append(listeParticule, *CreatParticle())
	if listeParticule == nil {
		t.Error("La fonction ne crée pas une particule et pourtant elle devrait.")
	}
}

// Test si les particules spawn bien dans la fenêtre quand elles spawn a des endroits aléatoires.
func TestScreenRandomSpawn(t *testing.T) {
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.RandomSpawn = true
	for spawninscreen := 0; spawninscreen < 10000; spawninscreen++ {
		var particule *Particle = CreatParticle()
		if particule.PositionX > float64(config.General.WindowSizeX) || particule.PositionX < 0 || particule.PositionY > float64(config.General.WindowSizeY) || particule.PositionY < 0 {
			t.Error("Les particules ne doivent normallement pas apparaitre or de la fenêtre.")
		}
	}
}

// Test qui permet de verifier si les particules apparaissent dans les positions définis par les paramétres SpawnX et SpawnY
func TestSpawnParticule(t *testing.T) {
	config.General.RandomSpawn = false
	var particule *Particle = CreatParticle()
	if particule.PositionX != float64(config.General.SpawnX) || CreatParticle().PositionY != float64(config.General.SpawnY) {
		t.Error("Les particules doivent toutes apparaitre dans les positions définis par les paramêtres SpawnX et SpawnY, ce qui n'est pas le cas ici.")
	}
}

// Test qui permet de vérifier si les particules ne dépassent pas le paramétre MaxSpeed.
func TestSpeed(t *testing.T) {
	config.General.TypeSpeed = 2
	config.General.MaxSpeed = 10

	var testParticle *Particle = CreatParticle()

	if testParticle.SpeedX > config.General.MaxSpeed || testParticle.SpeedY > config.General.MaxSpeed {
		t.Error("La vitesse des particules ne devrait pas dépasser le paramêtre MaxSpeed.")
	}
}

// Test qui permet de vérifier que la vitesse a une valeur négative lorsque le paramêtre TypeSpeed = 1
func TestTypeSpeed(t *testing.T) {
	config.General.TypeSpeed = 1

	var testparticle *Particle = CreatParticle()

	if testparticle.SpeedX > 0 || testparticle.SpeedY > 0 {
		t.Error("La vitesse devraient avoir une valeur négative.")
	}
}

// Test si les particules apparraissent bien de manière a former un cercle.
func TestCercleShape(t *testing.T) {
	config.General.CercleShape = true
	config.General.MoveCursor = false
	config.General.CercleRadius = 40
	var distance = config.General.CercleRadius
	var formuleX = distance * math.Cos(1*math.Pi/1)
	var formuleY = distance * math.Sin(1*math.Pi/1)

	for spawnincercle := 0; spawnincercle < 1000; spawnincercle++ {
		var particle *Particle = CreatParticle()
		particle.PositionX -= float64(config.General.WindowSizeX) / 2
		particle.PositionY -= float64(config.General.WindowSizeX) / 2
		if particle.PositionX > formuleX && particle.PositionY > formuleY {
			fmt.Println(formuleX)
			fmt.Println(particle.PositionX)
			t.Error()
		}
	}
}

// Test si le curseur est bien an haut à gauche de l'écran par défaut.
func TestCursorCercleInitiale(t *testing.T) {
	config.General.MoveCursor = false
	axeX, axeY := ebiten.CursorPosition()
	if axeX != 0 && axeY != 0 {
		t.Error("Le curseur est par défaut a la position X et position Y : 0, 0")
	}
}
