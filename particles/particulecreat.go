package particles

import (
	"math/rand"
	"project-particles/config"
)

// La fonction ParticuleCR() permet de créer une particule.
// La particule peut apparaitre a une position aléatoire ou définis en fonction du paramêtre RandomSpawn.
// Si le paramêtre RandomSpawn n'est pas activé, les paramêtres SpawnX et SpawnY définissent la position de la particule.
// Il lui est attribué une vitesse définis en fonction des paramêtres MaxSpeed et TypeSpeed.
// Plus exactement, TypeSpeed permet de décidé dans qu'elle direction vont les particules.
func CreatParticle() *Particle {
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

	var particule Particle = Particle{
		PositionX: PosX,
		PositionY: PosY,
		ScaleX:    3, ScaleY: 3,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		SpeedX:  TypeSpeedX,
		SpeedY:  TypeSpeedY,
		Opacity: 1,
	}

	return &particule
}