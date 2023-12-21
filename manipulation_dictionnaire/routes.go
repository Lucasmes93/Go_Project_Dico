package manipulation_dictionnaire

import "net/http"

// SetupRoutes configure les routes pour le dictionnaire
func SetupRoutes(dictionary *Dictionary) {
	http.HandleFunc("/add", dictionary.Add)
	http.HandleFunc("/get", dictionary.Get)
	http.HandleFunc("/remove", dictionary.Remove)
	http.HandleFunc("/list", dictionary.List)
	http.HandleFunc("/removeall", dictionary.RemoveAll)
	http.HandleFunc("/exporttofile", dictionary.ExportToFile)

}
