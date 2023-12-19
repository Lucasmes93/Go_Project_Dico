# Dictionnaire Simple en Go

Ce programme en Go propose une implémentation simple d'un dictionnaire avec des fonctionnalités de base pour gérer des mots et leurs définitions. Le dictionnaire est implémenté à l'aide d'une structure de données de type map.

## Fonctionnalités

1. **Créer un Nouveau Dictionnaire :**
    ```go
    d := NewDictionary()
    ```

2. **Ajouter des Mots et des Définitions :**
    ```go
    d.Add("chat", "Un animal domestique à quatre pattes.")
    d.Add("chien", "Un autre animal domestique à quatre pattes.")
    d.Add("oiseau", "Un animal volant.")
    ```

3. **Obtenir la Définition d'un Mot :**
    ```go
    definition, ok := d.Get("chat")
    if ok {
        fmt.Println("Définition du mot 'chat' :", definition)
    }
    ```

4. **Supprimer un Mot :**
    ```go
    d.Remove("oiseau")
    ```

5. **Lister les Mots par Ordre Alphabétique :**
    ```go
    mots := d.List()
    for _, mot := range mots {
        fmt.Println(mot)
    }
    ```

## Comment ça fonctionne

- La structure `Dictionary` utilise une map pour stocker les mots et leurs définitions.
- Les méthodes `Add`, `Get`, et `Remove` permettent respectivement d'ajouter un mot, d'obtenir la définition d'un mot, et de supprimer un mot.
- La méthode `List` renvoie une liste triée par ordre alphabétique des mots présents dans le dictionnaire.

## Exemple d'Utilisation

```go
// Créez un nouveau dictionnaire
d := NewDictionary()

// Ajoutez des mots et des définitions
	d.Add("chat", "Mammifère carnivore (félidé), sauvage ou domestique, au museau court et arrondi.")
	d.Add("chien", "Mammifère (canidé) carnivore aux multiples races, caractérisé par sa facilité à être domestiqué, par une course rapide, un excellent odorat et par son cri spécifique, l'aboiement.")
	d.Add("oiseau", "Vertébré ovipare, couvert de plumes et d'écailles cornées, à respiration pulmonaire, homéotherme, aux mâchoires sans dents revêtues d'un bec corné, et aux membres antérieurs, ou ailes, normalement adaptés au vol.")

// Obtenez la définition d'un mot
definition, ok := d.Get("chat")
if ok {
    fmt.Println("Définition du mot 'chat' :", definition)
}

// Supprimez un mot
d.Remove("oiseau")

// Obtenez la liste triée des mots
mots := d.List()
for _, mot := range mots {
    fmt.Println(mot)
}
