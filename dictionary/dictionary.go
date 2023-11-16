package dictionary

import (
	"encoding/json"
	"os"
)

// Dictionary represents the application for managing words and definitions.
type Dictionary struct {
	fileName string
}

// NewDictionary creates a new instance of Dictionary with an associated file.
func NewDictionary(fileName string) *Dictionary {
	return &Dictionary{fileName: fileName}
}

// Add adds a word and its definition to the file.
func (d *Dictionary) Add(word, definition string) error {
	entries, err := d.readFromFile()
	if err != nil {
		return err
	}

	entries[word] = definition

	return d.writeToFile(entries)
}

// Get returns the definition of a specific word.
func (d *Dictionary) Get(word string) (string, bool) {
	entries, err := d.readFromFile()
	if err != nil {
		return "", false
	}

	definition, found := entries[word]
	return definition, found
}

// Remove removes a word from the file.
func (d *Dictionary) Remove(word string) error {
	entries, err := d.readFromFile()
	if err != nil {
		return err
	}

	delete(entries, word)

	return d.writeToFile(entries)
}

// List returns a sorted list of words and their definitions.
func (d *Dictionary) List() (map[string]string, error) {
	return d.readFromFile()
}

// readFromFile reads the content from the file and converts it into a map.
func (d *Dictionary) readFromFile() (map[string]string, error) {
	file, err := os.Open(d.fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var entries map[string]string
	err = decoder.Decode(&entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

// writeToFile takes a map and writes its content to the file.
func (d *Dictionary) writeToFile(entries map[string]string) error {
	file, err := os.Create(d.fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(entries)
	if err != nil {
		return err
	}

	return nil
}
