package neoTransaction

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/blocktree/go-owcrypt"
)

type Vin struct {
	TxID string
	Vout uint16
}

func (vin *Vin) String() string {
	return fmt.Sprintf("Attribute : { txid : %s, vout : %d } ", vin.TxID, vin.Vout)
}

type Vout struct {
	Asset   string
	Address string
	Value   uint64
}

func (vout *Vout) String() string {
	return fmt.Sprintf("Attribute : { asset : %s, address : %s, value : %d } ", vout.Asset, vout.Address, vout.Value)
}

type Attribute struct {
	Attr AttributeType
	Data string
}

func (attr *Attribute) String() string {
	return fmt.Sprintf("Attribute : { usage : %s, data : %s } ", attr.Attr.jsonString, attr.Data)
}

type Script struct {
	Invocation   string
	Verification string
}

func (s *Script) String() string {
	return fmt.Sprintf(" Script : { invocation : %s, verification : %s }", s.Invocation, s.Verification)
}

// 创建未签名的空交易
// txType : 交易类型
// vins : 交易输入
// vouts : 交易输出
// attrs : 交易附加属性
func CreateEmptyRawTransaction(txType TransactionType, vins []Vin, vouts []Vout, attrs []Attribute) (string, error) {

	emptyTrans, err := newEmptyTransaction(txType, vins, vouts, attrs)
	if err != nil {
		return "", err
	}

	txBytes, err := emptyTrans.encodeToBytes()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(txBytes), nil
}

func CreateRawTransactionHashForSig(txHex string) ([]TxHash, error) {
	txBytes, err := hex.DecodeString(txHex)
	if err != nil {
		return nil, errors.New("Invalid transaction hex string!")
	}

	emptyTrans, err := DecodeRawTransaction(txBytes)
	if err != nil {
		return nil, err
	}

	return emptyTrans.getHashesForSig()
}

// 签名原始交易
// rawTx : 组装获得的原始交易
// priKey : 签名的私钥
func SignRawTransaction(rawTx string, prikey []byte) (*SignaturePubkey, error) {
	hash, err := hex.DecodeString(rawTx)
	if err != nil {
		return nil, errors.New("Invalid transaction hash!")
	}

	return calcSignaturePubkey(hash, prikey)
}

// 合并签名数据到空交易
// rawTx : 原始空交易
// txHashes : 签名信息
func InsertSignatureIntoEmptyTransaction(rawTx string, txHashes []TxHash) (string, error) {
	txBytes, err := hex.DecodeString(rawTx)
	if err != nil {
		return "", errors.New("Invalid transaction hex data!")
	}

	emptyTrans, err := DecodeRawTransaction(txBytes)
	if err != nil {
		return "", err
	}

	if txHashes == nil || len(txHashes) == 0 {
		return "", errors.New("No signature data found!")
	}

	if emptyTrans.Vins == nil || len(emptyTrans.Vins) == 0 {
		return "", errors.New("Invalid empty transaction,no input found!")
	}

	if emptyTrans.Vouts == nil || len(emptyTrans.Vouts) == 0 {
		return "", errors.New("Invalid empty transaction,no output found!")
	}

	if emptyTrans.Scripts == nil {
		emptyTrans.Scripts = make([]TxScript, 0)
	}

	for _, txHash := range txHashes {
		script, err := createTxScript(txHash.Normal.SigPub.Pubkey, txHash.Normal.SigPub.Signature)
		if err != nil {
			return "", err
		}

		emptyTrans.Scripts = append(emptyTrans.Scripts, *script)
	}

	ret, err := emptyTrans.encodeToBytes()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(ret), nil
}

func SignatureRawTransaction(rawTransHex string, signatureData []SignaturePubkey) (*[]byte, error) {

	rawTxBytes, err := hex.DecodeString(rawTransHex)
	if err != nil {
		return nil, err
	}

	rtx, err := DecodeRawTransaction(rawTxBytes)
	if err != nil {
		return nil, err
	}
	if rtx.Scripts == nil {
		rtx.Scripts = make([]TxScript, 0)
	}

	for _, sd := range signatureData {
		verifiBytes, err := BuildVerification(hex.EncodeToString(sd.Pubkey))
		if err != nil {
			return nil, err
		}
		invocationBytes := BuildInvocation(sd.Signature)
		rtx.Scripts = append(rtx.Scripts, *(NewEmptyTxScript(invocationBytes, verifiBytes)))
	}
	sigRawTxBytes, err := rtx.encodeToBytes()
	if err != nil {
		return nil, err
	}
	fmt.Println(rtx.String())
	return &sigRawTxBytes, nil
}

// 验证交易签名
// signedRawTx : 添加签名信息的原始交易
func VerifyRawTransaction(signedRawTx string) bool {
	txBytes, err := hex.DecodeString(signedRawTx)
	if err != nil {
		return false
	}

	signedTrans, err := DecodeRawTransaction(txBytes)
	if err != nil {
		return false
	}

	txHash, err := signedTrans.getHashesForSig()
	if err != nil {
		return false
	}

	for _, t := range txHash {
		th, _ := hex.DecodeString(t.Hash)
		if t.NRequired == 0 {
			pubkey := owcrypt.PointDecompress(t.Normal.SigPub.Pubkey, owcrypt.ECC_CURVE_SECP256R1)[1:]
			if owcrypt.Verify(pubkey, nil, 0, th, 32, t.Normal.SigPub.Signature, owcrypt.ECC_CURVE_SECP256R1) != owcrypt.SUCCESS {
				return false
			}
		} else {
			count := 0
			for i := 0; i < int(t.NRequired); i++ {
				for j := count; j < len(t.Multi); j++ {
					pubkey := owcrypt.PointDecompress(t.Multi[j].SigPub.Pubkey, owcrypt.ECC_CURVE_SECP256R1)[1:]
					if owcrypt.Verify(pubkey, nil, 0, th, 32, t.Multi[i].SigPub.Signature, owcrypt.ECC_CURVE_SECP256R1) == owcrypt.SUCCESS {
						count++
						break
					}
				}
			}
			if count != int(t.NRequired) {
				return false
			}
		}
	}
	return true
}
