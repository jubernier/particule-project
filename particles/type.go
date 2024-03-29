package particles

import "container/list"

// System définit un système de particules.
type System struct {
	Content *list.List
}

// Particle définit une particule.
type Particle struct {
	PositionX, PositionY            float64
	Rotation                        float64
	ScaleX, ScaleY                  float64
	ColorRed, ColorGreen, ColorBlue float64
	Opacity                         float64
	SpeedX, SpeedY                  float64
	LifeRate                        float64
	ExDistance                      float64 //Variable permetant de définir l'ancienne distance (rayon) avant chaque rotation
}
