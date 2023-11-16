package main

import (
	"dictionary/dictionary" 
	"fmt"
)

func main() {

    dict := dictionary.Dictionary{}

    
    dict.Add("gopher", "A small burrowing animal that lives underground.")
    dict.Add("go", "A programming language created at Google.")

    
    definition, found := dict.Get("gopher")
    if found {
        fmt.Printf("Definition of 'gopher': %s\n", definition)
    } else {
        fmt.Println("Word not found.")
    }

    
    dict.Remove("gopher")


    wordList := dict.List()
    fmt.Println("List of words and definitions:")
    for _, entry := range wordList {
        fmt.Println(entry)
    }
}
