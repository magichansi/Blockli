package blockli

import (
	"encoding/json"
	"log"
)

//Blockli sds
type Blockli struct {
	prevgenesisBlockHash string
	blockCache           *BlockCache
	currentBlockHash     string
}

//NewBlockli asdad
func NewBlockli() *Blockli {

	a := new(Blockli)
	a.blockCache = NewBlockCache("C:/dev/Go")
	return a
}

//Init sdsd
func (bc *Blockli) Init() string {
	bc.prevgenesisBlockHash = "AAA"
	genesisBlock := newBlock("genesis", bc.prevgenesisBlockHash)
	genesisBlockHash := genesisBlock.GetHash()
	bc.blockCache.Write(genesisBlock)
	bc.currentBlockHash = genesisBlockHash

	jsonBlock, _ := json.Marshal(bc.blockCache.Read(genesisBlockHash))
	log.Println("Generated GenesisBlock:", string(jsonBlock))
	return genesisBlockHash
}

//AddBlock sds
func (bc *Blockli) AddBlock(data string) string {
	newBlock := newBlock(data, bc.currentBlockHash)
	newBlockHash := newBlock.GetHash()
	bc.blockCache.Write(newBlock)
	bc.currentBlockHash = newBlockHash

	//jsonBlock, _ := json.Marshal(bc.blockarray[newBlockHash])
	//log.Println("Generated Block:",string(jsonBlock))
	return newBlockHash
}

func (bc *Blockli) getBlock(hash string) *Block {
	return bc.blockCache.Read(hash)
}

//ShowBlock sd
func (bc *Blockli) ShowBlock(hash string) string {
	jsonBlock, _ := json.Marshal(bc.blockCache.Read(hash))
	return string(jsonBlock)
}

//ValidateBlock sdsds
func (bc *Blockli) ValidateBlock(hash string) bool {
	//if hash genesis return true

	currentPreviousHash := bc.getBlock(hash).GetPreviousHash()
	genesisBlockHash := bc.prevgenesisBlockHash
	if currentPreviousHash == genesisBlockHash {
		return true
	}
	if bc.getBlock(hash) != nil {
		//else get prehash and validate
		return bc.ValidateBlock(currentPreviousHash)
	}
	log.Println("no genesis got", currentPreviousHash)
	return false

}

//GetGenesisHash dfd
func (bc *Blockli) GetGenesisHash() string {
	return bc.prevgenesisBlockHash
}
