package util

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/cbergoon/merkletree"
)

type Content struct {
	value string
}

func (c Content) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(c.value)); err != nil {
	  return nil, err
	}
  
	return h.Sum(nil), nil
}

func (c Content) Equals(other merkletree.Content) (bool, error) {
	return c.value == other.(Content).value, nil
}

func returnMerkleRoot(list []merkletree.Content) string {
	t, err := merkletree.NewTree(list)
	if err != nil {
		LogAnchor(err)
	}

	mr := t.MerkleRoot()
	rootHex := hex.EncodeToString(mr)

	return rootHex
}

//-- Merkle Node information
// type BlockListStruct struct {
// 	BlockHeight string
// 	BlockHash string
// 	TimeStamp string
// 	DataHashMerkle string
// 	ValidatorsHash string
// 	ConsensusHash string
// 	EvidenceHash string
// 	ProposerAddress string
// }
func MerkleRoot(parsedBlock []string) string {
	var list []merkletree.Content
	list = append(list, Content{value: parsedBlock[0]}) //BlockHeight
	list = append(list, Content{value: parsedBlock[1]}) //BlockHash
	list = append(list, Content{value: parsedBlock[2]}) //TimeStamp
	list = append(list, Content{value: parsedBlock[3]}) //DataHashMerkle
	list = append(list, Content{value: parsedBlock[4]}) //ValidatorHash
	list = append(list, Content{value: parsedBlock[5]}) //ConsensusHash
	list = append(list, Content{value: parsedBlock[6]}) //EvedenceHash
	list = append(list, Content{value: parsedBlock[7]}) //ProposerAddress

	return returnMerkleRoot(list)
}

//-- Super Merkle Root is merkle root of merkle roots
//   i.e. Create merkle root by using information of each block header 
//        -> Collect merkle roots (collect count is included in the config file)
//        -> When the number of collections reaches collect count, create merkle root again by using merkle roots
func SuperMerkleroot(merkleRootList []string) string {
	var list []merkletree.Content

	for _, merkleRoot := range(merkleRootList) {
		list = append(list, Content{value: merkleRoot})
	}

	return returnMerkleRoot(list)
}


