package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/brittonhayes/brittonhayes/config"
	"github.com/brittonhayes/brittonhayes/pkg/templates"
)

func main() {

	var document templates.Document
	_, err := toml.Decode(config.README, &document)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(document.Render())
}
