package parse

const BlockStructComponentNumber = 8

func ParsingBlock(clntResData interface{}) []string {
	var resultList []string

	blockHeight := clntResData.
	(map[string]interface{})["block"].
	(map[string]interface{})["header"].
	(map[string]interface{})["height"].(string)

	blockHash := clntResData.
	(map[string]interface{})["block_id"].
	(map[string]interface{})["hash"].(string)

	timeStamp := clntResData.
	(map[string]interface{})["block"].
	(map[string]interface{})["header"].
	(map[string]interface{})["time"].(string)

	dataHashMerkle := clntResData.
	(map[string]interface{})["block"].
	(map[string]interface{})["header"].
	(map[string]interface{})["data_hash"].(string)

	validatorsHash := clntResData.
	(map[string]interface{})["block"].
	(map[string]interface{})["header"].
	(map[string]interface{})["validators_hash"].(string)

	consensusHash := clntResData.
	(map[string]interface{})["block"].
	(map[string]interface{})["header"].
	(map[string]interface{})["consensus_hash"].(string)

	evidenceHash := clntResData.
	(map[string]interface{})["block"].
	(map[string]interface{})["header"].
	(map[string]interface{})["evidence_hash"].(string)

	proposerAddress := clntResData.
	(map[string]interface{})["block"].
	(map[string]interface{})["header"].
	(map[string]interface{})["proposer_address"].(string)

	resultList = append(resultList, blockHeight)
	resultList = append(resultList, blockHash)
	resultList = append(resultList, timeStamp)
	resultList = append(resultList, dataHashMerkle)
	resultList = append(resultList, validatorsHash)
	resultList = append(resultList, consensusHash)
	resultList = append(resultList, evidenceHash)
	resultList = append(resultList, proposerAddress)

	return resultList
}