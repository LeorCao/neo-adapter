/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package neocoin

import (
	"encoding/hex"
	"fmt"
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"testing"
)

func TestAddressDecoder_Encode(t *testing.T) {

	expResult := "ANYZ11AmUfwiZFLbAWHoExFyBuqgLmfz88"

	hash, _ := hex.DecodeString("4a43e85f3e0137a23998cdc6dbacfac0268bf038")

	cfg := NEO_mainnetAddressP2PKH

	addr := addressEncoder.AddressEncode(hash, cfg)

	if addr != expResult {
		t.Logf("addr encode error : expected result is : %s, but real result is : %s", expResult, addr)
	} else {
		t.Logf("addr: %s", addr)
	}

}

func TestAddressDecoder_Decode(t *testing.T) {

	expResult := "6eacd1e1d0d4885f42fdef0113273499ba9eccac"

	addr := "AXXYzk1kn9Bj8PHeqha921gqCpwJNRmuHC"

	cfg := NEO_mainnetAddressP2PKH

	hash, _ := addressEncoder.AddressDecode(addr, cfg)

	realResult := hex.EncodeToString(hash)

	if realResult != expResult {
		t.Logf("addr decode error : expected result is : %s, but real result is : %s", expResult, realResult)
	} else {
		t.Logf("hash: %s", hex.EncodeToString(hash))
	}

}

func initAddressDecode() *addressDecoder {
	tm := NewWalletManager()
	return NewAddressDecoder(tm)
}

// 测试私钥转WIF
func TestAddressDecoder_PrivateKeyToWIF(t *testing.T) {
	ad := initAddressDecode()
	privKeyBytes, err := hex.DecodeString("c0d97e2484b40e4a6f9cb471545973d635e495e1a469176d0604bdc62c441e1b")
	if err != nil {
		t.Error("Invalid private key!")
		return
	}
	wif, err := ad.PrivateKeyToWIF(privKeyBytes, true)
	if err != nil {
		t.Errorf("Private key to wif error : %s!", err.Error())
		return
	}
	wifExpect := "L3gau692aVdF8ESjuWKsoaew7Nu1uuRUfarf3VDc7LgkShCCkvyA"
	if wif != wifExpect {
		fmt.Println(fmt.Sprintf("wif not expected outcome!"))
	}
	fmt.Println(fmt.Sprintf("WIF : %s", wif))
}

// 测试公钥转地址
func TestAddressDecoder_PublicKeyToAddress(t *testing.T) {
	ad := initAddressDecode()
	pubKeyBytes, err := hex.DecodeString("036fb24ad1e6792d686fe8b37ceffbd1a5f2d30f05697bb10919bd778f91842a64")
	if err != nil {
		t.Error("Invalid public key!")
		return
	}
	addr, err := ad.PublicKeyToAddress(pubKeyBytes, true)
	if err != nil {
		t.Errorf("Public key to address error : %s", err.Error())
		return
	}
	addrExpect := "AemJEDk4ZAc6hvMWLrnrYigTsvKhujUGh2"
	if addr != addrExpect {
		t.Error("Public key to address not be expected outcome!")
		return
	}
	fmt.Println(fmt.Sprintf("Address : %s", addr))
}

// 测试WIF转私钥
func TestAddressDecoder_WIFToPrivateKey(t *testing.T) {
	ad := initAddressDecode()
	wifBytes := "L3gau692aVdF8ESjuWKsoaew7Nu1uuRUfarf3VDc7LgkShCCkvyA"
	privKeyBytes, err := ad.WIFToPrivateKey(wifBytes, true)
	if err != nil {
		t.Errorf("WIF to private key error : %s", err.Error())
	}

	privKeyExpect := "c0d97e2484b40e4a6f9cb471545973d635e495e1a469176d0604bdc62c441e1b"
	if hex.EncodeToString(privKeyBytes) != privKeyExpect {
		t.Error("WIF to private key not be expected outcome!")
	}
	fmt.Println(fmt.Sprintf("Private key : %s", hex.EncodeToString(privKeyBytes)))
}
