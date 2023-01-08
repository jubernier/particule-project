# Particule project

Pour compiler le projet, il vous faudra utiliser la commande `go build` depuis le dossier racine du projet pour obtenir un éxecutable adapté à votre système d'exploitation.

## Partie 4

### Description
    _Tous les paramêtres expliqués ci-dessous se trouve dans le fichier config.json, ils peuvent être modifié en fonction de l'effet voulut_

- [X] NewSystem():
    - La fonction ParticuleCr() permet :
        - de créer une particule
    - Les particules apparaissent en fonction du nombre attribué au parametre InitNumParticles
    - Si le paramêtre RandomSpawn est activé, les particules apparaissent de manière aléatoire dans la fenêtre sinon au coordonnée précisé par les paramètres SpawnX et SpawnY

- [X] Update(): vitesse des particules
    - Ajoute la vitesse au position des particules
    - Ici, la fonction ParticuleCr() permet :
        - de définir la vitesse SpeedX et la vitesse SpeedY de la particule créer, les vitesses sont aléatoires et en fonction de la valeur du paramêtre MaxSpeed
        - si le paramêtre TypeSpeed est a 0 les valeurs des vitesses peuvent être négatives
        - si le paramêtre TypeSpeed est a 1 les valeurs des vitesses ne seront que négatives
        - si le paramêtre TypeSpeed est a une autre valeur, les valeurs des vitesses ne seront que négatives 

- [X] Update(): génération de particules
    - Fait apparaitre un nombre de particules définis par le paramêtre SpawnRate au cours du temps

## Test

Expliquer test

Résultat du `go test ./tests` effectué dans le dossier projet:
```
ok      project-particles/tests

```

## Contribution
Basma Malki et Bernier Justine INFO 1 Groupe 1