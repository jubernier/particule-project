package particles

import (
	"math"
	"math/rand"
	"project-particles/config"
	"time"
)

// La fonction ParticuleCR() permet de créer une particule.
// La particule peut apparaitre a une position aléatoire ou définis en fonction du paramêtre RandomSpawn.
// Si le paramêtre RandomSpawn n'est pas activé, les paramêtres SpawnX et SpawnY définissent la position de la particule.
// Il lui est attribué une vitesse définis en fonction des paramêtres MaxSpeed et TypeSpeed.
// Plus exactement, TypeSpeed permet de décidé dans qu'elle direction vont les particules.
func CreatParticle() *Particle {
	rand.Seed(time.Now().UnixNano())
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

	if config.General.CercleShape {
		var axeX int = config.General.WindowSizeX / 2
		var axeY int = config.General.WindowSizeY / 2
		var i float64 = rand.Float64()
		var e float64 = rand.Float64()
		PosX = config.General.CercleRadius * math.Cos(e*math.Pi/i)
		PosY = config.General.CercleRadius * math.Sin(e*math.Pi/i)
		PosX = PosX - float64(axeX)
		PosY = PosY - float64(axeY)
	}

	var particule Particle = Particle{
		PositionX: PosX,
		PositionY: PosY,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 0, ColorBlue: 1,
		SpeedX:   TypeSpeedX,
		SpeedY:   TypeSpeedY,
		Opacity:  1,
		LifeRate: 0,
	}

	return &particule
}
