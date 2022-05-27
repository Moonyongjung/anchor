package rpc

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	eclient "github.com/Moonyongjung/anchor/client"
	"github.com/Moonyongjung/anchor/manage"
	"github.com/Moonyongjung/anchor/parse"
	ctypes "github.com/Moonyongjung/anchor/types"
	"github.com/Moonyongjung/anchor/util"
)


const defaultGasLimit = 1000000

func InitCheckContractState(channel ctypes.ChannelStruct) {
	channel.CallContractChan <- ctypes.CallContractChanStruct{
		InitBool: true,
		CallName: "callBlockList",
	}
}

func SendCallTransaction(channel ctypes.ChannelStruct) {
	for {
		select {
		case callContractInfo := <- channel.CallContractChan:
			//-- Input for Call contract is empty string
			callName := callContractInfo.CallName
			client := eclient.NewRpcClient()
			fromAddress := util.GetConfig().Get("AccountAddress")			
			byte20Address := util.AddressStringToByte20(fromAddress)
			_, _, _, gasPrice := util.GetConfigRpc().Get()

			byteData := GetAbiPack(callName, "", "", "")
			contractAddress := GetContractAddr()
			util.LogRpcClient("Call Contract address : " + contractAddress.Hex())

			msg := ethereum.CallMsg {
				From: byte20Address,
				To: &contractAddress,
				Gas: uint64(defaultGasLimit),
				GasPrice: gasPrice,
				Value: big.NewInt(0),
				Data: byteData,
			}

			res, err := client.Client.CallContract(client.Ctx, msg, nil)
			if err != nil {
				util.LogRpcClient(err)
			}

			result := GetAbiUnpack(callName, res)

			util.LogRpcClient("Contract Reponse : ", result)

			parse.ParsingCallTransactionRes(result, callContractInfo.InitBool, channel)

			if !callContractInfo.InitBool {
				channel.CallContractResultChan <- result	
			}
		}
	}
}



func SendInvokeTransaction(channel ctypes.ChannelStruct) {
	for {
		select {
		case invokcContractInfo := <- channel.InvokeContractChan:
			//-- input string
			invokeName := invokcContractInfo.InvokeName
			merkleRoot := invokcContractInfo.MerkleRoot
			startBlock := invokcContractInfo.StartBlock
			endBlock := invokcContractInfo.EndBlock
			
			client := eclient.NewRpcClient()

			//-- Private key for Transaction signing
			priKey := util.GetConfig().Get("PrivateKey")			
			priKey = priKey[2:]
			priKeyCrypto, _ := crypto.HexToECDSA(priKey)
			_, _, chainId, gasPrice := util.GetConfigRpc().Get()
			accountNonce := manage.NonceMng().NowNonce()

			byteData := GetAbiPack(invokeName, merkleRoot, startBlock, endBlock)
			contractAddress := GetContractAddr()

			tx := types.NewTransaction(
				accountNonce, 
				contractAddress, 
				big.NewInt(0),
				uint64(defaultGasLimit),
				gasPrice,
				byteData)	
			
			//-- ING... locl eth test chain id = 15
			// chainId = big.NewInt(15)
			signer := types.NewEIP155Signer(chainId)

			//-- Also, the version of blockchain based on ethereum is able to exist which don't use signer.
			//   the below function do not support signer.
			// signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, priKeyCrypto)
			signedTx, err := types.SignTx(tx, signer, priKeyCrypto)
			if err != nil {
				util.LogRpcClient(err)
			}

			err = client.Client.SendTransaction(client.Ctx, signedTx)
			if err != nil {
				util.LogRpcClient(err)
			}

			util.LogRpcClient("Account nonce : " + util.Uint64ToString(accountNonce))
			util.LogRpcClient("Transaction hash : " + signedTx.Hash().Hex())
			
			//-- Need to increase account nonce after sending transaction.
			manage.NonceMng().AddNonce()

			go waitTransactionReceipt(client, signedTx)
		}
	}
}

func waitTransactionReceipt(client *eclient.BcClient, signedTx *types.Transaction) {
	//-- Count is wait time (sec)
	count := 100
	for {
		receipt, err := client.Client.TransactionReceipt(client.Ctx, signedTx.Hash())
		if err != nil {
			count = count - 1
			if count < 0 {
				util.LogRpcClient("Block not mined, Timeout")
			}
			util.LogRpcClient("Transaction receipt is not arrived yet...")
			time.Sleep(time.Second*1)
		} else {
			util.LogRpcClient("Transaction Receipt : ", receipt)
			util.LogRpcClient("Transaction Hash : " , receipt.TxHash)
			util.LogRpcClient("Gas Used : " , receipt.GasUsed)
			util.LogRpcClient("Cumulative gas used", receipt.CumulativeGasUsed)
			break
		}
	}
}