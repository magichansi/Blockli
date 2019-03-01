package blockcache

import "testing"
import "blockli/Block"

func Test_add(t *testing.T) {

	blkch := New("C:/dev/Go")
	blk := Block.New("data", "key")
	blkch.insertIntoDB("a", blk)

	var newBlock *Block.Block
	newBlock = blkch.getFromDB("a")

	if newBlock.GetData() != "data" {
		t.Errorf("Error Serialization. expect %s. got %s", "data", newBlock.GetData())
	}

}
