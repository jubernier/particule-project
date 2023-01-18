package particles

import (
	"container/list"
	"project-particles/config"
)

// Update() mets à jour l'état du système de particules à chaque pas de temps. Elle est appellée exactement 60 fois par seconde de manière régulière par la fonction principale du projet.
// A chaque appel de Update() les paricules se déplace ou non en fonction de la valeur du paramêtre MaxSpeed et des particules sont crées ou non en fonction du paramêtre SpawnRate.

var indice float64

func (s *System) Update() {
	var comptetour int
	indice = indice + config.General.SpawnRate - float64(int(config.General.SpawnRate))
	var eNext *list.Element
	for e := s.Content.Front(); e != nil && comptetour < s.Content.Len()-config.NumberDeath; e = eNext {
		eNext = e.Next()
		// do something with e.Value
		particule, ok := e.Value.(*Particle)

		if ok {
			//définis la vitesse des particules
			particule.UpdatePosition()

			if config.General.CercleMouvement {
				//définis la vitesse angulaire des particules pour créer un effet de tourbillon qui vas vers le centre de cercle.
				particule.Velocity()
			}
			if config.General.Multicolor {
				//modifie la couleur de chaque particule en fonction du temps
				particule.Multicolor()
			}
			if config.General.Margin {
				//vérifie si la particule ne dépasse pas les marges imposées
				if particule.PositionX <= -config.General.MargeCoefficient ||
					particule.PositionY <= -config.General.MargeCoefficient ||
					particule.PositionX >= float64(config.General.WindowSizeX)+config.General.MargeCoefficient ||
					particule.PositionY >= float64(config.General.WindowSizeY)+config.General.MargeCoefficient {
					config.NumberDeath++
				}
			}

			if config.General.ChangeSize {
				particule.SizeModification()
			}

			if config.General.OpacityRate == 1 {
				particule.IncreaseOpacity()
			} else if config.General.OpacityRate == 2 {
				particule.DecreaseOpacity()
			}

			if config.General.AddLifeToParticle {
				//la vie de la particule augmente
				particule.LifeRate++
				if particule.LifeRate >= config.General.LifeRate {
					//Si le taux de vie de la particule a dépasser le niveau de vie définis dans les paramêtres, elle meurt et on la mets a la fin de la liste pour réutilier son espace mémoire.
					s.Content.MoveToBack(e)
					config.NumberDeath++
				}
			}
			comptetour++
		}
	}
	//Des particules sont créer en fonction du paramêtre SpawnRate
	for i := 0; i < int(config.General.SpawnRate); i++ {
		if config.NumberDeath >= 1 {
			//Ici, si il y a particule morte alors il y a recyclage, on la régénere.
			RecycleParticule(s)
		} else {
			s.Content.PushFront(CreatParticle())
		}
	}
	// créer une particule en fonction du paramêtre spawnrate si il n'est pas un entier
	if indice >= 1 {
		s.Content.PushFront(CreatParticle())
		indice--
	}
}
