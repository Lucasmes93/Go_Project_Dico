package main

import (
	"Go_Project_Dico/manipulation_dictionnaire"
	"fmt"
)

func main() {
	filePath := "dictionary.json"
	d := manipulation_dictionnaire.NewDictionary(filePath)

	// Ajout de mots et de définitions
	handleError(d.Add("chat", "Mammifère carnivore (félidé), sauvage ou domestique, au museau court et arrondi."), "Erreur lors de l'ajout")

	// Récupération de la définition du mot "chat"
	definition, ok := d.Get("chat")
	if ok {
		fmt.Println("Définition de 'chat':", definition)
	} else {
		fmt.Println("Mot non trouvé.")
	}

	// Suppression du mot "chien"
	handleError(d.Remove("chien"), "Erreur lors de la suppression")

	// Liste alphabétique des mots restants
	words, err := d.List()
	handleError(err, "Erreur lors de la récupération de la liste")
	fmt.Println("Mots restants:", words)
}

func handleError(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %v\n", message, err)
	}
}
