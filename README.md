
# Projet Dictionnaire en Go

Ce projet implémente un programme en Go permettant la manipulation d'un dictionnaire stocké dans un fichier JSON. Le dictionnaire peut être utilisé pour ajouter des mots avec leurs définitions, récupérer les définitions des mots, supprimer des mots, et lister les mots restants par ordre alphabétique.

## Structure du Projet

- **main.go**: Le fichier principal qui contient la fonction `main` du programme.

- **manipulation_dictionnaire**: Un package personnalisé contenant les fonctionnalités de manipulation du dictionnaire.

- **dictionary.json**: Le fichier JSON dans lequel le dictionnaire est stocké.

## Fonctionnalités

### Ajout d'un Mot et d'une Définition

Utilisez la méthode `Add` du dictionnaire pour ajouter un mot avec sa définition.

```go
handleError(d.Add("chat", "Mammifère carnivore (félidé), sauvage ou domestique, au museau court et arrondi."), "Erreur lors de l'ajout")
```

### Récupération de la Définition d'un Mot

Utilisez la méthode `Get` pour récupérer la définition d'un mot.

```go
definition, ok := d.Get("chat")
if ok {
    fmt.Println("Définition de 'chat':", definition)
} else {
    fmt.Println("Mot non trouvé.")
}
```

### Suppression d'un Mot

Utilisez la méthode `Remove` pour supprimer un mot du dictionnaire.

```go
handleError(d.Remove("chien"), "Erreur lors de la suppression")
```

### Liste Alphabétique des Mots Restants

Utilisez la méthode `List` pour obtenir une liste alphabétique des mots restants.

```go
words, err := d.List()
handleError(err, "Erreur lors de la récupération de la liste")
fmt.Println("Mots restants:", words)
```

## Installation et Utilisation

1. **Prérequis**: Assurez-vous que Go est installé sur votre système. Si ce n'est pas le cas, suivez les instructions d'installation sur [golang.org](https://golang.org/doc/install).

2. **Clonage du Dépôt**:
   ```bash
   git clone https://github.com/Lucasmes93/Go_Project_Dico.git
   ```

3. **Navigation dans le Répertoire du Projet**:
   ```bash
   cd Go_Project_Dico
   ```

4. **Exécution du Programme**:
   ```bash
   go run main.go
   ```

## Contribuer

N'hésitez pas à contribuer en ouvrant des issues pour signaler des problèmes ou en proposant des pull requests pour des améliorations. Toute contribution est la bienvenue !

## Auteur

- [MESSIA DOLIVEUX Lucas]
