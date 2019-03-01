package blockli

import "time"
import "crypto/sha256"
import "encoding/hex"

type Block struct {
	previousHash string `json:"previousHash"`
	data         string `json:"data"`
	timestamp    int64  `json:"timestamp"`
	Hash         string `json:"hash"`
}

func New(data string, previousHash string) *Block {
	returnBlock := new(Block)
	returnBlock.data = data
	returnBlock.timestamp = time.Now().Unix()
	returnBlock.previousHash = previousHash
	a := []byte(data + previousHash + string(returnBlock.timestamp))
	currentHash := sha256.Sum256(a)
	returnBlock.Hash = hex.EncodeToString(currentHash[:])
	return returnBlock
}

func (b Block) GetTimestamp() int64 {
	return b.timestamp
}
func (b Block) GetPreviousHash() string {
	return b.previousHash
}
func (b Block) GetHash() string {
	return b.Hash
}
func (b Block) GetData() string {
	return b.data
}
