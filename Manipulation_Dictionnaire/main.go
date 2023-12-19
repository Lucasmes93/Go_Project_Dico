package main

import (
	"fmt"
	"sort"
)

type Dictionary struct {
	m map[string]string
}

func NewDictionary() Dictionary {
	return Dictionary{m: make(map[string]string)}
}

func (d *Dictionary) Add(word, definition string) {
	d.m[word] = definition
}

func (d *Dictionary) Get(word string) (string, bool) {
	definition, ok := d.m[word]
	return definition, ok
}

func (d *Dictionary) Remove(word string) {
	delete(d.m, word)
}

func (d *Dictionary) List() []string {
	var words []string
	for word := range d.m {
		words = append(words, word)
	}
	sort.Strings(words)
	return words
}

func main() {
	// Créez un nouveau dictionnaire
	d := NewDictionary()
