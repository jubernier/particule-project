package particles

import (
	"project-particles/config"
)

// Update() mets à jour l'état du système de particules à chaque pas de temps. Elle est appellée exactement 60 fois par seconde de manière régulière par la fonction principale du projet.
// A chaque appel de Update() les paricules se déplace ou non en fonction de la valeur du paramêtre MaxSpeed et des particules sont crées ou non en fonction du paramêtre SpawnRate.

var indice float64

func (s *System) Update() {
	var comptetour int
	indice = indice + config.General.SpawnRate - float64(int(config.General.SpawnRate))
	for e := s.Content.Front(); e != nil && comptetour < s.Content.Len()-config.NumberDeath; e = e.Next() {
		// do something with e.Value
		particule, ok := e.Value.(*Particle)
		if ok {

			particule.UpdatePosition()

			if config.General.CercleSpeed {
				particule.Velocity()
			}

			if config.General.Margin {
				if particule.PositionX >= float64(config.General.WindowSizeX) || particule.PositionX < 0 || particule.PositionY > float64(config.General.WindowSizeY) {
					go s.Content.Remove(e)
					config.NumberDeath++
				}
			}

			if config.General.Design {
				particule.DesignParticle()
			}

			if config.General.OpacityRate == 1 {
				particule.IncreaseOpacity()
			} else if config.General.OpacityRate == 2 {
				particule.DecreaseOpacity()
			}

			if config.General.AddLifeToParticle {
				if particule.LifeRate >= config.General.LifeRate {
					s.Content.MoveToBack(e)
					config.NumberDeath++
	
				}
			}
			comptetour++
		}
	}
	for i := 0; i < int(config.General.SpawnRate); i++ {
		if config.NumberDeath >= 1 {
			RecycleParticule(s)
		} else {
			s.Content.PushFront(CreatParticle())
		}
	}
	if indice >= 1 {
		s.Content.PushFront(CreatParticle())
		indice--
	}

}
