package neoTransaction

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// 测试获取验证脚本的公钥
func TestGetVerificationPubkey(t *testing.T) {
	verification := "21036943c02168ce22fb2e48a3f92dd72336d295e793a52633beba22ac46916dc201ac"
	verifiBytes, err := hex.DecodeString(verification)
	if err != nil {
		t.Error("Invalid verificationScript!")
		return
	}
	ts := TxScript{
		verificationScript: verifiBytes,
	}
	pubKey, err := ts.GetPubKeyByVerificationScript()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("Verfication public key is : %s", hex.EncodeToString(pubKey)))
}

// 测试通过调用参数脚本获取签名内容
func TestTxScript_GetSignatureByInvocationScript(t *testing.T) {
	invocationScript := "4077d8a721291d9c8b8d4a587735b8f8ee10d8acdd3b71fbaa49d8168d89d5a3aeff89e360062e7048adff9b3d2a744651e46545c61286e81a7db402b2f1fc41f5"
	invocationScriptBytes, err := hex.DecodeString(invocationScript)
	if err != nil {
		t.Error("Invalid invocation script")
		return
	}
	ts := TxScript{
		invocationScript: invocationScriptBytes,
	}
	sign, err := ts.GetSignatureByInvocationScript()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("Invocation script signature is : %s", hex.EncodeToString(sign)))
}
