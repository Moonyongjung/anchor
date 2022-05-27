package parse

import (
	"strings"

	"github.com/Moonyongjung/anchor/manage"
	ctypes "github.com/Moonyongjung/anchor/types"
	"github.com/Moonyongjung/anchor/util"
)

var blockStandard = "1" //Terra pisco testnet genesis height

func ParsingCallTransactionRes(callRes []interface{}, initBool bool, channel ctypes.ChannelStruct) {
	if initBool {
		logicInitCallTransaction(callRes, channel)
	} else {
		util.LogRpcClient("Request : ", callRes)
	}
}

func logicInitCallTransaction(callRes []interface{}, channel ctypes.ChannelStruct) {
	res := util.ToString(callRes, "")
	res = strings.Replace(res, " ", "", -1)
	if res == "[]" {	
		manage.BlockListMng().NewLatestBlockHeight(blockStandard)
		util.LogRpcClient("No recorded block info")
		util.LogRpcClient("genesis block height : " + blockStandard)
	} else {
		manage.BlockListMng().NewLatestBlockHeight(callRes[2].(string))
		util.LogRpcClient("Saved latest block height : " + callRes[2].(string))
		manage.BlockListMng().IncreaseLatestBlockHeight()
	}
	channel.HttpClientStartSignal <- true
}