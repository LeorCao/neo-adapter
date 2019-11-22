package neoTransaction

type AddressPrefix struct {
	P2PKHPrefix  []byte
	P2WPKHPrefix []byte
	P2SHPrefix   []byte
	Bech32Prefix string
}

var (
	BTCMainnetAddressPrefix = AddressPrefix{[]byte{0x00}, []byte{0x05}, nil, "bc"}
	BTCTestnetAddressPrefix = AddressPrefix{[]byte{0x6F}, []byte{0xC4}, nil, "tb"}
)

// 交易类型
type TransactionType struct {
	jsonValue string
	hexValue  byte
	version   byte
}

var (
	MinerTransaction        = TransactionType{"MinerTransaction", 0x00, 0}
	IssueTransaction        = TransactionType{"IssueTransaction", 0x02, 0}
	ClaimTransaction        = TransactionType{"ClaimTransaction", 0x03, 0}
	DataFile                = TransactionType{"DataFile", 0x12, 0}
	EnrollmentTransaction   = TransactionType{"EnrollmentTransaction", 0x20, 0}
	RegisterTransaction     = TransactionType{"RegisterTransaction", 0x40, 0}
	ContractTransaction     = TransactionType{"ContractTransaction", 0x80, 0}
	RecordTransaction       = TransactionType{"RecordTransaction", 0x81, 0}
	StateTransaction        = TransactionType{"StateTransaction", 0x90, 0}
	StateUpdateTransaction  = TransactionType{"StateUpdateTransaction", 0x90, 0}
	StateUpdaterTransaction = TransactionType{"StateUpdaterTransaction", 0x91, 0}
	DestroyTransaction      = TransactionType{"DestroyTransaction", 0x18, 0}
	PublishTransaction      = TransactionType{"PublishTransaction", 0xd0, 0}
	InvocationTransaction   = TransactionType{"InvocationTransaction", 0xd1, 1}
)

// 交易附加参数类型
type AttributeType struct {
	jsonString      string
	value           byte
	maxDataLength   uint16
	fixedDataLength uint16
}

var (
	AttrContractHash   = AttributeType{"ContractHash", 0x00, 32, 32}
	AttrECDH02         = AttributeType{"ECDH02", 0x02, 32, 32}
	AttrECDH03         = AttributeType{"ECDH03", 0x03, 32, 32}
	AttrScript         = AttributeType{"Script", 0x20, 20, 20}
	AttrVote           = AttributeType{"Vote", 0x30, 0, 0}
	AttrDescriptionUrl = AttributeType{"DescriptionUrl", 0x81, 255, 0}
	AttrDescription    = AttributeType{"Description", 0x90, 65535, 0}

	AttrHash1  = AttributeType{"Hash1", 0xa1, 32, 32}
	AttrHash2  = AttributeType{"Hash2", 0xa2, 32, 32}
	AttrHash3  = AttributeType{"Hash3", 0xa3, 32, 32}
	AttrHash4  = AttributeType{"Hash4", 0xa4, 32, 32}
	AttrHash5  = AttributeType{"Hash5", 0xa5, 32, 32}
	AttrHash6  = AttributeType{"Hash6", 0xa6, 32, 32}
	AttrHash7  = AttributeType{"Hash7", 0xa7, 32, 32}
	AttrHash8  = AttributeType{"Hash8", 0xa8, 32, 32}
	AttrHash9  = AttributeType{"Hash9", 0xa9, 32, 32}
	AttrHash10 = AttributeType{"Hash10", 0xaa, 32, 32}
	AttrHash11 = AttributeType{"Hash11", 0xab, 32, 32}
	AttrHash12 = AttributeType{"Hash12", 0xac, 32, 32}
	AttrHash13 = AttributeType{"Hash13", 0xad, 32, 32}
	AttrHash14 = AttributeType{"Hash14", 0xae, 32, 32}
	AttrHash15 = AttributeType{"Hash15", 0xaf, 32, 32}

	AttrRemark   = AttributeType{"Remark", 0xf0, 65535, 0}
	AttrRemark1  = AttributeType{"Remark1", 0xf1, 65535, 0}
	AttrRemark2  = AttributeType{"Remark2", 0xf2, 65535, 0}
	AttrRemark3  = AttributeType{"Remark3", 0xf3, 65535, 0}
	AttrRemark4  = AttributeType{"Remark4", 0xf4, 65535, 0}
	AttrRemark5  = AttributeType{"Remark5", 0xf5, 65535, 0}
	AttrRemark6  = AttributeType{"Remark6", 0xf6, 65535, 0}
	AttrRemark7  = AttributeType{"Remark7", 0xf7, 65535, 0}
	AttrRemark8  = AttributeType{"Remark8", 0xf8, 65535, 0}
	AttrRemark9  = AttributeType{"Remark9", 0xf9, 65535, 0}
	AttrRemark10 = AttributeType{"Remark10", 0xfa, 65535, 0}
	AttrRemark11 = AttributeType{"Remark11", 0xfb, 65535, 0}
	AttrRemark12 = AttributeType{"Remark12", 0xfc, 65535, 0}
	AttrRemark13 = AttributeType{"Remark13", 0xfd, 65535, 0}
	AttrRemark14 = AttributeType{"Remark14", 0xfe, 65535, 0}
	AttrRemark15 = AttributeType{"Remark15", 0xff, 65535, 0}
)

const (
	PublicKeySize        = 33
	DefaultTxVersion     = uint32(0)
	MaxScriptElementSize = 520
)

const (
	NeoAssetId    = "c56f33fc6ecfcd0c225c4ab356fee59390af8560be0e930faebe74a6daff7c9b" // NEO 的 asset id
	NeoGasAssetId = "602c79718b16e442de58778e148d0b1084e3b2dffd5de6b7b16cee7969282de7" // NEOGAS 的 asset id
)

const (
	SequenceFinal        = uint32(0xFFFFFFFF)
	SequenceMaxBip125RBF = uint32(0xFFFFFFFD)
)

const (
	SegWitSymbol  = byte(0)
	SegWitVersion = byte(1)
	SigHashAll    = byte(1)
)

const (
	//OpCodeHash160     = byte(0xA9)
	//OpCodeEqual       = byte(0x87)
	//OpCodeEqualVerify = byte(0x88)
	//OpCodeCheckSig    = byte(0xAC)
	//OpCodeDup         = byte(0x76)
	//OpCode_1          = byte(0x51)
	//OpCheckMultiSig   = byte(0xAE)
	//OpPushData1       = byte(0x4C)
	//OpPushData2       = byte(0x4D)
	//OpPushData3       = byte(0x4E)

	OpPushBytes64 = byte(0x40)
	OpPushBytes33 = byte(0x21)
	OpCheckSig    = byte(0xac)

	OpPush2         = byte(0x60)
	OpCheckMultiSig = byte(0xae)
)

var (
	CurveOrder     = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0xBA, 0xAE, 0xDC, 0xE6, 0xAF, 0x48, 0xA0, 0x3B, 0xBF, 0xD2, 0x5E, 0x8C, 0xD0, 0x36, 0x41, 0x41}
	HalfCurveOrder = []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x5D, 0x57, 0x6E, 0x73, 0x57, 0xA4, 0x50, 0x1D, 0xDF, 0xE9, 0x2F, 0x46, 0x68, 0x1B, 0x20, 0xA0}
)
