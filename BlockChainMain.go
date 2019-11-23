package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type BlockChain struct {
	blocks []*Block
}

type Block struct {
	Timestamp int64
	Hash      []byte
	Data      []byte
	PrevHash  []byte
}

func (b *Block) DeriveHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	info := bytes.Join([][]byte{b.Data, b.PrevHash, timestamp}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func NodeStart() *Block {
	return CreateBlock("NodeStart", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NodeStart()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after NodeStart")
	chain.AddBlock("Second Block after first")
	chain.AddBlock("Third Block after Second")
	chain.AddBlock("Four Block after Third")

	for _, block := range chain.blocks {
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
