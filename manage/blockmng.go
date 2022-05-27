package manage

import (
	"sync"

	"github.com/Moonyongjung/anchor/util"
)

var blockListInstance *BlockList
var blockListOnce sync.Once

//-- Recording block structure includes start/end block number and the entire merkle root 
//   of the result of each block header's merkle tree information
type BlockList struct {
	EndBlock string `json: end_block`
	StartBlock string `json: start_block`
	MerkleRoot string `json: merkle_root`
}

func BlockListMng() *BlockList {
	blockListOnce.Do( func() {
		blockListInstance = &BlockList{}
	})
	return blockListInstance
}

func (b *BlockList) NewLatestBlockHeight(EndBlock string) {
	b.EndBlock = EndBlock
}

func (b *BlockList) NowLatestBlockHeight() string{
	return b.EndBlock
}

func (b *BlockList) IncreaseLatestBlockHeight() {
	temp, err := util.ToInt(b.EndBlock)
	if err != nil {
		util.LogAnchor(err)
	}
	temp = temp + 1
	increasedBlockHeight := util.ToString(temp, "")
	b.EndBlock = increasedBlockHeight
}

func (b *BlockList) DecreaseLatestBlockHeight() {
	temp, err := util.ToInt(b.EndBlock)
	if err != nil {
		util.LogAnchor(err)
	}
	temp = temp - 1
	increasedBlockHeight := util.ToString(temp, "")
	b.EndBlock = increasedBlockHeight
}



