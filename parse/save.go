package parse

import (
	"time"

	"github.com/Moonyongjung/anchor/manage"
	ctypes "github.com/Moonyongjung/anchor/types"
	"github.com/Moonyongjung/anchor/util"
)

const requestBiggerHeightErr = "requested block height is bigger then the chain length"

var merklerootList []string
var heightList []string

func SaveBlock(responseData interface{}, savedLatestBlockHeight string, channel ctypes.ChannelStruct) {	
	errCheck := responseData.(map[string]interface{})["error"]

	if errCheck != nil {
		if errCheck.(string) == requestBiggerHeightErr {
			util.LogHttpClient("Waiting For creating new block...")
			manage.BlockListMng().DecreaseLatestBlockHeight()
			time.Sleep(time.Second)
		}
	} else if errCheck == nil {
		contractInvokeName := util.GetConfig().Get("contractInvokeName")
		collectBlockCount := util.GetConfig().Get("collectBlockCount")
		intCount, _ := util.ToInt(collectBlockCount)

		parsedBlock := ParsingBlock(responseData)
		util.LogHttpClient("Received Data : ", parsedBlock)
		merkleRoot := util.MerkleRoot(parsedBlock)

		heightList = append(heightList, parsedBlock[0])
		merklerootList = append(merklerootList, merkleRoot)

		util.LogHttpClient("Merkle Root List : ", merklerootList)

		if len(merklerootList) == intCount {
			superMerkleRoot := util.SuperMerkleroot(merklerootList)
			util.LogHttpClient("Save merkleroot : " + superMerkleRoot)
			util.LogHttpClient("Save startBlock : " + heightList[0])
			util.LogHttpClient("Save EndBlock : " + heightList[intCount-1])

			channel.InvokeContractChan <- ctypes.InvokeContractChanStruct{
				InvokeName: contractInvokeName,
				MerkleRoot: superMerkleRoot,
				StartBlock: heightList[0],
				EndBlock: heightList[intCount-1],
			}

			heightList = nil
			merklerootList = nil
		}
	}
	
}