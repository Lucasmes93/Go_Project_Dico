package main

import (
	"Go_Project_Dico/manipulation_dictionnaire"
	"fmt"
	"net/http"
)

func main() {
	// Initialisez votre dictionnaire
	dictionary := manipulation_dictionnaire.NewDictionary()

	// Configurez les routes
	manipulation_dictionnaire.SetupRoutes(dictionary)

	// Démarrer le serveur sur le port 8080
	port := 8080
	fmt.Printf("Serveur en cours d'exécution sur le port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
