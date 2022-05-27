# anchor
Private chain (based on Tendermint core) is anchoring to Public chain (based on Ethereum core)

## To start
```
go build anchor.go
./anchor
```

## Prerequisites
- `mnemonic` that use to create anchor account set in the `config/configKey.json`.
- Smart contract ABI & bytecode set in the `rpc/contracts`. (Default contract is set)
- If smart contract is changed, use functions(`rpc.GolangBindings(), rpc.DeployAnchorContract()`) in the `anchor.go` without operating RPC client and HTTP client/server in order to deploy
## Overview

### Configuration
For test, private chain is the Terra testnet (pisco) and public chain is the BSC testnet. In the config file(config.json), the BSC testnet is set as the target public chain URL, and it is recorded assuming the terra testnet is a private chain. However, since RPC comm. is requested using go-ethereum functions, it is possible to send transactions to other ethereum-based blockchains other than BSC. The private chain request used the [API](https://docs.tendermint.com/master/rpc/) of the blockchain using Tenderming core.

HTTP server is provided to call transaction, and API call Sample is in the `/gw/serverapisample`. `gatewayServerPort` in the `config.json` is the TCP port number of HTTP server. 

By adjusting `collectBlockCount`, can determine the number of blocks to be transferred to a transaction at one time.

### Recorded information in the smart contract of public chain

The information that are `EndBlock`, `StartBlock`, `MerkleRoot` is recorded to the smart contract of public chain. `MerkleRoot` is entire merkle root of the result of each block header's merkle tree information. i.e. Create merkle root by using info of each block header -> Collect merkle roots (collect count is included in the config file) -> When the number of collections reaches collect count, create merkle root again by using merkle roots. Merkle nodes are 
- `BlockHeight`
- `BlockHash`
- `TimeStamp`
- `DataHashMerkle`
- `ValidatorHash`
- `ConsensusHash`
- `EvidenceHash`
- `ProposerAddress`

So, it is not need to use entire merkle root. simply, can just save the normal merkle root. However, a lot of information is implied to reduce the number of transactions to reduce the gas cost. Verification using the entire merkle root can be verified by making the block headers recorded in the private chain as the merkle root and going through the decoding process twice.

