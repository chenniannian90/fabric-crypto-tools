package pkg

import (
	"bytes"
	"crypto/sha256"
	"encoding/asn1"
	"math/big"

	"github.com/Hyperledger-TWGC/ccs-gm/sm3"
	"github.com/hyperledger/fabric-protos-go/common"
)

type HashFunc func(data []byte) []byte

// sha256 哈希函数
func Sha256Hash(data []byte) []byte {
	res := sha256.Sum256(data)
	return res[:]
}

// 国密哈希函数
func Sm3Hash(data []byte) []byte {
	return sm3.SumSM3(data)
}

type asn1Header struct {
	Number       *big.Int
	PreviousHash []byte
	DataHash     []byte
}

func BlockHeaderBytes(b *common.BlockHeader) []byte {
	asn1Header := asn1Header{
		PreviousHash: b.PreviousHash,
		DataHash:     b.DataHash,
		Number:       new(big.Int).SetUint64(b.Number),
	}
	result, err := asn1.Marshal(asn1Header)
	if err != nil {
		panic(err)
	}
	return result
}

// 计算区块头部 hash
func BockHeaderHash(block common.Block, hashFunc HashFunc) []byte {
	headerBytes := BlockHeaderBytes(block.Header)
	hashData := hashFunc(headerBytes)
	return hashData
}

// 计算区块 data 部分 hash
func BockDataHash(block common.Block, hashFunc HashFunc) []byte {
	hashData := hashFunc(bytes.Join(block.Data.Data, nil))
	return hashData
}
