package blockli

import (
	"blockli/block"
	"blockli/blockcache"
	"encoding/json"
	"log"
)

type Blockli struct {
	prevgenesisBlockHash string
	blockCache           *blockcache.BlockCache
	currentBlockHash     string
}

func New() *Blockli {

	a := new(Blockli)
	a.blockCache = blockcache.New("C:/dev/Go")
	return a
}

func (bc *Blockli) Init() string {
	bc.prevgenesisBlockHash = "AAA"
	genesisBlock := block.New("genesis", bc.prevgenesisBlockHash)
	genesisBlockHash := genesisBlock.GetHash()
	bc.blockCache.Write(genesisBlock)
	bc.currentBlockHash = genesisBlockHash

	jsonBlock, _ := json.Marshal(bc.blockCache.Read(genesisBlockHash))
	log.Println("Generated GenesisBlock:", string(jsonBlock))
	return genesisBlockHash
}

func (bc *Blockli) AddBlock(data string) string {
	newBlock := block.New(data, bc.currentBlockHash)
	newBlockHash := newBlock.GetHash()
	bc.blockCache.Write(newBlock)
	bc.currentBlockHash = newBlockHash

	//jsonBlock, _ := json.Marshal(bc.blockarray[newBlockHash])
	//log.Println("Generated Block:",string(jsonBlock))
	return newBlockHash
}

func (bc *Blockli) getBlock(hash string) *block.Block {
	return bc.blockCache.Read(hash)
}

func (bc *Blockli) ShowBlock(hash string) string {
	jsonBlock, _ := json.Marshal(bc.blockCache.Read(hash))
	return string(jsonBlock)
}

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
	} else {
		log.Println("no genesis got", currentPreviousHash)
		return false
	}

}

func (bc *Blockli) GetGenesisHash() string {
	return bc.prevgenesisBlockHash
}
