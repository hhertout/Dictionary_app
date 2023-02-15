package main

import (
	"flag"
	"fmt"
	"training/dictionary_project/cli"
	"training/dictionary_project/dictionary"
)

func main() {
	action := flag.String("action", "list", "Action to perform on the dictionary")

	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()

	flag.Parse()
	switch *action {
	case "list":
		cli.ActionList(d)
	case "add":
		cli.ActionAdd(d, flag.Args())
	case "define":
		cli.ActionDefine(d, flag.Args())
	case "remove":
		cli.ActionRemove(d, flag.Args())
	default:
		fmt.Printf("Unknow action: %v\n", action)
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
	}
}
