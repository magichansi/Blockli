package block_test

import "testing"

func Test_Constructor(t *testing.T) {
	data := "data"
	testBlock := New(data, "AAA")
	if testBlock.data != data {
		t.Errorf("Data missmatch %s %s", testBlock.data, data)
	}
	if testBlock.previousHash != "AAA" {
		t.Errorf("Hash missmatch %s %s", testBlock.previousHash, "AAA")
	}
}
