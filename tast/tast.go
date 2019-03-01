package main

import "fmt"
import "blockli/blockli"

func main() {

	bc := blockli.New()
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
