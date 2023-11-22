package main

import (
    "fmt"
    "sync"
    "GOLANG/dictionary"
)

func main() {
    // Création d'une instance de Dictionary avec un fichier associé
    dict := dictionary.NewDictionary("dictionary.json")

    var wg sync.WaitGroup
    wg.Add(2)

    // Utilisation de la concurrence pour effectuer simultanément des opérations d'ajout et de suppression
    go func() {
        defer wg.Done()
        dict.Add("gopher", "A small burrowing animal that lives underground.")
    }()

    go func() {
        defer wg.Done()
        dict.Remove("gopher")
    }()

    // Attendre la fin des opérations concurrentes
    wg.Wait()

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
