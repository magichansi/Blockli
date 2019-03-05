package blockli

import (
	"log"
)

//BlockCache asdadasda
type BlockCache struct {
	blockStore *BlockStore
	blockarray map[string]*Block
}

//NewBlockCache  sdadadaddasd
func NewBlockCache(savePath string) *BlockCache {
	dbh := new(BlockCache)
	dbh.blockStore = NewBlockStore(savePath)
	var err error
	if err != nil {
		log.Fatal(err)
		log.Fatal("01")
	}
	dbh.blockarray = make(map[string]*Block)
	return dbh
}

func (db *BlockCache) Read(key string) *Block {
	var returnBlock *Block

	returnBlock = db.blockarray[key]
	if returnBlock != nil {
		return returnBlock
	}

	returnBlock = db.blockStore.Retrieve(key)
	if returnBlock != nil {
		return returnBlock
	}
	return nil
}

func (db *BlockCache) Write(value *Block) string {
	db.blockarray[value.GetHash()] = value
	db.blockStore.Save(value.GetHash(), value)
	return value.GetHash()
}
