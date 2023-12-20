package manipulation_dictionnaire

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
