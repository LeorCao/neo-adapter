package neoTransaction

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
)

type TxOut struct {
	asset   []byte
	value   []byte
	address []byte
}

func newTxOutForEmptyTrans(vout []Vout) ([]TxOut, error) {
	if vout == nil || len(vout) == 0 {
		return nil, errors.New("No address to send when create an empty transaction!")
	}
	var ret []TxOut

	for _, v := range vout {
		assetId, err := hex.DecodeString(v.Asset)
		assetId = reverseByteArray(assetId)
		if err != nil {
			log.Error("Empty transaction asset to bytes error : ", err.Error())
		}
		value := uint64ToLittleEndianBytes(v.Value * 100000000)
		_, address, err := DecodeCheck(v.Address)
		fmt.Println(len(address))
		if err != nil {
			return nil, errors.New("Invalid address")
		}

		ret = append(ret, TxOut{assetId, value, address})
	}
	return ret, nil
}

func decodeTxOutFromRawTrans(txBytes []byte, index int) ([]TxOut, int, error) {
	var txOuts = make([]TxOut, 0)
	var voutCount = txBytes[index]
	index++
	if voutCount == 0 {
		return nil, index, errors.New("Invalid transaction vout count")
	}

	for i := byte(0); i < voutCount; i++ {
		var txOut = TxOut{}
		if index+32 > len(txBytes) {
			return nil, index, errors.New("Invalid transaction vout assetid length ")
		}
		txOut.asset = txBytes[index : index+32]
		index += 32
		if index+8 > len(txBytes) {
			return nil, index, errors.New("Invalid transaction vout value length ")
		}
		txOut.value = txBytes[index : index+8]
		index += 8
		if index+20 > len(txBytes) {
			return nil, index, errors.New("Invalid transaction vout out address length ")
		}
		txOut.address = txBytes[index : index+20]
		index += 20
		txOuts = append(txOuts, txOut)
	}
	return txOuts, index, nil
}

func (out TxOut) toBytes() ([]byte, error) {
	if out.value == nil || len(out.value) != 8 {
		return nil, errors.New("Invalid amount for a transaction output!")
	}

	ret := []byte{}
	ret = append(ret, out.asset...)
	//fmt.Println(fmt.Sprintf("Transaction out asset %x", out.asset))
	ret = append(ret, out.value...)
	//fmt.Println(fmt.Sprintf("Transaction out value %x", out.value))
	ret = append(ret, out.address...)
	//fmt.Println(fmt.Sprintf("Transaction out address %x", out.address))
	return ret, nil
}

func (to *TxOut) setEmpty() {
	to.asset = []byte{}
	to.value = []byte{}
	to.address = []byte{}
}

func (to *TxOut) String() string {
	return fmt.Sprintf("{ asset : %x, value : %x, address : %x}", to.asset, to.value, to.address)
}
