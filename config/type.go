package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
// Vous pouvez ajouter des champs et ils seront automatiquement lus dans le
// fichier de config. Vous devrez le faire plusieurs fois durant le projet.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	RandomSpawn              bool
	SpawnX, SpawnY           int
	SpawnRate                float64
	Multicolor               bool
	ColorRandom              bool
	ColorRed                 float64
	ColorGreen               float64
	ColorBlue                float64
	MaxSpeed                 float64
	TypeSpeed                int
	Gravity                  bool
	GravityCoefficient       float64
	AddLifeToParticle        bool
	LifeRate                 float64
	Margin                   bool
	MargeCoefficient         float64
	ChangeSize               bool
	RandomSize               bool

	OpacityRate     int
	MoveCursor      bool
	CercleShape     bool
	CercleRadius    float64
	CercleMouvement bool
	RotationSpeed   float64
	MouseMouvement  bool
	Opacity         float64
	SizeX           float64
	SizeY           float64
}

var General Config
var NumberDeath int
