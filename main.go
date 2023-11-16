package main

import (
	"GOLANG/dictionary"
	"fmt"
)

func main() {
    // Création d'une instance de Dictionary avec un fichier associé
    dict := dictionary.NewDictionary("dictionary.json")

    // Utilisation de la méthode Add pour ajouter des mots et des définitions
    if err := dict.Add("gopher", "A small burrowing animal that lives underground."); err != nil {
        fmt.Println("Error adding entry:", err)
    }

    if err := dict.Add("go", "A programming language created at Google."); err != nil {
        fmt.Println("Error adding entry:", err)
    }

    // Utilisation de la méthode Get pour afficher la définition d'un mot spécifique
    definition, found := dict.Get("gopher")
    if found {
        fmt.Printf("Definition of 'gopher': %s\n", definition)
    } else {
        fmt.Println("Word not found.")
    }

    // Utilisation de la méthode Remove pour supprimer un mot du fichier
    if err := dict.Remove("gopher"); err != nil {
        fmt.Println("Error removing entry:", err)
    }

    // Utilisation de la méthode List pour obtenir la liste triée des mots et de leurs définitions
    entries, err := dict.List()
    if err != nil {
        fmt.Println("Error listing entries:", err)
    } else {
        fmt.Println("List of words and definitions:")
        for word, definition := range entries {
            fmt.Printf("%s: %s\n", word, definition)
        }
    }
}
