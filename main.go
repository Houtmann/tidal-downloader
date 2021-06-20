package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

var tidal *Tidal

func main() {
	var err error
	tidal, err = NewTidal()
	if err != nil {
		return 
	}

	err = tidal.Configure()
	if err != nil {
		return
	}

	p := prompt.New(
		executor,
		completer,
	)
	p.Run()

}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "downloads", Description: "downloads music, or playlist"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}


func executor(t string) {
	if t == "downloads" {
		fmt.Println(t)
	}
	return
}