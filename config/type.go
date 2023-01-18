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
	ColorRandom              bool // Permet d'initialiser la couleur des particules aléatoirement
	ColorRed                 float64
	ColorGreen               float64
	ColorBlue                float64
	MaxSpeed                 float64 //Vitesse maximale d'une particule
	TypeSpeed                int     // Permet de changer la direction du déplacement des particules
	Gravity                  bool    // Activer l'extension gravitée
	GravityCoefficient       float64 // Coefficient de graitée (plus il est élevé, plus les particules sont attirés vers le bas ou vers le haut rapidement)
	AddLifeToParticle        bool
	LifeRate                 float64 // Activer l'extension Durée de vie
	Margin                   bool    // Activer l'extension extérieur de l'écran
	MargeCoefficient         float64 // Marge après laquelle les particules sont tuées si elle est dépassée
	ChangeSize               bool
	RandomSize               bool

	Opacity         float64 // Permet de définir l'intensité de l'opacité des particules
	OpacityRate     int     // Permet de définir l'intensité de l'opacité
	MoveCursor      bool    // Permet de déplacer le point de spawn des particules avec la souris
	CercleShape     bool    // Permet d'activer la forme circulaire des particules
	CercleRadius    float64 // Permet de définir la taille du rayon du cercle généré
	CercleMouvement bool    // Permet d'activer la rotation du cercle
	RotationSpeed   float64 // Permet de gérer la vitesse de rotation
	SizeX           float64 // Permet de définir la taille des particules
	SizeY           float64 // Permet de définir la taille des particules
}

var General Config
var NumberDeath int // Permet de renseigner sur les particules qui sont mortes
