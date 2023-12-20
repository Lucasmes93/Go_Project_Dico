package manipulation_dictionnaire

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Dictionary struct {
	entries map[string]string
	mu      sync.RWMutex
}

func NewDictionary() *Dictionary {
	return &Dictionary{
		entries: make(map[string]string),
	}
}

func (d *Dictionary) handleMethodNotAllowed(w http.ResponseWriter, r *http.Request, allowedMethod string) {
	if r.Method != allowedMethod {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}
}

func (d *Dictionary) Add(w http.ResponseWriter, r *http.Request) {
	d.handleMethodNotAllowed(w, r, http.MethodPost)

	var entry map[string]string
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()
	d.entries[entry["mot"]] = entry["definition"]
	w.WriteHeader(http.StatusCreated)
}

func (d *Dictionary) Get(w http.ResponseWriter, r *http.Request) {
	d.handleMethodNotAllowed(w, r, http.MethodGet)

	word := r.URL.Query().Get("mot")

	d.mu.RLock()
	defer d.mu.RUnlock()

	definition, exists := d.entries[word]
	if !exists {
		http.Error(w, "Mot non trouvé", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mot": word, "definition": definition})
}

func (d *Dictionary) Remove(w http.ResponseWriter, r *http.Request) {
	d.handleMethodNotAllowed(w, r, http.MethodDelete)

	word := r.URL.Query().Get("mot")

	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.entries, word)
	w.WriteHeader(http.StatusOK)
}

func (d *Dictionary) List(w http.ResponseWriter, r *http.Request) {
	d.handleMethodNotAllowed(w, r, http.MethodGet)

	d.mu.RLock()
	defer d.mu.RUnlock()

	var entries []map[string]string
	for word, definition := range d.entries {
		entry := map[string]string{"mot": word, "definition": definition}
		entries = append(entries, entry)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
