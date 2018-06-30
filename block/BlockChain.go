package block

type Chain struct {
	Blocks []*Block
}

func (chain *Chain) PrintBlockChain() {
	for _, bl := range chain.Blocks {
		println(bl.Index)
		println(bl.PreHash)
		println(bl.Hash)
		println(bl.TimesTamp)
		println(bl.Data)
		println()
	}
}

func NewBlockChain() *Chain {
	geneticBlock := GenerateGeneticBlock()
	blockChain := Chain{}
	blockChain.AppendBlock(&geneticBlock)
	return &blockChain
}

func (chain *Chain) SendData(data string) {
	//geneticBlock := GenerateGeneticBlock()
	preBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := GenerateNewBlock(preBlock, data)
	chain.AppendBlock(&newBlock)
}

func (chain *Chain) AppendBlock(block *Block) {
	if len(chain.Blocks) <= 0 {
		chain.Blocks = append(chain.Blocks, block)
		return
	}

	if isValid(block, chain.Blocks[len(chain.Blocks)-1]) {
		chain.Blocks = append(chain.Blocks, block)
	} else {
		println("invalid block")
	}
}
func isValid(block *Block, old *Block) bool {
	if block.Index-1 != old.Index {
		return false
	}

	if block.PreHash != old.Hash {
		return false
	}

	if CaculateHash(block) != block.Hash {
		return false
	}
	return true
}
