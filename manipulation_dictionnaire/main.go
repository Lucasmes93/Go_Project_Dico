package main

import (
	"fmt"
	"sort"
)

type Dictionary map[string]string

func NewDictionary() Dictionary {
	return make(Dictionary)
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionary) Get(word string) (string, bool) {
	definition, ok := d[word]
	return definition, ok
}

func (d Dictionary) Remove(word string) {
	delete(d, word)
}

func (d Dictionary) List() []string {
	var words []string
	for word := range d {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}

func main() {
	d := NewDictionary()

	d.Add("chat", "Mammifère carnivore (félidé), sauvage ou domestique, au museau court et arrondi.")
	d.Add("chien", "Mammifère (canidé) carnivore aux multiples races, caractérisé par sa facilité à être domestiqué, par une course rapide, un excellent odorat et par son cri spécifique, l'aboiement.")
	d.Add("oiseau", "Vertébré ovipare, couvert de plumes et d'écailles cornées, à respiration pulmonaire, homéotherme, aux mâchoires sans dents revêtues d'un bec corné, et aux membres antérieurs, ou ailes, normalement adaptés au vol.")

	definition, ok := d.Get("chat")
	if ok {
		fmt.Println("Définition de 'chat':", definition)
	}

	d.Remove("oiseau")

	words := d.List()
	for _, word := range words {
		fmt.Println(word)
	}
}
