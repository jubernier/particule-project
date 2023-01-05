package particles

import "project-particles/config"

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
var buffer float64

func (s *System) Update() {
	buffer = buffer + config.General.SpawnRate - float64(int(config.General.SpawnRate))
	for e := s.Content.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		particule, ok := e.Value.(*Particle)
		if ok {
			particule.SpeedY = particule.SpeedY + config.General.Gravity
			particule.PositionX = particule.PositionX + particule.SpeedX
			particule.PositionY = particule.PositionY + particule.SpeedY
			particule.LifeRate++

			if config.General.Design {
				particule.DesignParticle()
			}
			if config.General.Margin {
				if particule.PositionX >= float64(config.General.WindowSizeX) || particule.PositionX < 0 || particule.PositionY > float64(config.General.WindowSizeY) {
					particule.PositionX = float64(config.General.WindowSizeX) + 100
					go s.Content.Remove(e)
				}
			}
			if particule.LifeRate >= config.General.LifeRate {
				s.Content.Remove(e)
			}
		}
	}
	for i := 0; i < int(config.General.SpawnRate); i++ {
		s.Content.PushFront(ParticuleCr())
	}
	if buffer >= 1 {
		s.Content.PushFront(ParticuleCr())
		buffer--
	}
}
