# Projet Dictionnaire en Go

Ce projet implémente un serveur HTTP en Go pour la manipulation d'un dictionnaire. Le dictionnaire est géré à l'aide de trois fichiers principaux : `main.go`, `manipulation_dictionnaire/dictionary.go`, et `manipulation_dictionnaire/routes.go`. Le serveur expose des endpoints pour ajouter des mots, récupérer des définitions, supprimer des mots, et lister tous les mots du dictionnaire.

## Structure du Projet

- **main.go**: Fichier principal du projet contenant la fonction `main`. Il initialise le dictionnaire, configure les routes, et démarre le serveur HTTP.

- **manipulation_dictionnaire/dictionary.go**: Contient la définition du type `Dictionary` et les fonctions associées pour ajouter, récupérer, supprimer et lister des entrées dans le dictionnaire.

- **manipulation_dictionnaire/routes.go**: Configure les routes du serveur HTTP pour les opérations du dictionnaire, telles que l'ajout, la récupération, la suppression et la liste.

## Fonctionnalités

### Ajout d'un Mot et d'une Définition

Utilisez la route `/add` pour ajouter un mot avec sa définition. Exemple en utilisant curl :

```bash
curl -X POST -d '{"mot": "example", "definition": "This is an example."}' http://localhost:8080/add
```

### Récupération de la Définition d'un Mot

Utilisez la route `/get` pour récupérer la définition d'un mot. Exemple en utilisant curl :

```bash
curl http://localhost:8080/get?mot=example
```

### Suppression d'un Mot

Utilisez la route `/remove` pour supprimer un mot du dictionnaire. Exemple en utilisant curl :

```bash
curl -X DELETE http://localhost:8080/remove?mot=example
```

### Liste de Tous les Mots

Utilisez la route `/list` pour obtenir une liste de tous les mots du dictionnaire. Exemple en utilisant curl :

```bash
curl http://localhost:8080/list
```

## Installation et Utilisation

1. **Prérequis**: Assurez-vous que Go est installé sur votre système. Si ce n'est pas le cas, suivez les instructions d'installation sur [golang.org](https://golang.org/doc/install).

2. **Clonage du Dépôt**:
   ```bash
   git clone https://github.com/Lucasmes93/Go_Project_Dico.git
   cd Go_Project_Dico
    ```

3. **Exécution du Programme**:
   ```bash
   go run main.go
   ```

Le serveur sera en cours d'exécution sur le port 8080. Vous pouvez maintenant utiliser les routes mentionnées ci-dessus pour manipuler le dictionnaire.

## Contributions

Les contributions sont les bienvenues ! Si vous souhaitez améliorer ce projet, n'hésitez pas à ouvrir une pull request.

## Auteurs

- [MESSIA DOLIVEUX Lucas - DINSPEL Ludwig.Jr]
