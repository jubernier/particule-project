package main

// Layout définit la taille en pixels de la zone d'affichage des particules  en fonction de la taille en pixels de la fenêtre.
func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
