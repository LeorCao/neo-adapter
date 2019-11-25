package neoTransaction

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
)

// 交易脚本
type TxScript struct {
	invocationScript   []byte // 调用脚本
	verificationScript []byte // 验证脚本
}

func NewEmptyTxScript(invocation, verification []byte) *TxScript {
	return &TxScript{
		invocationScript:   invocation,
		verificationScript: verification,
	}
}

// 构建验证脚本
func BuildVerification(pubkey string) ([]byte, error) {
	verif, err := hex.DecodeString(pubkey)
	if err != nil {
		return nil, errors.New("Invalid public key!")
	}
	verif = append([]byte{OpPushBytes33}, verif...)
	verif = append(verif, []byte{OpCheckSig}...)
	return verif, nil
}

// 获取验证脚本中的公钥
func (ts *TxScript) GetPubKeyByVerificationScript() ([]byte, error) {
	if len(ts.verificationScript) != 35 {
		return nil, errors.New("Invalid verificationScript script length!")
	}
	return ts.verificationScript[1:34], nil
}

// 获取调用参数中的签名内容
func (ts *TxScript) GetSignatureByInvocationScript() ([]byte, error) {
	if len(ts.invocationScript) != 65 {
		return nil, errors.New("Invalid invocationScript script length")
	}
	return ts.invocationScript[1:], nil
}

// 构建参数脚本
func BuildInvocation(signByte []byte) []byte {
	signByte = append([]byte{OpPushBytes64}, signByte...)
	return signByte
}

func createTxScript(pubkey, signBytes []byte) (*TxScript, error) {
	invocation := BuildInvocation(signBytes)
	verification, err := BuildVerification(hex.EncodeToString(pubkey))
	if err != nil {
		return nil, err
	}
	return &TxScript{
		invocationScript:   invocation,
		verificationScript: verification,
	}, nil
}

func decodeTxScriptVerificationFromRawTrans(txByte []byte, index int) ([]TxScript, int, error) {
	var ret = make([]TxScript, 0)
	scriptsCount := txByte[index]
	index++
	for i := byte(0); i < scriptsCount; i++ {
		index++
		if index+65 > len(txByte) {
			return ret, index, errors.New("Invalid transaction tx script invocationScript")
		}
		invocationScript := txByte[index : index+65]
		index += 65
		index++
		if index+35 > len(txByte) {
			return ret, index, errors.New("Invalid transaction tx script verificationScript")
		}
		verificationScript := txByte[index : index+35]
		index += 35
		ret = append(ret, TxScript{invocationScript: invocationScript, verificationScript: verificationScript})
	}
	return ret, index, nil
}

func (ts TxScript) toBytes() ([]byte, error) {
	var ret = make([]byte, 0)
	ret = append(ret, byte(len(ts.invocationScript)))
	ret = append(ret, ts.invocationScript...)
	ret = append(ret, byte(len(ts.verificationScript)))
	ret = append(ret, ts.verificationScript...)
	return ret, nil
}

func (ts *TxScript) setEmpty() {
	ts.invocationScript = []byte{}
	ts.verificationScript = []byte{}
}

func (ts *TxScript) String() string {
	return fmt.Sprintf("{ invocationScript : %x, verificationScript : %x }", ts.invocationScript, ts.verificationScript)
}
