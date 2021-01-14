package pkg

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
)

//// json 文件转成 block
//func Json2Block(data []byte) (*common.Block, error){
//	block := &common.Block{}
//	err := JSONToMessage(data, block)
//	if err != nil {
//		return nil, err
//	}
//	return block, nil
//}

// 二进制文件转成 block
func Bytes2Block(data []byte) (*common.Block, error) {
	block := &common.Block{}
	err := proto.Unmarshal(data, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

//func JSONToMessage(bs []byte, msg proto.Message) error {
//	buffer := bytes.NewBuffer(bs)
//	err := protolator.DeepUnmarshalJSON(buffer, msg)
//	if err != nil {
//		return fmt.Errorf("failed to DeepUnmarshalJSON, err:%v", err)
//	}
//	return nil
//}
