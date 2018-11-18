package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//block valuable information.
type Block struct{
	Timestamp		int64
	Data			[]byte
	PrevBlockHash	[]byte
	Hash			[]byte
}

func (block *Block) SetHash() {
	Timestamp := []byte(strconv.FormatInt(block.Timestamp,10))
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp})

	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

