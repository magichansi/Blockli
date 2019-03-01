package block

import "time"
import "crypto/sha256"
import "encoding/hex"

type Block struct {
	PreviousHash string `json:"previousHash"`
	Data         string `json:"data"`
	Timestamp    int64  `json:"timestamp"`
	Hash         string `json:"hash"`
}

func New(data string, previousHash string) *Block {
	returnBlock := new(Block)
	returnBlock.Data = data
	returnBlock.Timestamp = time.Now().Unix()
	returnBlock.PreviousHash = previousHash
	a := []byte(data + previousHash + string(returnBlock.Timestamp))
	currentHash := sha256.Sum256(a)
	returnBlock.Hash = hex.EncodeToString(currentHash[:])
	return returnBlock
}

func (b Block) GetTimestamp() int64 {
	return b.Timestamp
}
func (b Block) GetPreviousHash() string {
	return b.PreviousHash
}
func (b Block) GetHash() string {
	return b.Hash
}
func (b Block) GetData() string {
	return b.Data
}
