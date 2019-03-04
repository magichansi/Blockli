package block

import "testing"

func Test_Constructor(t *testing.T) {
	data := "data"
	testBlock := New(data, "AAA")
	if testBlock.Data != data {
		t.Errorf("Data missmatch %s %s", testBlock.Data, data)
	}
	if testBlock.PreviousHash != "AAA" {
		t.Errorf("Hash missmatch %s %s", testBlock.PreviousHash, "AAA")
	}
}
