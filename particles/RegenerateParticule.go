package particles

import (
	"math"
	"math/rand"
	"project-particles/config"
	"time"
)

func AddParticule(s *System) {
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
		var axeX = config.General.WindowSizeX / 2
		var axeY = config.General.WindowSizeY / 2
		var i float64 = rand.Float64()
		var e float64 = rand.Float64()
		PosX = config.General.CercleRadius * math.Cos(e*math.Pi/i)
		PosY = config.General.CercleRadius * math.Sin(e*math.Pi/i)
		PosX = PosX + float64(axeX)
		PosY = PosY + float64(axeY)
	}

	e := s.Content.Back()
	particule, _ := e.Value.(*Particle)
	particule.ColorGreen = 1
	particule.ColorRed = 1
	particule.Opacity = 1
	particule.SpeedX = TypeSpeedX
	particule.SpeedY = TypeSpeedY
	particule.PositionX = PosX
	particule.PositionY = PosY
	particule.LifeRate = 0
	s.Content.MoveToFront(e)
}
