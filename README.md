# Dictionary Go

Un simple programme Go représentant un dictionnaire avec des opérations de base telles que l'ajout de mots, la récupération de définitions, la suppression de mots et la liste triée des mots.

#Le projet est organisé en plusieurs composants :

- **manipulation_dictionnaire**: Ce répertoire contient le package qui gère la manipulation du dictionnaire.
  - **dico.go**: Ce fichier définit le type Dictionary et ses méthodes pour ajouter, récupérer, supprimer et lister des mots avec leurs définitions.

- **main**: Ce répertoire contient l'application principale.
  - **main.go**: Ce fichier utilise le package manipulation_dictionnaire pour interagir avec le dictionnaire, en ajoutant des mots, récupérant des définitions, supprimant des mots et listant le dictionnaire.

## Utilisation

1. **Créez un nouveau dictionnaire :**
   ```go
   d := NewDictionary()
   ```

2. **Ajoutez des mots et des définitions :**
   ```go
   d.Add("chat", "Mammifère carnivore (félidé), sauvage ou domestique, au museau court et arrondi.")
   d.Add("chien", "Mammifère (canidé) carnivore aux multiples races, caractérisé par sa facilité à être domestiqué, par une course rapide, un excellent odorat et par son cri spécifique, l'aboiement.")
   d.Add("oiseau", "Vertébré ovipare, couvert de plumes et d'écailles cornées, à respiration pulmonaire, homéotherme, aux mâchoires sans dents revêtues d'un bec corné, et aux membres antérieurs, ou ailes, normalement adaptés au vol.")
   ```

3. **Obtenez la définition d'un mot :**
   ```go
   definition, ok := d.Get("chat")
   if ok {
       fmt.Println("Définition de 'chat':", definition)
   }
   ```

4. **Supprimez un mot :**
   ```go
   d.Remove("oiseau")
   ```

5. **Obtenez la liste triée des mots :**
   ```go
   words := d.List()
   for _, word := range words {
       fmt.Println(word)
   }
   ```

## Structure du Code

- `main.go`: Contient le code principal pour utiliser le dictionnaire.
- `dictionary.go`: Définit le type `Dictionary` et ses méthodes pour les opérations de base.
