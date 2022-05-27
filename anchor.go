package main

import (
	"os"
	"os/signal"
	"syscall"

	eclient "github.com/Moonyongjung/anchor/client"
	"github.com/Moonyongjung/anchor/gw"
	"github.com/Moonyongjung/anchor/util"
	"github.com/Moonyongjung/anchor/rpc"
	ctypes "github.com/Moonyongjung/anchor/types"
)

var configPath = "./config/config.json"
var configKeyPath = "./config/configKey.json"
var configContractPath = "./config/configContract.json"

var Channel ctypes.ChannelStruct

func init() {
	util.GetConfig().Read(configPath)
	util.GetConfig().Read(configKeyPath)
	util.GetConfig().Read(configContractPath)
}

func main() {
	Channel.HttpClientStartSignal = make(chan bool)
	Channel.CallContractChan = make(chan ctypes.CallContractChanStruct)
	Channel.InvokeContractChan = make(chan ctypes.InvokeContractChanStruct)
	Channel.CallContractResultChan = make(chan interface{})
	
	eclient.InitRpcClient()

	go rpc.InitCheckContractState(Channel)
	go rpc.SendInvokeTransaction(Channel)
	go rpc.SendCallTransaction(Channel)
	go gw.RunHttpServer(Channel)
	go gw.RunHttpClient(Channel)

	//-- If smart contract is changed, use this functions without operating rpc and gw in order to deploy
	// rpc.GolangBindings()
	// rpc.DeployAnchorContract()	

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	util.LogAnchor("Shutting down the server...")
	util.LogAnchor("Server gracefully stopped")	
}

