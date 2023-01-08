package particles

import (
	"container/list"
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le retourne à la fonction principale du projet, qui se chargera de l'afficher.
// Cette fonction NewSystem permet de créer le nombre de particules qui est precisés dans le paramétre InitNumParticles qui se trouve dans le fichier config.json

func NewSystem() System {
	l := list.New()
	for i := 0; i < config.General.InitNumParticles; i++ {
		l.PushFront(CreatParticle())
	}

	return System{Content: l}
}
