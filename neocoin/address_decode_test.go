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
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcrypt"
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

	expResult := "4a43e85f3e0137a23998cdc6dbacfac0268bf038"

	addr := "ANYZ11AmUfwiZFLbAWHoExFyBuqgLmfz88"

	cfg := NEO_mainnetAddressP2PKH

	hash, _ := addressEncoder.AddressDecode(addr, cfg)

	realResult := hex.EncodeToString(hash)

	if realResult != expResult {
		t.Logf("addr decode error : expected result is : %s, but real result is : %s", expResult, realResult)
	} else {
		t.Logf("hash: %s", hex.EncodeToString(hash))
	}

}

func TestAddressDecoder_PublicKeyToAddress(t *testing.T) {
	addr := "031a6c6fbbdf02ca351745fa86b9ba5a9452d785ac4f7fc2b7548ca2a46c4fcf4a"

	cfg := NEO_mainnetAddressP2PKH

	hash, err := addressEncoder.AddressDecode(addr, cfg)
	if err != nil {
		t.Errorf("AddressDecode failed unexpected error: %v\n", err)
		return
	}
	t.Logf("hash: %s", hex.EncodeToString(hash))
}

func TestAddressDecoder_ScriptPubKeyToBech32Address(t *testing.T) {

	scriptPubKey, _ := hex.DecodeString("002079db247b3da5d5e33e036005911b9341a8d136768a001e9f7b86c5211315e3e1")

	addr, err := scriptPubKeyToBech32Address(scriptPubKey, true)
	if err != nil {
		t.Errorf("ScriptPubKeyToBech32Address failed unexpected error: %v\n", err)
		return
	}
	t.Logf("addr: %s", addr)

	t.Logf("addr: %s", addr)
}

func TestAddressDecoder_WIFToP2WPKH_nested_in_P2SH(t *testing.T) {
	wif := "KxDgvEKzgSBPPfuVfw67oPQBSjidEiqTHURKSDL1R7yGaGYAeYnr"

	priv := "1dd37fba80fec4e6a6f13fd708d8dcb3b29def768017052f6c930fa1c5d90bbb"
	pub := "031a6c6fbbdf02ca351745fa86b9ba5a9452d785ac4f7fc2b7548ca2a46c4fcf4a"

	privkey, err := addressEncoder.AddressDecode(wif, NEO_mainnetPrivateWIFCompressed)
	if err != nil {
		t.Errorf("AddressDecode failed unexpected error: %v\n", err)
		return
	}
	if priv != hex.EncodeToString(privkey) {
		t.Failed()
	}
	t.Logf("privkey: %s", hex.EncodeToString(privkey))

	pubkey, _ := owcrypt.GenPubkey(privkey, owcrypt.ECC_CURVE_SECP256R1)
	pubkey = owcrypt.PointCompress(pubkey, owcrypt.ECC_CURVE_SECP256R1)

	if pub != hex.EncodeToString(pubkey) {
		t.Failed()
	}

	t.Logf("pubkey: %s", hex.EncodeToString(pubkey))

	pubkey = append([]byte{0x21}, pubkey...)
	pubkey = append(pubkey, 0xac)

	hash := owcrypt.Hash(pubkey, 0, owcrypt.HASH_ALG_SHA256)
	hash = owcrypt.Hash(hash, 0, owcrypt.HASH_ALG_RIPEMD160)

	t.Logf("hash: %s", hex.EncodeToString(hash))

	addr := addressEncoder.AddressEncode(hash, NEO_mainnetAddressP2PKH)

	t.Logf("addr: %s", addr)
}
