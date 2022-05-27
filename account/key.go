package account

import (
	"os/exec"
	"strings"

	"github.com/Moonyongjung/anchor/util"
)

func CreateAccountFromMnemonic() {
	var strList []string
	cmd := exec.Command("node", "./lib/ethers/generateAddr.js")
	output, err := cmd.Output()
	if err != nil {
		util.LogRpcClient(err)
	}
	strOutput := string(output)
	strList = strings.Split(strOutput, "\n")
	
	util.StoreKey(strList)	
}