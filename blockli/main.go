package main

import "fmt"
import "github.com/magichansi/blockli/blocklichain"

func main() {

	bc := blocklichain.New()
	bc.Init()
	bc.AddBlock("sdsd")
	bc.AddBlock("sdsd")
	bc.AddBlock("sdsd")
	bc.AddBlock("sdsd")
	ahash := bc.AddBlock("sdsd")
	if bc.ValidateBlock(ahash) {
		fmt.Println("geil")
	} else {
		fmt.Println("ned geil")
	}

}
