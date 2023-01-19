# Particule project

# Partie 5

Pour compiler le projet, il vous faudra utiliser la commande `go build` depuis le dossier racine du projet pour obtenir un éxecutable adapté à votre système d'exploitation.

## Exemple de configuration

Pour passer d'une config.json à une autre il suffit d'utiliser les touches du clavier qui sont précisé dans la partie 5.10 ci-dessous.

1. La configuration du fichier config1.json montre le système de particules qui apprraissent formant un cercle et qui rejoigne le cercle créant une sorte de tourbillon.
2. La configuration du fichier config2.json montre le système de particules qui spawn en cercle et bougent suivant le curseur de notre souris
3. La configuration du fichier config3.json affiche nos particules qui spawn aléatoirement en utulisant de la gravité.
4.  La configuration du fichier config4.json affiche une explosion de particule .
5.  La configuration du fichier config5.json affiche des particules qui spawns autour d'un cercle en utilisant de la vitesse mais sans être attirées vers le centre .
    - Les trois config.json permetent de contrôler et changer la couleur des particules avec les touches suivantes :
            La touche B pour la couleur Bleu.
            La touche Q pour la couleur Aqua.
            La touche Y pour la couleur Jaune. 
            La touche R pour la couleur Rouge.
            La touche P pour la couleur Violet.
            La touche G pour la couleur Verte.
            * Par défaut les couleurs des particules sont générées aléatoirement.


## Variables du fichier config.json en fonction des extensions

> ### Extension 5.1 - Gravitée
> - Activable depuis la variable gravity (bool) du fichier config.json.
> - Force de gravitée modifiable avec la variable GravityCoefficient
>
> ### Extension 5.2 - Extérieur de l'écran
> - Activable depuis la variable Margin (bool) du fichier config.json.
> - Marge modifiable avec la variable MargeCoefficient.
>

> ### Extension 5.3 - Durée de vie
> - Activable depuis la variable AddLifeToParticle (bool) du fichier config.json .
>  
> - Variables :
>   - LifeRate  (Permet de régler la  vie d'une particule)
>

> ### Extension 5.4 - Variation de transparence , couleur et échelle 
> - Pour la transparence : Peut être regler depuis la variable Opacity
> - Pour l'echelle des tailles : Activable depuis la variable ChangeSize(bool) qui permet de diminuer la taille des particules avec le temps ou bien les genérer aléatoirement si on active la variable RandomSize.
> -  La variable Multicolor (bool) dans config.json permet de modifier la couleur d'une seule particule au fil du temps si elle est activée .
>

> ### Extension 5.6 - Encore plus d'optimisation de la mémoire
> - Fonctionne avec le paramêtre AddLifeToParticle.

> ### Extension 5.7 - Forme du générateur
> Activable en changeant la variable CercleShape (bool) du fichier config.json qui permet d'afficher le cercle.
>
> - Pour afficher un cercle : les particules spawn en autour 
>   Les variables à modifier sont  :
>  * CercleRadius(int) pour définir la distance de notre rayon par rapport au centre.
>  * CercleMouvement(bool) qui permet d'afficher un tourbillon.
>  * RotationSpeed (int) qui définit la vitesse angulaire.
>
> - Pour afficher une explosion : 
>  Il suffit d'utiliser la Touche E pour afficher cette forme de générateur.
>

> ### Extension 5.9 - Déplacement du générateur
>
> Activable depuis la variable MoveCursor (bool) du fichier config.json.
>
> L'activation de cette extension entraine nécessairement la désactivation du CercleMouvement 
>
> Les différentes options de cette extension sont :
>
> - Les touches directionnelles pour changer le lieu de spawn des particules, les variables à modifier sont :
> - Le curseur pour changer la position de spawn des particules, sous conditions que MoveGenerator soit activée . 
>

> ### Extension 5.10 - Modification dynamique du système de particules
> Les différentes options de cette extension :
> Les touches qui permettent de modifier la couleur : (Il faut rappuyer sur la touche de la config desirée à chaque fois avant de changer de couleur )
>   * La touche B pour la couleur Bleu.
>   * La touche Q pour la couleur Aqua.
>   * La touche Y pour la couleur Jaune. 
>   * La touche R pour la couleur Rouge.
>   * La touche P pour la couleur Violet.
>   * La touche G pour la couleur Verte.
> 
> La touche E permet de créer un effet d'explosion 
> Ainsi les touches : 
> * U : permet d'afficher des particules qui spawns autour d'un cercle sans être attirées vers le centre.
> * X : permet de visualiser l'utulisation de la gravité sur les particules.
> * Z : permet de visualiser le déplacement du générateur en fonction du curseur .
> La touche espace permet d'afficher la config de base (tourbillon)
>

## Extensions réalisées & fonctionnement

- [x] 5.1 - Gravité
  - Ajouter dans le Json un booleen pour activer la gravitée
  - Ajouter une valeur à la vitesse Y de chaque particule (avec la variable GravityCoefficient)
  - [x] Tests
  
- [x] 5.2 - Extérieur de l’écran
  - Vérifier si la position des particules est toujours situés dans l'écran (avec une marge)
  - Marge modifiable dans le fichier json (avec MargeCoefficient)
  - Lorsque la particule dépasse ces marges, la considérer comme morte (variable bool dans la structure d'une particule)
  - [x] Tests

- [x] 5.3 - Durée de vie
  - durée de vie modifiable dans le fichier json
  - 1 compteur d'appel par particule (variable dans la structure d'une particule à ajouter dans la fonctions CreateParticle)
  - Une particule qui n'a plus de durée de vie est considérée comme morte
  - L'opacitée d'une particule baisse automatiquement et en conséquences en fonction de sa durée de vie
  - [x] Tests

- [x] 5.4 - Variation de transparence , couleur et échelle :
  - modifie la taille des particules au fil du  temps 
  - modifie la couleur des particules au fil du temps ainsi que la transparence dédinies par la variable "Opacity" dans config.json
    
  - [x] Tests
  
- [x] 5.6 - Recyclage Optimisation de la mémoire 

  - [x] Tests

- [x] 5.7 - Forme du générateur
  - [x] Explosion
  - [x] Cercle 

  - [x] Tests
  
- [x] 5.9 - Déplacement du générateur
  - Vérifier la pression des touches du clavier ou la position de la souris et changer les champs de spawnX et spawnY du fichier json
  - Les tests pour cette extension ne sont pas réalisables car elle nécéssite une action de l'utilisateur sur le clavier ou la souris.

- [x] 5.10 - Modification dynamique du système de particules
  - La modification de couleurs ainsi que de config à travers des touches .
  - Les tests pour cette extension ne sont pas réalisables car elle nécéssite une action de l'utilisateur sur le clavier ou la souris.


## Tests
Résultat du `go test .\particles\` depuis le dossier racine du projet :

```
ok      project-particles/particles     0.306s
```