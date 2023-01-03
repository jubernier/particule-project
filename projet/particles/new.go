package particles

import (
	"container/list"
	"project-particles/config"
	"math/rand"
	"time"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	l := list.New()
	var PosX = float64(config.General.SpawnX)
	var PosY = float64(config.General.SpawnY)
	rand.Seed(time.Now().UnixNano())
	for i := 0 ; i < config.General.InitNumParticles ; i++{
		if config.General.RandomSpawn{
			PosX = rand.Float64() * (float64(config.General.WindowSizeX)
			PosY = rand.Float64() * (float64(config.General.WindowSizeY)
		}
		l.PushFront(&Particle{
			PositionX: PosX,
			PositionY: PosY,
			ScaleX:    1, ScaleY: 1,
			ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
			Opacity: 1,
		})
	}

	return System{Content: l}
}
