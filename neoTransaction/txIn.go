package neoTransaction

import (
	"errors"
	"fmt"
)

type TxIn struct {
	txID []byte
	vout []byte
}

func (in TxIn) GetTxID() string {
	return reverseBytesToHex(in.txID)
}

func (in TxIn) GetVout() uint32 {
	return littleEndianBytesToUint32(in.vout)
}

func newTxInForEmptyTrans(vin []Vin) ([]TxIn, error) {
	if vin == nil || len(vin) == 0 {
		return nil, errors.New("No input found when create an empty transaction!")
	}
	var ret []TxIn

	for _, v := range vin {
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

func (in TxIn) toBytes() ([]byte, error) {
	var ret []byte
	ret = append(ret, in.txID...)
	ret = append(ret, in.vout...)
	return ret, nil
}

func (in *TxIn) setEmpty() {
	in.txID = []byte{}
	in.vout = []byte{}
}

func (ti *TxIn) String() string {
	return fmt.Sprintf("{ txId : %x, vout : %x }", ti.txID, ti.vout)
}
