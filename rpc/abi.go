package rpc

import (
	"github.com/ethereum/go-ethereum/common"
	
	"github.com/Moonyongjung/anchor/util"
)

func GetAbiPack(callName string, merkleRoot string, startBlock string, endBlock string) []byte {
	contractAbi, err := AnchoringMetaData.GetAbi()
	if err != nil {
		util.LogAnchor(err)
	}	

	var abiByteData []byte
	if merkleRoot == "" {
		abiByteData, err = contractAbi.Pack(callName)
		if err != nil {
			util.LogAnchor(err)
		}
	} else {
		abiByteData, err = contractAbi.Pack(callName, merkleRoot, startBlock, endBlock)
		if err != nil {
			util.LogAnchor(err)
		}
	}

	return abiByteData
}

func GetAbiUnpack(callName string, data []byte) []interface{} {
	contractAbi, err := AnchoringMetaData.GetAbi()
	if err != nil {
		util.LogAnchor(err)
	}

	unpacked, err := contractAbi.Unpack(callName, data)
	if err != nil {
		util.LogAnchor(err)
	}

	return unpacked
}

func GetContractAddr() (common.Address) {
	contractDir := util.GetConfig().Get("contractDir")
	contractAddressDir := util.GetConfig().Get("contractAddressDir")

	var addressJsonStruct util.ContractAddressStruct
	addressData := util.JsonUnmarshal(addressJsonStruct, contractDir+contractAddressDir)
	address := addressData.(map[string]interface{})["ContractAddress"].(string)
	
	return util.AddressStringToByte20(address)
}