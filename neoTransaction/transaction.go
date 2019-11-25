package neoTransaction

import (
	"encoding/hex"
	"errors"
	"github.com/blocktree/go-owcrypt"
)

type Vin struct {
	TxID string
	Vout uint16
}

type Vout struct {
	/*
		{
			"n": 0,
			"asset": "0xc56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b",
			"value": "100",
			"address": "AJABcaJHDpbovCPCwszBkhK6QwWkC8ogWR"
		}
	*/

	Asset   string
	Address string
	Value   uint64
}

type Attribute struct {
	Attr AttributeType
	Data string
}

type Script struct {
	Invocation   string
	Verification string
}

type TxUnlock struct {
	LockScript string
}

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

func SignRawTransactionHash(txHash string, prikey []byte) (*SignaturePubkey, error) {
	hash, err := hex.DecodeString(txHash)
	if err != nil {
		return nil, errors.New("Invalid transaction hash!")
	}

	return calcSignaturePubkey(hash, prikey)
}

func InsertSignatureIntoEmptyTransaction(txHex string, txHashes []TxHash) (string, error) {
	txBytes, err := hex.DecodeString(txHex)
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

	script, err := createTxScript(txHashes[0].Normal.SigPub.Pubkey, txHashes[0].Normal.SigPub.Signature)
	if err != nil {
		return "", err
	}

	if emptyTrans.Scripts == nil {
		emptyTrans.Scripts = make([]TxScript, 0)
	}

	emptyTrans.Scripts = append(emptyTrans.Scripts, *script)

	ret, err := emptyTrans.encodeToBytes()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(ret), nil
}

func SignatureRawTransaction(rawTransHex, pubKey string, signatureData []byte) (*[]byte, error) {
	verifiBytes, err := BuildVerification(pubKey)
	if err != nil {
		return nil, err
	}

	invocationBytes := BuildInvocation(signatureData)

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
	rtx.Scripts = append(rtx.Scripts, *(NewEmptyTxScript(invocationBytes, verifiBytes)))
	sigRawTxBytes, err := rtx.encodeToBytes()
	if err != nil {
		return nil, err
	}
	return &sigRawTxBytes, nil
}

func VerifyRawTransaction(txHex string) bool {
	txBytes, err := hex.DecodeString(txHex)
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
