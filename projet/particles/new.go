package particles

import (
	"container/list"
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	l := list.New()
	l.PushFront(&Particle{
		PositionX: float64(config.General.WindowSizeX) / 2,
		PositionY: float64(config.General.WindowSizeY) / 2,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity: 1,
	})
	return System{Content: l}
}
