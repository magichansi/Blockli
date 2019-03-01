package blockcache_test

import "testing"
import "github.com/magichansi/blockli/block"

func Test_add(t *testing.T) {

	blkch := New("C:/dev/Go")
	blk := block.New("data", "key")
	hash := blkch.Write(blk)

	var newBlock *block.Block
	newBlock = blkch.Read(hash)

	if newBlock.GetData() != "data" {
		t.Errorf("Error Serialization. expect %s. got %s", "data", newBlock.GetData())
	}

}
