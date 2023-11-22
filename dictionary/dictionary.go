package dictionary

import (
	"encoding/json"
	"os"
	"sync"
)

// Dictionary représente l'application de gestion de mots et de définitions.
type Dictionary struct {
	fileName string
	mu       sync.Mutex
	entries  map[string]string
	addCh    chan entryOperation
	removeCh chan entryOperation
}

// entryOperation représente une opération d'ajout ou de suppression.
type entryOperation struct {
	word       string
	definition string
}

// NewDictionary crée une nouvelle instance de Dictionary avec un fichier associé.
func NewDictionary(fileName string) *Dictionary {
	dict := &Dictionary{
		fileName: fileName,
		entries:  make(map[string]string),
		addCh:    make(chan entryOperation),
		removeCh: make(chan entryOperation),
	}

	// Lancer le worker pour traiter les opérations d'ajout et de suppression
	go dict.entryWorker()

	return dict
}

// entryWorker traite les opérations d'ajout et de suppression de manière concurrente.
func (d *Dictionary) entryWorker() {
	for {
		select {
		case entry := <-d.addCh:
			d.add(entry.word, entry.definition)
		case entry := <-d.removeCh:
			d.remove(entry.word)
		}
	}
}

// Add ajoute un mot et sa définition au fichier de manière concurrente.
func (d *Dictionary) Add(word, definition string) {
	d.addCh <- entryOperation{word, definition}
}

// Remove supprime un mot du fichier de manière concurrente.
func (d *Dictionary) Remove(word string) {
	d.removeCh <- entryOperation{word, ""}
}

// add ajoute un mot et sa définition à la map et écrit dans le fichier.
func (d *Dictionary) add(word, definition string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.entries[word] = definition
	d.writeToFile()
}

// remove supprime un mot de la map et écrit dans le fichier.
func (d *Dictionary) remove(word string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	delete(d.entries, word)
	d.writeToFile()
}

// List renvoie une liste triée des mots et de leurs définitions.
func (d *Dictionary) List() (map[string]string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	return d.entries, nil
}

// readFromFile lit le contenu du fichier et le convertit en map.
func (d *Dictionary) readFromFile() error {
	file, err := os.Open(d.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&d.entries)
	if err != nil {
		return err
	}

	return nil
}

// writeToFile prend la map actuelle et écrit son contenu dans le fichier.
func (d *Dictionary) writeToFile() error {
	file, err := os.Create(d.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(d.entries)
	if err != nil {
		return err
	}

	return nil
}
