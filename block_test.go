package blockli

import "testing"

func Test_BlockConstructor(t *testing.T) {
	data := "data"
	testBlock := newBlock(data, "AAA")
	if testBlock.Data != data {
		t.Errorf("Data missmatch %s %s", testBlock.Data, data)
	}
	if testBlock.PreviousHash != "AAA" {
		t.Errorf("Hash missmatch %s %s", testBlock.PreviousHash, "AAA")
	}
}
