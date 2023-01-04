package particles

import "project-particles/config"

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	s.buffer = s.buffer + config.General.SpawnRate
	for e := s.Content.Front(); e != nil; e = e.Next() {
		// do something with e.Value
		particule, ok := e.Value.(*Particle)
		if ok {
			particule.PositionX = particule.PositionX + particule.SpeedX + config.General.Gravity
			particule.PositionY = particule.PositionY - particule.SpeedY + particule.LifeRate*config.General.Gravity
			particule.LifeRate++
			if particule.LifeRate > 20 {
				particule.Opacity = particule.Opacity - 0.1
			}
		}
	}
	if s.buffer >= 1 {
		s.Content.PushFront(ParticuleCr())
		s.buffer--
	}
}
