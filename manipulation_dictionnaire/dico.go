package manipulation_dictionnaire

import (
	"encoding/json"
	"io/ioutil"
	"sort"
)

type Entry struct {
	Word       string `json:"word"`
	Definition string `json:"definition"`
}

type Dictionary struct {
	FilePath string
	Entries  []Entry
}

func NewDictionary(filePath string) *Dictionary {
	return &Dictionary{FilePath: filePath}
}

func (d *Dictionary) Add(word, definition string) error {
	entry := Entry{Word: word, Definition: definition}
	d.Entries = append(d.Entries, entry)
	return d.saveToFile()
}
func (d *Dictionary) Get(word string) (string, bool) {
	for _, entry := range d.Entries {
		if entry.Word == word {
			return entry.Definition, true
		}
	}
	return "", false
}
func (d *Dictionary) Remove(word string) error {
	var newEntries []Entry
	for _, entry := range d.Entries {
		if entry.Word != word {
			newEntries = append(newEntries, entry)
		}
	}
	d.Entries = newEntries
	return d.saveToFile()
}
func (d *Dictionary) List() ([]string, error) {
	var words []string
	for _, entry := range d.Entries {
		words = append(words, entry.Word)
	}
	sort.Strings(words)
	return words, nil
}
func (d *Dictionary) saveToFile() error {
	jsonData, err := json.MarshalIndent(d.Entries, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(d.FilePath, jsonData, 0644)
}
