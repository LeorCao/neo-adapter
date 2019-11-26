package neoTransaction

import (
	"errors"
	"fmt"
)

type TxIn struct {
	txID []byte
	vout []byte
}

// 获取交易ID
func (in TxIn) GetTxID() string {
	return reverseBytesToHex(in.txID)
}

// 获取对应索引
func (in TxIn) GetVout() uint16 {
	return littleEndianBytesToUint16(in.vout)
}

// 创建交易输入 并将字段值序列化
// vins : 交易输入
func newTxInForEmptyTrans(vins []Vin) ([]TxIn, error) {
	if vins == nil || len(vins) == 0 {
		return nil, errors.New("No input found when create an empty transaction!")
	}
	var ret []TxIn

	for _, v := range vins {
		txid, err := reverseHexToBytes(v.TxID)
		reverseByteArray(txid)
		if err != nil || len(txid) != 32 {
			return nil, errors.New("Invalid previous transaction id!")
		}
		vout := uint16ToLittleEndianBytes(v.Vout)
		ret = append(ret, TxIn{txid, vout})
	}
	return ret, nil
}

// 反序列化交易输入
// txBytes : 交易输入的序列化值
// index : 字段值在序列化数组中的索引
func decodeTxInFromRawTrans(txBytes []byte, index int) ([]TxIn, int, error) {
	var txIns = make([]TxIn, 0)
	vinCount := txBytes[index]
	index++
	if vinCount == 0 {
		return nil, index, errors.New("Invalid transaction vin count")
	}

	for i := byte(0); i < vinCount; i++ {
		var txIn = TxIn{}

		if index+32 > len(txBytes) {
			return nil, index, errors.New("Invalid transaction vin txid length")
		}
		txIn.txID = txBytes[index : index+32]
		index += 32
		if index+2 > len(txBytes) {
			return nil, index, errors.New("Invalid transaction vin vout length")
		}
		txIn.vout = txBytes[index : index+2]
		index += 2
		txIns = append(txIns, txIn)
	}
	return txIns, index, nil
}

// 将交易转换为字节数组
func (in TxIn) toBytes() ([]byte, error) {
	var ret []byte
	ret = append(ret, in.txID...)
	ret = append(ret, in.vout...)
	return ret, nil
}

// 将输入设置为空
func (in *TxIn) setEmpty() {
	in.txID = []byte{}
	in.vout = []byte{}
}

func (ti *TxIn) String() string {
	return fmt.Sprintf("{ txId : %x, vout : %x }", ti.txID, ti.vout)
}
