package blockli

import "testing"

func Test_addBlockCache(t *testing.T) {

	blkch := NewBlockCache("C:/dev/Go")
	blk := newBlock("data", "key")
	hash := blkch.Write(blk)

	var newBlock *Block
	newBlock = blkch.Read(hash)

	if newBlock.GetData() != "data" {
		t.Errorf("Error Serialization. expect %s. got %s", "data", newBlock.GetData())
	}

}
