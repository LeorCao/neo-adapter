package neoTransaction

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
)

func byteArrayCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			return false
		}
	}
	return true
}

//reverseBytes endian reverse
func reverseBytes(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

//reverseHexToBytes decode a hex string to an byte array,then change the endian
func reverseHexToBytes(hexVar string) ([]byte, error) {
	if len(hexVar)%2 == 1 {
		return nil, errors.New("Invalid TxHash!")
	}
	ret, err := hex.DecodeString(hexVar)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

//reverseBytesToHex change the endian of the input byte array then encode it to hex string
func reverseBytesToHex(bytesVar []byte) string {
	return hex.EncodeToString(reverseBytes(bytesVar))
}

//uint16ToLittleEndianBytes
func uint16ToLittleEndianBytes(data uint16) []byte {
	tmp := [2]byte{}
	binary.LittleEndian.PutUint16(tmp[:], data)
	return tmp[:]
}

func uint16ToLittleEndianBytesAndReverse(data uint16) []byte {
	tmp := [2]byte{}
	binary.LittleEndian.PutUint16(tmp[:], data)
	return reverseByteArray(tmp[:])
}

func bytesFillZero(src []byte, minLength int) ([]byte, error) {
	if len(src) > minLength {
		return nil, errors.New(fmt.Sprintf("Source length more than min length!"))
	}
	ret := make([]byte, minLength)
	if len(src) < minLength {
		for i, v := range src {
			ret[i] = v
		}
	}
	if len(src) == minLength {
		ret = src
	}
	return ret, nil
}

//littleEndianBytesToUint16
func littleEndianBytesToUint16(data []byte) uint16 {
	return binary.LittleEndian.Uint16(data)
}

//uint32ToLittleEndianBytes
func uint32ToLittleEndianBytes(data uint32) []byte {
	tmp := [4]byte{}
	binary.LittleEndian.PutUint32(tmp[:], data)
	return tmp[:]
}

func uint32ToLittleEndianBytesAndReverse(data uint32) []byte {
	tmp := [4]byte{}
	binary.LittleEndian.PutUint32(tmp[:], data)
	return reverseByteArray(tmp[:])
}

// 反转byte数组
func reverseByteArray(arr []byte) []byte {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}

func reverseAssetId(hex string) string {
	cleanInput := cleanHexPrefix(hex)
	ret := ""
	for i := len(cleanInput) - 2; i >= 0; i -= 2 {
		ret += cleanInput[i : i+2]
	}
	return ret
}

func cleanHexPrefix(input string) string {
	if containsHexPrefix(input) {
		return input[2:]
	}
	return input
}

func containsHexPrefix(input string) bool {
	return input != "" && input != " " && input != "\n" &&
		len(input) > 1 && input[0] == '0' && input[1] == 'x'
}

//littleEndianBytesToUint32
func littleEndianBytesToUint32(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
}

func littleEndianBytesToUint32Reverse(data []byte) uint32 {
	return binary.LittleEndian.Uint32(reverseByteArray(data))
}

//uint64ToLittleEndianBytes
func uint64ToLittleEndianBytes(data uint64) []byte {
	tmp := [8]byte{}
	binary.LittleEndian.PutUint64(tmp[:], data)
	return tmp[:]
}

//uint64ToLittleEndianBytes
func uint64ToLittleEndianBytesAndReverse(data uint64) []byte {
	tmp := [8]byte{}
	binary.LittleEndian.PutUint64(tmp[:], data)
	reverseByteArray(tmp[:])
	return tmp[:]
}

//littleEndianBytesToUint64
func littleEndianBytesToUint64(data []byte) uint64 {
	return binary.LittleEndian.Uint64(data)
}

func fixedDataLen(src []byte, length int) []byte {
	if len(src) > length {
		return src[:length]
	} else {
		return src
	}
}

func writeLength(v int64) (ret []byte, err error) {
	if v < 0 {
		return ret, errors.New("Length is error")
	}
	if v < 0xfd {
		ret = append(ret, byte(v))
	} else if v <= 0xFFFF {
		ret = append(ret, byte(0xFD))
		ret = append(ret, byte(v))
	} else if v <= 0xFFFFFFFF {
		ret = append(ret, byte(0xFE))
		ret = append(ret, byte(v))
	} else {
		ret = append(ret, byte(0xFF))
		ret = append(ret, byte(v))
	}
	return ret, nil
}

func getAttributeTypeByUsage(usage byte) *AttributeType {
	switch usage {
	case AttrContractHash.value:
		return &AttrContractHash
	case AttrECDH02.value:
		return &AttrECDH02
	case AttrECDH03.value:
		return &AttrECDH03
	case AttrScript.value:
		return &AttrScript
	case AttrVote.value:
		return &AttrVote
	case AttrDescriptionUrl.value:
		return &AttrDescriptionUrl
	case AttrDescription.value:
		return &AttrDescription
	case AttrHash1.value:
		return &AttrHash1
	case AttrHash2.value:
		return &AttrHash2
	case AttrHash3.value:
		return &AttrHash3
	case AttrHash4.value:
		return &AttrHash4
	case AttrHash5.value:
		return &AttrHash5
	case AttrHash6.value:
		return &AttrHash6
	case AttrHash7.value:
		return &AttrHash7
	case AttrHash8.value:
		return &AttrHash8
	case AttrHash9.value:
		return &AttrHash9
	case AttrHash10.value:
		return &AttrHash10
	case AttrHash11.value:
		return &AttrHash11
	case AttrHash12.value:
		return &AttrHash12
	case AttrHash13.value:
		return &AttrHash13
	case AttrHash14.value:
		return &AttrHash14
	case AttrHash15.value:
		return &AttrHash15
	case AttrRemark.value:
		return &AttrRemark
	case AttrRemark1.value:
		return &AttrRemark1
	case AttrRemark2.value:
		return &AttrRemark2
	case AttrRemark3.value:
		return &AttrRemark3
	case AttrRemark4.value:
		return &AttrRemark4
	case AttrRemark5.value:
		return &AttrRemark5
	case AttrRemark6.value:
		return &AttrRemark6
	case AttrRemark7.value:
		return &AttrRemark7
	case AttrRemark8.value:
		return &AttrRemark8
	case AttrRemark9.value:
		return &AttrRemark9
	case AttrRemark10.value:
		return &AttrRemark10
	case AttrRemark11.value:
		return &AttrRemark11
	case AttrRemark12.value:
		return &AttrRemark12
	case AttrRemark13.value:
		return &AttrRemark13
	case AttrRemark14.value:
		return &AttrRemark14
	case AttrRemark15.value:
		return &AttrRemark15
	}
	return nil
}
