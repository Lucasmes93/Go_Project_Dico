package manipulation_dictionnaire

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Dictionary struct représente votre dictionnaire
type Dictionary struct {
	entries map[string]string
	mu      sync.RWMutex
}

// NewDictionary initialise un nouveau dictionnaire
func NewDictionary() *Dictionary {
	return &Dictionary{
		entries: make(map[string]string),
	}
}

// Add permet d'ajouter une entrée au dictionnaire (Route: /add)
func (d *Dictionary) Add(w http.ResponseWriter, r *http.Request) {
	// Assurez-vous que la méthode de la requête est POST
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Exemple : Parsez le corps de la requête pour obtenir les données
	var entry map[string]string
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du corps de la requête", http.StatusBadRequest)
		return
	}

	// Ajoutez la nouvelle entrée au dictionnaire
	d.mu.Lock()
	defer d.mu.Unlock()
	d.entries[entry["mot"]] = entry["definition"]

	// Répondez avec un statut de succès
	w.WriteHeader(http.StatusCreated)
}

// Get permet de récupérer une définition par mot (Route: /get)
func (d *Dictionary) Get(w http.ResponseWriter, r *http.Request) {
	// Assurez-vous que la méthode de la requête est GET
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Exemple : Obtenez le mot à partir des paramètres de requête
	word := r.URL.Query().Get("mot")

	// Obtenez la définition du mot à partir du dictionnaire
	d.mu.RLock()
	defer d.mu.RUnlock()
	definition, exists := d.entries[word]
	if !exists {
		http.Error(w, "Mot non trouvé", http.StatusNotFound)
		return
	}

	// Répondez avec la définition du mot
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"mot": word, "definition": definition})
}

// Remove permet de supprimer une entrée par mot (Route: /remove)
func (d *Dictionary) Remove(w http.ResponseWriter, r *http.Request) {
	// Assurez-vous que la méthode de la requête est DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Exemple : Obtenez le mot à partir des paramètres de requête
	word := r.URL.Query().Get("mot")

	// Supprimez l'entrée du dictionnaire
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.entries, word)

	// Répondez avec un statut de succès
	w.WriteHeader(http.StatusOK)
}

// List permet d'afficher tous les éléments du dictionnaire (Route: /list)
func (d *Dictionary) List(w http.ResponseWriter, r *http.Request) {
	// Assurez-vous que la méthode de la requête est GET
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Obtenez la liste complète des entrées du dictionnaire
	d.mu.RLock()
	defer d.mu.RUnlock()

	// Construisez une liste de paires mot-définition
	var entries []map[string]string
	for word, definition := range d.entries {
		entry := map[string]string{"mot": word, "definition": definition}
		entries = append(entries, entry)
	}

	// Répondez avec la liste complète au format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
