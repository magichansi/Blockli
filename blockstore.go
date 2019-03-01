package blockstore

import (
	"encoding/json"

	Block "github.com/magichansi/blockli/block"

	log "github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

//BlockStore sdad
type BlockStore struct {
	dbPath string
}

//New sdssdd
func New(path string) *BlockStore {

	returnStore := new(BlockStore)
	returnStore.dbPath = path
	return returnStore
}

//Save sdsdsd
func (BlockStore *BlockStore) Save(key string, value *Block.Block) {
	db, err := leveldb.OpenFile(BlockStore.dbPath, nil)
	if err != nil {
		log.WithFields(log.Fields{}).Fatal(err.Error())
	}
	defer db.Close()

	jsonBlock, _ := json.Marshal(value)
	err = db.Put([]byte(key), []byte(jsonBlock), nil)
	if err != nil {
		log.WithFields(log.Fields{}).Warning(err.Error())
	}
}

//Retrieve sdsds
func (BlockStore *BlockStore) Retrieve(key string) *Block.Block {
	db, err := leveldb.OpenFile(BlockStore.dbPath, nil)
	if err != nil {
		log.WithFields(log.Fields{}).Fatal(err.Error())
	}
	defer db.Close()

	data, err := db.Get([]byte(key), nil)
	if err != nil {
		log.WithFields(log.Fields{}).Warning(err.Error())
	}

	var returnBlock *Block.Block
	err = json.Unmarshal(data, &returnBlock)

	if err != nil {
		log.WithFields(log.Fields{}).Fatal(err.Error())
	}
	return returnBlock
}
