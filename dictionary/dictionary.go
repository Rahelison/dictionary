package dictionary

import (
	"fmt"
	"sort"
)

// Dictionary représente l'application de gestion de mots et de définitions.
type Dictionary struct {
    entries map[string]string
}

// Add ajoute un mot et sa définition à la map.
func (d *Dictionary) Add(word, definition string) {
    if d.entries == nil {
        d.entries = make(map[string]string)
    }
    d.entries[word] = definition
}

// Get renvoie la définition d'un mot spécifique.
func (d *Dictionary) Get(word string) (string, bool) {
    definition, found := d.entries[word]
    return definition, found
}

// Remove supprime un mot de la map.
func (d *Dictionary) Remove(word string) {
    delete(d.entries, word)
}

// List renvoie une liste triée des mots et de leurs définitions.
func (d *Dictionary) List() []string {
    var wordList []string
    for word := range d.entries {
        wordList = append(wordList, word)
    }
    sort.Strings(wordList)

    var resultList []string
    for _, word := range wordList {
        resultList = append(resultList, fmt.Sprintf("%s: %s", word, d.entries[word]))
    }
    return resultList
}
