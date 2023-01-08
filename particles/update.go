package particles

import "project-particles/config"

// Update() mets à jour l'état du système de particules à chaque pas de temps. Elle est appellée exactement 60 fois par seconde de manière régulière par la fonction principale du projet.
// A chaque appel de Update() les paricules se déplace ou non en fonction de la valeur du paramêtre MaxSpeed et des particules sont crées ou non en fonction du paramêtre SpawnRate.

var indice float64

func (s *System) Update() {
	indice = indice + config.General.SpawnRate - float64(int(config.General.SpawnRate))
	for e := s.Content.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		particule, ok := e.Value.(*Particle)
		if ok {
			particule.PositionX = particule.PositionX + particule.SpeedX
			particule.PositionY = particule.PositionY + particule.SpeedY
		}
	}
	for i := 0; i < int(config.General.SpawnRate); i++ {
		s.Content.PushFront(CreatParticle())
	}
	if indice >= 1 {
		s.Content.PushFront(CreatParticle())
		indice--
	}
}
