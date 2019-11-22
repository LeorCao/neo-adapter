package neoTransaction

import (
	"encoding/hex"
	"errors"
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

func InsertSignatureIntoEmptyTransaction(txHex string, txHashes []TxHash, unlockData, SegwitON bool) (string, error) {
	//txBytes, err := hex.DecodeString(txHex)
	//if err != nil {
	//	return "", errors.New("Invalid transaction hex data!")
	//}
	//
	//emptyTrans, err := DecodeRawTransaction(txBytes, SegwitON)
	//if err != nil {
	//	return "", err
	//}
	//
	//if unlockData == nil || len(unlockData) == 0 {
	//	return "", errors.New("No unlock data found!")
	//}
	//
	//if txHashes == nil || len(txHashes) == 0 {
	//	return "", errors.New("No signature data found!")
	//}
	//
	//if emptyTrans.Vins == nil || len(emptyTrans.Vins) == 0 {
	//	return "", errors.New("Invalid empty transaction,no input found!")
	//}
	//
	//if emptyTrans.Vouts == nil || len(emptyTrans.Vouts) == 0 {
	//	return "", errors.New("Invalid empty transaction,no output found!")
	//}
	//
	//if len(emptyTrans.Vins) != len(unlockData) {
	//	return "", errors.New("The number of transaction inputs and the unlock data are not match!")
	//}

	//segwit, err := isSegwit(unlockData, SegwitON)
	//if err != nil {
	//	return "", nil
	//}
	//if segwit && !SegwitON {
	//	return "", errors.New("Segwit transaction found while SegwitON is set to false!")
	//}
	//emptyTrans.Witness = segwit

	//for i := 0; i < len(emptyTrans.Vins); i++ {
	//	_, redeem, inType, err := checkScriptType(unlockData[i].LockScript, unlockData[i].RedeemScript)
	//	if err != nil {
	//		return "", nil
	//	}
	//	if inType == TypeP2PKH {
	//		emptyTrans.Vins[i].inType = int(inType)
	//		emptyTrans.Vins[i].scriptPub = nil
	//		script, err := txHashes[i].encodeToScript(redeem, SegwitON)
	//		if err != nil {
	//			return "", err
	//		}
	//		emptyTrans.Vins[i].scriptSig = script
	//	} else if inType == TypeP2WPKH {
	//		if !SegwitON {
	//			return "", errors.New("Pay to Witness-pubkey-hash type transaction found while the SegwitON is set to false!")
	//		}
	//		emptyTrans.Vins[i].inType = int(inType)
	//		emptyTrans.Vins[i].scriptPub = redeem
	//		script, err := txHashes[i].encodeToScript(redeem, SegwitON)
	//		if err != nil {
	//			return "", err
	//		}
	//		emptyTrans.Vins[i].scriptSig = script
	//	} else if inType == TypeBech32 {
	//		if !SegwitON {
	//			return "", errors.New("Bech32 type transaction found while the SegwitON is set to false!")
	//		}
	//		emptyTrans.Vins[i].inType = int(inType)
	//		emptyTrans.Vins[i].scriptPub = nil
	//		script, err := txHashes[i].encodeToScript(redeem, SegwitON)
	//		if err != nil {
	//			return "", err
	//		}
	//		emptyTrans.Vins[i].scriptSig = script
	//	} else if inType == TypeMultiSig {
	//		emptyTrans.Vins[i].inType = int(inType)
	//		//if segwit {
	//		//	redeemHash := owcrypt.Hash(redeem, 0, owcrypt.HASH_ALG_SHA256)
	//		//	redeemHash = append([]byte{0x23, 0x22, 0x00, 0x20}, redeemHash...)
	//		//	emptyTrans.Vins[i].scriptPub = redeemHash
	//		//} else {
	//		//	emptyTrans.Vins[i].scriptPub = nil
	//		//}
	//		emptyTrans.Vins[i].scriptSig = nil
	//		script, err := txHashes[i].encodeToScript(redeem, SegwitON)
	//		if err != nil {
	//			return "", err
	//		}
	//		emptyTrans.Vins[i].scriptMulti = script
	//	}
	//}

	//ret, err := emptyTrans.encodeToBytes(SegwitON)
	//if err != nil {
	//	return "", err
	//}

	//return hex.EncodeToString(ret), nil
	return "", nil
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

func VerifyRawTransaction(txHex string, SegwitON bool, addressPrefix AddressPrefix) bool {
	//txBytes, err := hex.DecodeString(txHex)
	//if err != nil {
	//	return false
	//}
	//
	//signedTrans, err := DecodeRawTransaction(txBytes, SegwitON)
	//if err != nil {
	//	return false
	//}
	//
	//if len(signedTrans.Vins) != len(unlockData) {
	//	return false
	//}
	//
	//emptyTrans := signedTrans.cloneEmpty()
	//
	//txHash, err := emptyTrans.getHashesForSig(unlockData, SegwitON, addressPrefix)
	//if err != nil {
	//	return false
	//}
	//
	//for i := 0; i < len(signedTrans.Vins); i++ {
	//	if signedTrans.Vins[i].inType == TypeP2PKH || signedTrans.Vins[i].inType == TypeP2WPKH || signedTrans.Vins[i].inType == TypeBech32 {
	//		sigpub, sigType, err := decodeFromScriptBytes(signedTrans.Vins[i].scriptSig)
	//		if err != nil {
	//			return false
	//		}
	//
	//		txHash[i].Normal.SigPub = *sigpub
	//		txHash[i].Normal.SigType = sigType
	//	} else if signedTrans.Vins[i].inType == TypeMultiSig {
	//		sigpub, sigType, err := decodeMultiBytes(signedTrans.Vins[i].scriptMulti)
	//		if err != nil {
	//			return false
	//		}
	//		for j := 0; j < len(sigpub); j++ {
	//			txHash[i].Multi[j].SigPub = sigpub[j]
	//			txHash[i].Multi[j].SigType = sigType[j]
	//		}
	//	}
	//}
	//
	//for _, t := range txHash {
	//	th, _ := hex.DecodeString(t.Hash)
	//	if t.NRequired == 0 {
	//		pubkey := owcrypt.PointDecompress(t.Normal.SigPub.Pubkey, owcrypt.ECC_CURVE_SECP256K1)[1:]
	//		if owcrypt.Verify(pubkey, nil, 0, th, 32, t.Normal.SigPub.Signature, owcrypt.ECC_CURVE_SECP256K1) != owcrypt.SUCCESS {
	//			return false
	//		}
	//	} else {
	//		count := 0
	//		for i := 0; i < int(t.NRequired); i++ {
	//			for j := count; j < len(t.Multi); j++ {
	//				pubkey := owcrypt.PointDecompress(t.Multi[j].SigPub.Pubkey, owcrypt.ECC_CURVE_SECP256K1)[1:]
	//				if owcrypt.Verify(pubkey, nil, 0, th, 32, t.Multi[i].SigPub.Signature, owcrypt.ECC_CURVE_SECP256K1) == owcrypt.SUCCESS {
	//					count++
	//					break
	//				}
	//			}
	//		}
	//		if count != int(t.NRequired) {
	//			return false
	//		}
	//	}
	//}
	return true
}
