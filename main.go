package main

import (
	"Go_Project_Dico/manipulation_dictionnaire"
	"fmt"
)

func main() {
	// Crée un nouveau dictionnaire
	d := manipulation_dictionnaire.NewDictionary()

	// Ajoute des mots et leurs définitions au dictionnaire
	d.Add("chat", "Mammifère carnivore (félidé), sauvage ou domestique, au museau court et arrondi.")
	d.Add("chien", "Mammifère (canidé) carnivore aux multiples races, caractérisé par sa facilité à être domestiqué, par une course rapide, un excellent odorat et par son cri spécifique, l'aboiement.")
	d.Add("oiseau", "Vertébré ovipare, couvert de plumes et d'écailles cornées, à respiration pulmonaire, homéotherme, aux mâchoires sans dents revêtues d'un bec corné, et aux membres antérieurs, ou ailes, normalement adaptés au vol.")

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
