package main

import (
	"blockchain/blockstore"
	"blockchain/block"
	"fmt"

)

func main() {

	a := blockstore.New("./hier.db")
	b := block.New("sdsd","AAAA")
	a.Save("sds",b)
	fmt.Println(a.Retrieve("sds").GetData())

}
