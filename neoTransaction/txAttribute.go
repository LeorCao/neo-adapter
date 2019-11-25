package neoTransaction

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
)

const (
	Attribute_Usage_ContractHash   = 0   // 合约脚本哈希 length 32
	Attribute_Usage_ECDH02         = 2   // 用于ECDH密钥交换的公钥 length 32
	Attribute_Usage_ECDH03         = 3   // 用于ECDH密钥交换的公钥 length 32
	Attribute_Usage_Script         = 32  // 交易额外的验证 length 20
	Attribute_Usage_Vote           = 48  // 投票payload length 需要指定 最大255个字节
	Attribute_Usage_DescriptionUrl = 129 // 描述说明的URL length 需要指定 最大255个字节
	Attribute_Usage_Description    = 144 // 说明 length 需要指定 最大255个字节
	// Usage 161 - 175 自定义的存储哈希 length 32
	// Usage 240 - 255 自定义的一般备注 length 需要指定 最大255个字节
)

type TxAttribute struct {
	usage  byte
	length []byte
	data   []byte
}

func newTxAttributeForEmptyTrans(attrs []Attribute) ([]TxAttribute, error) {

	ret := make([]TxAttribute, 0)

	if attrs == nil {
		return ret, nil
	}

	for _, attr := range attrs {
		data, err := hex.DecodeString(attr.Data)
		if err != nil {
			return nil, err
		}
		txAttr := TxAttribute{usage: attr.Attr.value}
		if attr.Attr.fixedDataLength != 0 {
			data = fixedDataLen(data, int(attr.Attr.fixedDataLength))
			txAttr.length = uint16ToLittleEndianBytes(uint16(len(data)))
		}
		txAttr.data = data
		ret = append(ret, txAttr)
	}
	return ret, nil
}

func decodeTxAttributeFromRawTrans(txByte []byte, index int) ([]TxAttribute, int, error) {
	var txAttrs = make([]TxAttribute, 0)
	var attrCount = txByte[index]
	index++
	if attrCount == 0 {
		return txAttrs, index, nil
	}

	for i := byte(0); i < attrCount; i++ {
		var txAttr = TxAttribute{}
		if index > len(txByte) {
			return nil, index, errors.New("Invalid transaction vout attribute length")
		}
		txAttr.usage = txByte[index]
		index++
		attrType := getAttributeTypeByUsage(txAttr.usage)
		if attrType.fixedDataLength == 0 {
			if index+2 > len(txByte) {
				return nil, index, errors.New("Invalid transaction vout attribute length")
			}
			txAttr.length = txByte[index : index+2]
			index += 2
			dataLen := int(littleEndianBytesToUint16(txAttr.length))
			if index+dataLen > len(txByte) {
				return nil, index, errors.New("Invalid transaction vout attribute length")
			}
			txAttr.data = txByte[index : index+dataLen]
			index += dataLen
			txAttrs = append(txAttrs, txAttr)
			continue
		}
		if index+int(attrType.fixedDataLength) > len(txByte) {
			return nil, index, errors.New("Invalid transaction vout attribute length")
		}
		txAttr.data = txByte[index : index+int(attrType.fixedDataLength)]
		txAttrs = append(txAttrs, txAttr)
	}
	return txAttrs, index, nil
}

func (ta TxAttribute) toBytes() ([]byte, error) {
	if ta.usage < 0 || ta.usage > 0xff {
		return nil, errors.New(fmt.Sprintf("Invalid usage value : %d", ta.usage))
	}
	ret := []byte{}
	ret = append(ret, ta.usage)
	lengthData := littleEndianBytesToUint16(ta.length)
	if lengthData > 0 {
		ret = append(ret, ta.length...)
	}
	ret = append(ret, ta.data...)
	return ret, nil
}

func (ta *TxAttribute) setEmpty() {
	ta.usage = byte(0)
	ta.length = []byte{}
	ta.data = []byte{}
}

func (tx *TxAttribute) String() string {
	return fmt.Sprintf("{ usage : %x, length : %x, data : %x }", tx.usage, tx.length, tx.data)
}
