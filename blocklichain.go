package blockli

import (
	"encoding/json"
	"log"

	"github.com/magichansi/blockli/blockcache"

	"github.com/magichansi/blockli/block"
)

type blockli struct {
	prevgenesisBlockHash string
	blockCache           *blockcache.BlockCache
	currentBlockHash     string
}

func New() *blockli {

	a := new(blockli)
	a.blockCache = blockcache.New("C:/dev/Go")
	return a
}

func (bc *blockli) Init() string {
	bc.prevgenesisBlockHash = "AAA"
	genesisBlock := block.New("genesis", bc.prevgenesisBlockHash)
	genesisBlockHash := genesisBlock.GetHash()
	bc.blockCache.Write(genesisBlock)
	bc.currentBlockHash = genesisBlockHash

	jsonBlock, _ := json.Marshal(bc.blockCache.Read(genesisBlockHash))
	log.Println("Generated GenesisBlock:", string(jsonBlock))
	return genesisBlockHash
}

func (bc *blockli) AddBlock(data string) string {
	newBlock := block.New(data, bc.currentBlockHash)
	newBlockHash := newBlock.GetHash()
	bc.blockCache.Write(newBlock)
	bc.currentBlockHash = newBlockHash

	//jsonBlock, _ := json.Marshal(bc.blockarray[newBlockHash])
	//log.Println("Generated Block:",string(jsonBlock))
	return newBlockHash
}

func (bc *blockli) getBlock(hash string) *block.Block {
	return bc.blockCache.Read(hash)
}

func (bc *blockli) ShowBlock(hash string) string {
	jsonBlock, _ := json.Marshal(bc.blockCache.Read(hash))
	return string(jsonBlock)
}

func (bc *blockli) ValidateBlock(hash string) bool {
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

func (bc *blockli) GetGenesisHash() string {
	return bc.prevgenesisBlockHash
}
