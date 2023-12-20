package main

import (
	"Go_Project_Dico/manipulation_dictionnaire"
	"fmt"
	"net/http"
)

const port = 8080

func main() {
	dictionary := manipulation_dictionnaire.NewDictionary()
	manipulation_dictionnaire.SetupRoutes(dictionary)

	fmt.Printf("Serveur en cours d'exécution sur le port %d...\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Erreur lors du démarrage du serveur: %s\n", err)
	}
}
