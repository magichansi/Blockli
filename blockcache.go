package blockcache

import (
	"log"

	"github.com/magichansi/blockli/block"
	"github.com/magichansi/blockli/blockstore"
)

//BlockCache asdadasda
type BlockCache struct {
	blockStore *blockstore.BlockStore
	blockarray map[string]*block.Block
}

//New  sdadadaddasd
func New(savePath string) *BlockCache {
	dbh := new(BlockCache)
	dbh.blockStore = blockstore.New(savePath)
	var err error
	if err != nil {
		log.Fatal(err)
		log.Fatal("01")
	}
	dbh.blockarray = make(map[string]*block.Block)
	return dbh
}

func (db *BlockCache) Read(key string) *block.Block {
	var returnBlock *block.Block

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

func (db *BlockCache) Write(value *block.Block) string {
	db.blockarray[value.GetHash()] = value
	db.blockStore.Save(value.GetHash(), value)
	return value.GetHash()
}
