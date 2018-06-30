package main

import "awesomeProject/block"

func main() {
	blockChain := block.NewBlockChain()
	blockChain.SendData("Hello")
	blockChain.SendData("World")

	blockChain.PrintBlockChain()
}
