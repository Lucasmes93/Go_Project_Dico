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

	// Récupère et affiche la définition du mot "chat"
	definition, ok := d.Get("chat")
	if ok {
		fmt.Println("Définition de 'chat':", definition)
	}

	// Supprime le mot "oiseau" du dictionnaire
	d.Remove("oiseau")

	// Récupère la liste alphabétique des mots et les affiche
	words := d.List()
	for _, word := range words {
		fmt.Println(word)
	}
}
