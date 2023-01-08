package main

import "project-particles/particles"

// game est un type qui implémente l'interface Game de la bibliothèque Ebiten
// car il dispose d'une méthode Update, d'une méthode Draw et d'une méthode
// Layout.
type game struct {
	system particles.System
}
