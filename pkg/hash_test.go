package pkg

import (
	"encoding/hex"
	"io/ioutil"
	"testing"
)

func TestHeaderHash(t *testing.T) {
	block11FilePath := "../data/block11.json"
	block12FilePath := "../data/block12.json"

	block11Content, err := ioutil.ReadFile(block11FilePath)
	if err != nil {
		t.Errorf("failed to open %s", block11FilePath)
	}

	block12Content, err := ioutil.ReadFile(block12FilePath)
	if err != nil {
		t.Errorf("failed to open %s", block12FilePath)
	}

	block11, err := Bytes2Block(block11Content)
	if err != nil {
		t.Errorf("Bytes2Block failed %s", string(block11Content))
	}

	block12, err := Bytes2Block(block12Content)
	if err != nil {
		t.Errorf("Bytes2Block failed %s", string(block12Content))
	}

	block11DataHash := hex.EncodeToString(block11.Header.DataHash)
	block11DataHashCal := hex.EncodeToString(BockDataHash(*block11, Sha256Hash))
	block11HashCal := hex.EncodeToString(BockHeaderHash(*block11, Sha256Hash))

	block12DataHash := hex.EncodeToString(block12.Header.DataHash)
	block12DataHashCal := hex.EncodeToString(BockDataHash(*block12, Sha256Hash))
	block11Hash := hex.EncodeToString(block12.Header.PreviousHash)

	if block11DataHash != block11DataHashCal{
		t.Errorf("block11DataHash:%s != block11DataHashCal:%s", block11DataHash, block11DataHashCal)
	}

	if block11HashCal != block11Hash{
		t.Errorf("block11HashCal:%s != block11Hash:%s", block11HashCal, block11Hash)
	}

	if block12DataHash != block12DataHashCal{
		t.Errorf("block12DataHash:%s != block12DataHashCal:%s", block12DataHash, block12DataHashCal)
	}
}

