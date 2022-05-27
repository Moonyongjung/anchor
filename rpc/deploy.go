package rpc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Moonyongjung/anchor/client"
	"github.com/Moonyongjung/anchor/manage"
	"github.com/Moonyongjung/anchor/util"
)

// var gasLimit = 21000
var gasLimit = 1000000

type contractBytecode struct {
	contractBytecode string
}

func DeployAnchorContract() {
	//-- New RPC client
	bcClient := client.NewRpcClient()

	//-- Get config	
	address := util.GetConfig().Get("AccountAddress")
	priKey := util.GetConfig().Get("PrivateKey")
	_, _, chainId, gasPrice := util.GetConfigRpc().Get()

	accountNonce := manage.NonceMng().NowNonce()
	util.LogAnchor("Deploy smart contract")
	util.LogAnchor("Account Nonce : ", accountNonce)
	util.LogAnchor("Chain ID : ", chainId)	
	util.LogAnchor("Suggested gas price (wai) : ", gasPrice, "[usually 10gwai")	
	
	//-- User address, prikey
	byte20Address := util.AddressStringToByte20(address)
	util.LogAnchor("Account Address : ", byte20Address.Hex())

	//-- Remove 0x
	priKey = priKey[2:]
	priKeyCrypto, _ := crypto.HexToECDSA(priKey)

	//-- Generate Transactor and set transactor's values
	//-- ING... locl eth test chain id = 15
	// chainId = big.NewInt(15)
	contractAuth, _ := bind.NewKeyedTransactorWithChainID(priKeyCrypto, chainId)

	//-- ING... The below function not using chain ID
	// contractAuth := bind.NewKeyedTransactor(priKeyCrypto)
	contractAuth.Nonce = big.NewInt(int64(accountNonce))
	contractAuth.Value = big.NewInt(0)
	contractAuth.GasLimit = uint64(gasLimit)
	contractAuth.GasPrice = gasPrice
	util.LogAnchor("Contract Authentication info : ", contractAuth)

	//-- Deploy contract
	ctrtAddr, tx, _, err := DeployAnchoring(contractAuth, bcClient.Client)
	if err != nil {
		util.LogAnchor(err)
	}

	util.LogAnchor("Deploy Complete")
	util.LogAnchor("Contract address : ", ctrtAddr)
	util.LogAnchor("Deploy Transaction : ", tx)	
	
	//-- Auto create contract address json file
	util.StoreContractAddress(ctrtAddr)

}
