package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct{
	Timestamp		int64
	Data			[]byte
	PrevBlockHash	[]byte
	Hash			[]byte
}

fun (block *Block) SetHash() {
	Timestamp := []byte(strconv.FormatInt(block.Timestamp,10))
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp})

	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}

