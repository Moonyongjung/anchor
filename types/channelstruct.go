package types

type ChannelStruct struct {
	HttpClientStartSignal chan bool
	CallContractChan chan CallContractChanStruct
	InvokeContractChan chan InvokeContractChanStruct
	CallContractResultChan chan interface{}
}

type CallContractChanStruct struct {
	InitBool bool
	CallName string
}

type InvokeContractChanStruct struct {
	InvokeName string
	MerkleRoot string
	StartBlock string
	EndBlock string
}