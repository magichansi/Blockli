package blocklichain

import "testing"

func Test_genesisBlock(t *testing.T) {
	bc := New()
	bc.Init()
	if bc.prevgenesisBlockHash != "AAA" {
		t.Fail()
	}
}

func Test_addBlock(t *testing.T) {
	bc := New()
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

func Test_validate(t *testing.T) {
	bc := New()
	genhash := bc.Init()

	if !bc.ValidateBlock(genhash) {
		t.Errorf("Genesis validation failed. Got %s.", genhash)
	}

	hash := bc.AddBlock("data")
	if !bc.ValidateBlock(hash) {
		t.Errorf("validation failed.got %s", hash)
	}
}

func Benchmark_SpeedAdd(b *testing.B) {
	bc := New()
	bc.Init()
	for i := 0; i < 10000000; i++ {
		bc.AddBlock("data" + string(i))
	}
}

func Benchmark_SpeedValidate(b *testing.B) {
	bc := New()
	bc.Init()
	for i := 0; i < 3000000; i++ {
		bc.AddBlock("data" + string(i))
	}
	hash := bc.AddBlock("data")
	if !bc.ValidateBlock(hash) {
		b.Errorf("validation failed.got %s", hash)
	}
}
