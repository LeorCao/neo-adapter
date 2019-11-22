package neoTransaction

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
)

// 交易脚本
type TxScript struct {
	invocation   []byte // 调用脚本
	verification []byte // 验证脚本
}

func NewEmptyTxScript(invocation, verification []byte) *TxScript {
	return &TxScript{
		invocation:   invocation,
		verification: verification,
	}
}

// 构建验证脚本
func BuildVerification(pubkey string) ([]byte, error) {
	verif, err := hex.DecodeString(pubkey)
	if err != nil {
		return nil, err
	}
	verif = append([]byte{OpPushBytes33}, verif...)
	verif = append(verif, []byte{OpCheckSig}...)
	return verif, nil
}

func BuildInvocation(signByte []byte) []byte {
	signByte = append([]byte{OpPushBytes64}, signByte...)
	return signByte
}

func decodeTxScriptVerificationFromRawTrans(txByte []byte, index int) ([]TxScript, int, error) {
	var ret = make([]TxScript, 0)
	scriptsCount := txByte[index]
	index++
	for i := byte(0); i < scriptsCount; i++ {
		index++
		if index+35 > len(txByte) {
			return ret, index, errors.New("Invalid transaction tx script verification")
		}
		verificationScript := txByte[index : index+35]
		index += 35
		index++
		if index+65 > len(txByte) {
			return ret, index, errors.New("Invalid transaction tx script invocation")
		}
		invocationScript := txByte[index : index+65]
		index += 65
		ret = append(ret, TxScript{invocation: invocationScript, verification: verificationScript})
	}
	return ret, index, nil
}

// 构建调用脚本
func (ts *TxScript) buildInvocation() error {
	return nil
}

func (ts TxScript) toBytes() ([]byte, error) {
	var ret = make([]byte, 0)
	ret = append(ret, byte(len(ts.invocation)))
	ret = append(ret, ts.invocation...)
	//fmt.Println(fmt.Sprintf("Transaction tx script invocation %x", ts.invocation))
	ret = append(ret, byte(len(ts.verification)))
	ret = append(ret, ts.verification...)
	//fmt.Println(fmt.Sprintf("Transaction tx script verification %x", ts.invocation))
	return ret, nil
}

func (ts *TxScript) setEmpty() {
	ts.invocation = []byte{}
	ts.verification = []byte{}
}

func (ts *TxScript) String() string {
	return fmt.Sprintf("{invocation : %x, verification : %x}", ts.invocation, ts.verification)
}
