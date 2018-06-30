package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int64
	TimesTamp int64
	Hash      string
	PreHash   string

	Data string
}

func CaculateHash(block *Block) string {
	blockStr := string(block.Index) + string(block.TimesTamp) + block.PreHash + block.Data
	sum256 := sha256.Sum256([]byte(blockStr))
	return hex.EncodeToString(sum256[:])
}

func GenerateNewBlock(preBlock *Block, data string) Block {
	block := Block{}
	block.PreHash = preBlock.Hash
	block.TimesTamp = time.Now().Unix()
	block.Index = preBlock.Index + 1
	block.Data = data
	block.Hash = CaculateHash(&block)
	return block
}

func GenerateGeneticBlock() Block {

	oBlock := Block{}
	oBlock.Index = -1
	oBlock.Hash = ""

	return GenerateNewBlock(&oBlock, "Genetic Block")
}
