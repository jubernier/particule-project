# Particule project

# Partie 5

Pour compiler le projet, il vous faudra utiliser la commande `go build` depuis le dossier racine du projet pour obtenir un éxecutable adapté à votre système d'exploitation.

## Exemple de configuration

Pour passer d'une config.json à une autre il suffit d'utiliser les fléches du clavier .

1. La configuration du fichier config1.json montre le système de particules avec un cercle tourne en créant une sorte de tourbillon.
2. La configuration du fichier config2.json montre le système de particules qui spawn en cercle et bougent suivant le curseur de notre souris
3. La configuration du fichier config3.json affiche nos particules qui spawn aléatoirement .
    - Les trois config.json permetent de contrôler et changer la couleur des particules avec les touches suivantes :
            La touche B pour la couleur Bleu .
            La touche Q pour la couleur Aqua.
            La touche Y pour la couleur Jaune. 
            La touche R pour la couleur Rouge.
            La touche P pour la couleur Violet.
            La touche G pour la couleur Verte.
            * Par défaut les couleurs des particules sont générées aléatoirement.


## Variables du fichier config.json en fonction des extensions

> ### Extension 5.1 - Gravitée
> - Activable depuis la variable gravity (bool) du fichier config.json .
> - Force de gravitée modifiable avec la variable GravityCoefficient
>
> Les extensions fonctionnant avec celle-ci :
> - 5.2 Extérieur de l'écran
> - 5.3 Durée de vie
> - 5.4 Variations de couleur, d’échelle, de rotation, de transparence
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur (peu importe le type de spawn)
> - 5.9 Déplacement du générateur
> - 5.10 Modification dynamique du systeme de particules

> ### Extension 5.2 - Extérieur de l'écran
> - Activable depuis la variable Margin (bool) du fichier config.json .
> - Marge modifiable avec la variable MargeCoefficient
>
> Les extensions fonctionnant avec celle-ci :
> - 5.1 Gravitée
> - 5.3 Durée de vie
> - 5.4 Variations de couleur, d’échelle, de rotation, de transparence
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur (peu importe le type de spawn)
> - 5.9 Déplacement du générateur
> - 5.10 Modification dynamique du systeme de particules

> ### Extension 5.3 - Durée de vie
> - Activable depuis la variable AddLifeToParticle (bool) du fichier config.json .
>  
> - Variables :
>   - LifeRate  (Permet de régler la  vie d'une particule)
>
> Les extensions fonctionnant avec celle-ci :
> - 5.1 Gravitée
> - 5.2 Extérieur de l'écran
> - 5.4 Variations de couleur, d’échelle, de rotation, de transparence
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur :
>   - Peu importe le type de spawn
> - 5.9 Déplacement du générateur
> - 5.10 Modification dynamique du systeme de particules

> ### Extension 5.4 - Variation de transparence , couleur et échelle 
> - Pour la transparence : Peut être regler depuis la variable Opacity
> - Pour l'echelle des tailles : Activable depuis la variable ChangeSize(bool) qui permet de diminuer la taille des particules avec le temps ou bien les genérer aléatoirement si on active la variable RandomSize.
> -  La variable Multicolor (bool) dans config.json permet de modifier la couleur d'une seule particule au fil du temps si elle est activée .
>
> Les extensions fonctionnant avec celle-ci :
> - 5.1 Gravitée
> - 5.2 Extérieur de l'écran
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur :  
>   - Peu importe le type de spawn
> - 5.9 Déplacement du générateur
> - 5.10 Modification dynamique du systeme de particules
> ### Extension 5.5 - Optimisation de la mémoire
> Fonctionne automatiquement avec toutes les extensions

> ### Extension 5.7 - Forme du générateur
> Activable en changeant la variable CercleShape (bool) du fichier config.json qui permet d'afficher notre cercle .
>
> - Pour afficher un cercle : les particules spawn en autour 
>   Les variables à modifier sont  :
>  * CercleRadius(int) pour définir la taille de notre rayon .
>  * CercleMouvement(bool) qui permet d'afficher notre tourbillon .
>  * RotationSpeed (int) qui définit la vitesse de la rotation .
>
> - Pour afficher une explosion : 
>  Il suffit d'utiliser la Touche E pour afficher cette forme de générateur .
>
> Les extensions fonctionnant avec celle-ci sont les suivantes :
> - 5.1 Gravitée
> - 5.2 Extérieur de l'écran
> - 5.3 Durée de vie
> - 5.4 Variations de couleur, d’échelle, de rotation, de transparence
> - 5.5 Optimisation de la mémoire
> - 5.9 Sauf le tourbillon 
> - 5.10 Modification dynamique du systeme de particules

> ### Extension 5.9 - Déplacement du générateur
>
> Activable depuis la variable MoveCursor (bool) du fichier config.json .
>
> L'activation de cette extension entraine nécessairement la désactivation du CercleMouvement (je ne suis pas sure)
>
> Les différentes options de cette extension sont :
>
> - Les touches directionnelles pour changer le lieu de spawn des particules, les variables à modifier sont :
> - Le curseur pour changer la position de spawn des particules, sous conditions que MoveGenerator soit activée . 
>
> Les extensions fonctionnant avec celle-ci sont les suivantes :
>
> - 5.1 Gravitée
> - 5.2 Extérieur de l'écran
> - 5.3 Durée de vie
> - 5.4 Variations de couleur, d’échelle, de rotation, de transparence
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur (sauf tourbillon)
> - 5.10 Modification dynamique du systeme de particules

> ### Extension 5.10 - Modification dynamique du système de particules
> Les différentes options de cette extension :
> Les touches qui permettent de modifier la couleur : 
>   * La touche B pour la couleur Bleu .
>   * La touche Q pour la couleur Aqua.
>   * La touche Y pour la couleur Jaune. 
>   * La touche R pour la couleur Rouge.
>   * La touche P pour la couleur Violet.
>   * La touche G pour la couleur Verte.
> 
> La touche E permet de créer un effet d'explosion 
> 
> Les extensions fonctionnant avec celle-ci sont :
> - 5.1 Gravitée
> - 5.2 Extérieur de l'écran
> - 5.3 Durée de vie
> - 5.4 Variations de couleur, d’échelle, de rotation, de transparence
> - 5.5 Optimisation de la mémoire
> - 5.7 Forme du générateur
> - 5.9 Déplacement du générateur

## Extensions réalisées & fonctionnement

- [x] 5.1 - Gravité
  - Ajouter dans le Json un booleen pour activer la gravitée
  - Ajouter une valeur à la vitesse Y de chaque particule (avec la variable GravityCoefficient)
  - [x] Tests
  
- [x] 5.2 - Extérieur de l’écran
  - Vérifier si la position des particules est toujours situés dans l'écran (avec une marge)
  - Marge modifiable dans le fichier json (avec MargeCoefficient)
  - Lorsque la particule dépasse ces marges, la considérer comme morte (variable bool dans la structure d'une particule)
  - Ne plus l'afficher (dans le fichier draw.go)
  - [x] Tests

- [x] 5.3 - Durée de vie
  - durée de vie modifiable dans le fichier json
  - 1 compteur d'appel par particule (variable dans la structure d'une particule à ajouter dans la fonctions CreateParticle)
  - Une particule qui n'a plus de durée de vie est considérée comme morte
  - L'opacitée d'une particule baisse automatiquement et en conséquences en fonction de sa durée de vie
  - [x] Tests

- [x] 5.5 - Optimisation de la mémoire
  - La méthode add() prend en compte :
    - Si il y a des particules mortes, on les réutilises
    - Sinon on en ajoute au tableau
  - Variable LastInLife dans laquelle est stocké l'index de la dernière particule vivante
  - [x] Tests
  
- [x] 5.7 - Forme du générateur
  - [x] Explosion
    - Définir si une particule est "explosive" ou non
    - D'autres particules non explosives serront crées à partir de la position de la particule explosive
  - [x] Cercle 

  - [x] Tests
  
- [x] 5.9 - Déplacement du générateur
  - Vérifier la pression des touches du clavier ou la position de la souris et changer les champs de spawnX et spawnY du fichier json
  - Les tests pour cette extension ne sont pas réalisables car elle nécéssite une action de l'utilisateur sur le clavier ou la souris.

- [x] 5.10 - Modification dynamique du système de particules
  - Lorsque la touche G est pressée, si le menu n'est pas affiché, il est affiché sinon il est caché (et ne permet plus la modification de valeurs)
  - La navigation du menu se fait avec les touches + et -. Retour au premier paramètre du menu si on est au dernier, et au dernier si on est au premier.
  - La modification des valeurs se fait par la molette, un paramètre permet de savoir si une valeur peut être négative et se bloque à 0 si ce n'est pas le cas.
  - Les tests pour cette extension ne sont pas réalisables car elle nécéssite une action de l'utilisateur sur le clavier ou la souris.


## Tests
Résultat du `go test ./tests` depuis le dossier racine du projet :

```
ok      project-particles/tests 0.017s
```