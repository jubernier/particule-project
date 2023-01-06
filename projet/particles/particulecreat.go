package particles

import (
	"math/rand"
	"project-particles/config"
)

func ParticuleCr() *Particle {
	PosX := float64(config.General.SpawnX)
	PosY := float64(config.General.SpawnY)
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

	var particule Particle = Particle{
		PositionX: PosX,
		PositionY: PosY,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		SpeedX:   TypeSpeedX,
		SpeedY:   TypeSpeedY,
		Opacity:  1,
		LifeRate: 0,
	}

	return &particule
}
