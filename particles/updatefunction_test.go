package particles

import (
	"container/list"
	"project-particles/config"

	"testing"
)

// Test qui permet de vérifier si la position de la particule change en un appel de la fonction Update().
func TestUpdate(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.TypeSpeed = 1
	config.General.MaxSpeed = 10
	var l *list.List = list.New()
	l.PushFront(CreatParticle())
	var SystemParticle *System = &System{Content: l}
	var particule *Particle = l.Front().Value.(*Particle)
	var posX float64 = particule.PositionX
	var posY float64 = particule.PositionY

	SystemParticle.Update()

	if posX == SystemParticle.Content.Front().Value.(*Particle).PositionX || posY == SystemParticle.Content.Front().Value.(*Particle).PositionY {
		t.Error("Les particules ne se déplace pas et pourtant elle devrait.")
	}
}

// Test qui permet de vérifier si il y a bien création de particules comme le définis le paramêtre SpawnRate.
func TestRandomSpawnUpdate(t *testing.T) {
	config.General.SpawnRate = 1
	var l *list.List = list.New()
	l.PushFront(CreatParticle())

	var systema *System = &System{Content: l}
	var systemb *System = &System{Content: list.New()}

	systema.Update()
	//systema est le system initiale comparer au systemb qui est modifier par la fonction Update()
	if systema.Content.Len() == systemb.Content.Len() {
		t.Error("Les particules ne sont pas crées au cours du temps alors que cela devraient être possible.")
	}
}

// Test qui vérifi que la vitesse n'est bien pas la même si speedy est avec gravité et sans a la fin
func TestUpdateGravity(t *testing.T) {
	config.General.MaxSpeed = 0
	config.General.Gravity = true
	config.General.GravityCoefficient = 0.6
	var l *list.List = list.New()
	l.PushFront(CreatParticle())
	var SystemEx *System = &System{Content: l}
	var particule *Particle = l.Front().Value.(*Particle)
	SystemEx.Update()
	if particule.SpeedY == 0 {
		t.Error("Les particules ne devraient subir la gravité ce n'est pas le cas ici.")
	}
}

// Test si la particule vieillit bien a chaque appel de update.
func TestUpdateLifeRate(t *testing.T) {
	config.General.AddLifeToParticle = true
	config.General.LifeRate = 10
	var l *list.List = list.New()
	l.PushFront(CreatParticle())
	var SystemEx *System = &System{Content: l}
	var particule *Particle = l.Front().Value.(*Particle)
	SystemEx.Update()
	if particule.LifeRate == 0 {
		t.Error("La vie de la particule devrait augmenter.")
	}
}

// Test si les particules meurt bien a la fin de leur vie.
func TestUpdateNumberDeath(t *testing.T) {
	config.General.AddLifeToParticle = true
	config.General.LifeRate = 1
	var l *list.List = list.New()
	l.PushFront(CreatParticle())
	var SystemEx *System = &System{Content: l}
	SystemEx.Update()
	if config.NumberDeath == 0 {
		t.Error("La particule devrait mourir a la fin de sa vie ce qui n'est pas le cas ici.s")
	}
}

// Test si une particule quand sa position dépasse les marges est bien morte.
func TestUpdateMargin(t *testing.T) {
	config.General.Margin = true
	config.General.MargeCoefficient = 0
	config.General.RandomSpawn = false
	config.General.SpawnX = config.General.WindowSizeX + 100
	config.General.AddLifeToParticle = false

	var l *list.List = list.New()
	l.PushFront(CreatParticle())
	var SystemBis *System = &System{Content: l}
	SystemBis.Update()

	if config.NumberDeath == 0 {
		t.Error("La particule devrait mourir car elle apparait au delà de la marge extérieur et pourtant ce n'est pas le cas.")
	}
}
