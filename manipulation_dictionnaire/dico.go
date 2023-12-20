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
