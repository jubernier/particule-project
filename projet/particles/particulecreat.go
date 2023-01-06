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
	var particule Particle = Particle{
		PositionX: PosX,
		PositionY: PosY,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:  1,
		LifeRate: 0,
	}

	return &particule
}
