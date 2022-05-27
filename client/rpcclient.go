package client

import (
	"context"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	erpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/Moonyongjung/anchor/account"
	"github.com/Moonyongjung/anchor/manage"
	"github.com/Moonyongjung/anchor/util"
)

var keyStorePath = "./account/keystore/key.json"

type BcClient struct {
	Ctx    context.Context
	Client *ethclient.Client
}

func InitRpcClient() {
	if _, err := os.Stat(keyStorePath); err != nil {
		account.CreateAccountFromMnemonic()
	}
	util.GetConfig().Read(keyStorePath)
	block, _ := QueryBlock()
	currentNonce, _ := QueryNonce()
	chainId, _ := QueryChainId()
	suggestGasPrice, _ := QuerySuggestGasPrice()

	//-- Save info which are latest block number, account nonce, 
	//   Chain ID and Gas price
	util.GetConfigRpc().Set(block, currentNonce, chainId, suggestGasPrice)
	manage.NonceMng().NewNonce()
}

//-- Connect to Ethereum blockchain
//   For testing, connect to BSC based on ethereum core
func NewRpcClient() *BcClient {
	ctx, _ := context.WithCancel(context.Background())
	// defer cancel()

	//-- Target blockchain node URL
	bscUrlForTestnet := util.GetConfig().Get("bscUrlForTestnet")
	httpDefaultTransport := http.DefaultTransport
	defaultTransportPointer, ok := httpDefaultTransport.(*http.Transport)
	if !ok {
		util.LogRpcClient(ok)
	}
	defaultTransport := *defaultTransportPointer
	defaultTransport.DisableKeepAlives = true

	httpClient := &http.Client{Transport: &defaultTransport}
	rpcClient, err := erpc.DialHTTPWithClient(bscUrlForTestnet, httpClient)
	if err != nil {
		util.LogRpcClient(err)
	}

	ethClient := ethclient.NewClient(rpcClient)

	return &BcClient{ctx, ethClient}
}
