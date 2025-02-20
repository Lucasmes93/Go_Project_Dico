package manipulation_dictionnaire

import (
	"log"
	"net/http"
)

// SetupRoutes configure toutes les routes avec leurs méthodes HTTP correctes
func SetupRoutes(mux *http.ServeMux, dictionary *Dictionary) {
	mux.Handle("/add", methodHandler(dictionary.Add, http.MethodPost))
	mux.Handle("/update", methodHandler(dictionary.Update, http.MethodPut))
	mux.Handle("/remove", methodHandler(dictionary.Remove, http.MethodDelete))
	mux.Handle("/list", methodHandler(dictionary.List, http.MethodGet))
	mux.Handle("/search", methodHandler(dictionary.Search, http.MethodGet))
	mux.Handle("/count", methodHandler(dictionary.Count, http.MethodGet))
	mux.Handle("/health", methodHandler(HealthCheck, http.MethodGet))
}

// Middleware pour forcer une seule méthode HTTP par route
func methodHandler(handlerFunc http.HandlerFunc, allowedMethod string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
			return
		}
		log.Printf("📡 [%s] %s", r.Method, r.URL.Path)
		handlerFunc(w, r)
	})
}
