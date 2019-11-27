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
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/openwallet"
)

type AddressDecoder interface {
	openwallet.AddressDecoder
	ScriptPubKeyToBech32Address(scriptPubKey []byte) (string, error)
}

type addressDecoder struct {
	wm *WalletManager //钱包管理者
}

func (decoder *addressDecoder) RedeemScriptToAddress(pubs [][]byte, required uint64, isTestnet bool) (string, error) {
	panic("implement me")
}

func (decoder *addressDecoder) ScriptPubKeyToBech32Address(scriptPubKey []byte) (string, error) {
	panic("implement me")
}

//NewAddressDecoder 地址解析器
func NewAddressDecoder(wm *WalletManager) *addressDecoder {
	decoder := addressDecoder{}
	decoder.wm = wm
	return &decoder
}

//PrivateKeyToWIF 私钥转WIF
func (decoder *addressDecoder) PrivateKeyToWIF(priv []byte, isTestnet bool) (string, error) {
	cfg := NEO_mainnetPrivateWIFCompressed
	if decoder.wm.Config.IsTestNet {
		cfg = NEO_testnetPrivateWIFCompressed
	}

	return addressEncoder.AddressEncode(priv, cfg), nil
}

//PublicKeyToAddress 公钥转地址
func (decoder *addressDecoder) PublicKeyToAddress(pub []byte, isTestnet bool) (string, error) {
	cfg := NEO_mainnetAddressP2PKH
	if decoder.wm.Config.IsTestNet {
		cfg = NEO_testnetAddressP2PKH
	}

	pub = append([]byte{0x21}, pub...)
	pub = append(pub, 0xac)

	sha256result := owcrypt.Hash(pub, 0, owcrypt.HASH_ALG_SHA256)
	pkHash := owcrypt.Hash(sha256result, 0, owcrypt.HASH_ALG_RIPEMD160)

	address := addressEncoder.AddressEncode(pkHash, cfg)

	return address, nil

}

//WIFToPrivateKey WIF转私钥
func (decoder *addressDecoder) WIFToPrivateKey(wif string, isTestnet bool) ([]byte, error) {

	cfg := NEO_mainnetPrivateWIFCompressed
	if decoder.wm.Config.IsTestNet {
		cfg = NEO_testnetPrivateWIFCompressed
	}

	priv, err := addressEncoder.AddressDecode(wif, cfg)
	if err != nil {
		return nil, err
	}

	return priv, err

}
