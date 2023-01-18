package particles

import (
	"math"
	"math/rand"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

func (particle *Particle) UpdatePosition() {
	if config.General.Gravity {
		particle.SpeedY += config.General.GravityCoefficient
	}
	particle.PositionX += particle.SpeedX
	particle.PositionY += particle.SpeedY
}

func (particle *Particle) DesignParticle() {
	particle.ColorBlue = 0
	particle.ScaleX = 1.5
	particle.Rotation = 1.5
}

func (particle *Particle) IncreaseOpacity() {
	particle.Opacity = particle.Opacity + 0.1
}

func (particle *Particle) DecreaseOpacity() {
	particle.Opacity = particle.Opacity - 0.1
}

func (particle *Particle) CircleShape() {

	var axeX = config.General.WindowSizeX / 2
	var axeY = config.General.WindowSizeY / 2

	if config.General.CursorCercle {
		axeX, axeY = ebiten.CursorPosition()
	}
	var i float64 = rand.Float64()
	var e float64 = rand.Float64()
	particle.PositionX = config.General.CercleRadius * math.Cos(e*math.Pi/i)
	particle.PositionY = config.General.CercleRadius * math.Sin(e*math.Pi/i)
	particle.PositionX += float64(axeX)
	particle.PositionY += float64(axeY)
}

func (particle *Particle) Velocity() {
	distanceX := float64(config.General.WindowSizeX/2) - (particle.PositionX)
	distanceY := float64(config.General.WindowSizeY/2) - (particle.PositionY)
	distance := math.Sqrt((distanceX * distanceX) + (distanceY * distanceY))
	if distanceX/distance < 100 {
		particle.SpeedX += distanceX / distance
	}
	if distanceY/distance < 100 {
		particle.SpeedY += distanceY / distance
	}
	if particle.ExDistance > distance {
		particle.SpeedX -= distanceX / distance
		particle.SpeedY -= distanceY / distance
	}
	particle.ExDistance = distance
}

// func (p *Particle) Explose(s *System, i int) {
// 	// La méthode Explose() vérifie si une particule est "explosive" ou non et crée en conséquences
// 	// le nombre de particules définit dans le fichier config.json
// 	if Particle.IsExplosive {
// 		for j := 0; j < config.General.NbExplose; j++ {
// 			s.add(false, i)
// 		}
// 		if !config.General.Trainee {
// 			p.Kill()
// 		}
// 	}
// }

// func (p *Particle) SquareSpawn() {
// 	// Fait spawn les particules sur les bords d'un carré de demi-diagonale RayonApparition.
// 	// On réutilise la variable IsExplosible de chaque particule pour savoir si elle est spawnable ou non.
// 	// Par défaut une particule est spawnable (IsExplosive = true) et donc dès que sa position permet son apparition,
// 	// elle n'est plus spawnable (IsExplosive = false). Cela permet d'éviter de la faire réapparaitre en boucle quand elle replis les conditions d'apparitions.

// 	if p.IsExplosive && (p.PositionX < config.General.SpawnX-config.General.RayonApparition ||
// 		p.PositionX > config.General.SpawnX+config.General.RayonApparition ||
// 		p.PositionY < config.General.SpawnY-config.General.RayonApparition ||
// 		p.PositionY > config.General.SpawnY+config.General.RayonApparition) {

// 		p.IsExplosive = false
// 		p.Opacity = 1
// 		if config.General.ActiveLife {
// 			p.Life = p.OriginalLife
// 		}
// 	}
// }
