package particles

import (
	"math/rand"
	"project-particles/config"
)

func RecycleParticule(s *System) {

	var	colorblue = config.General.ColorBlue
	var	colorgreen = config.General.ColorGreen
	var	colorred = config.General.ColorRed
	if config.General.ColorRandom {
		colorblue = rand.Float64()
		colorgreen = rand.Float64()
		colorred = rand.Float64()
	}
	
	var PosX float64 = float64(config.General.SpawnX)
	var PosY float64 = float64(config.General.SpawnY)
	if config.General.RandomSpawn {
		PosX = rand.Float64() * float64(config.General.WindowSizeX)
		PosY = rand.Float64() * float64(config.General.WindowSizeY)
	}

	var indice []int = []int{1, -1}
	var TypeSpeedX = rand.Float64() * config.General.MaxSpeed
	var TypeSpeedY = rand.Float64() * config.General.MaxSpeed

	if config.General.TypeSpeed == 0 {
		TypeSpeedX = TypeSpeedX * float64(indice[rand.Intn(2)])
		TypeSpeedY = TypeSpeedY * float64(indice[rand.Intn(2)])
	} else if config.General.TypeSpeed == 1 {
		TypeSpeedX = -(rand.Float64()) * config.General.MaxSpeed
		TypeSpeedY = -(rand.Float64()) * config.General.MaxSpeed
	}

	e := s.Content.Back()
	particule, _ := e.Value.(*Particle)
	particule.ColorBlue = colorblue
	particule.ColorGreen = colorgreen
	particule.ColorRed = colorred
	particule.Opacity = 1
	particule.SpeedX = TypeSpeedX
	particule.SpeedY = TypeSpeedY
	particule.PositionX = PosX
	particule.PositionY = PosY
	particule.LifeRate = 0
	s.Content.MoveToFront(e)

	if config.General.CercleShape {
		particule.CircleShape()
	}
}
