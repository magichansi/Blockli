package blockli

import "testing"

func Test_genesisBlockBlockli(t *testing.T) {
	bc := NewBlockli()
	bc.Init()
	if bc.prevgenesisBlockHash != "AAA" {
		t.Fail()
	}
}

func Test_addBlockBlockli(t *testing.T) {
	bc := NewBlockli()
	bc.Init()
	hash := bc.AddBlock("data")
	if bc.getBlock(hash).GetHash() != hash {
		t.Fail()
	}
	if bc.currentBlockHash != hash {
		t.Fail()
	}

	hash2 := bc.AddBlock("data2")
	if bc.getBlock(hash2).GetHash() != hash2 {
		t.Fail()
	}
	if bc.currentBlockHash != hash2 {
		t.Fail()
	}

}

func Test_validateBlockli(t *testing.T) {
	bc := NewBlockli()
	genhash := bc.Init()

	if !bc.ValidateBlock(genhash) {
		t.Errorf("Genesis validation failed. Got %s.", genhash)
	}

	hash := bc.AddBlock("data")
	if !bc.ValidateBlock(hash) {
		t.Errorf("validation failed.got %s", hash)
	}
}

func Benchmark_SpeedAddBlockli(b *testing.B) {
	bc := NewBlockli()
	bc.Init()
	for i := 0; i < 10000000; i++ {
		bc.AddBlock("data" + string(i))
	}
}

func Benchmark_SpeedValidateBlockli(b *testing.B) {
	bc := NewBlockli()
	bc.Init()
	for i := 0; i < 3000000; i++ {
		bc.AddBlock("data" + string(i))
	}
	hash := bc.AddBlock("data")
	if !bc.ValidateBlock(hash) {
		b.Errorf("validation failed.got %s", hash)
	}
}
