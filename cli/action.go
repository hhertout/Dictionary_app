package cli

import (
	"fmt"
	"training/dictionary_project/dictionary"
)

func ActionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Println("Dictionary content")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func ActionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleErr(err)
	fmt.Printf("'%v' added to the dictionnary \n", word)
}
func ActionDefine(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	handleErr(err)
	fmt.Println(entry)
}
func ActionRemove(d *dictionary.Dictionary, args []string) {
	word := args[0]
	err := d.Remove(word)
	handleErr(err)
	fmt.Printf("'%v' was remove from the dictionnary", word)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
	}
}
