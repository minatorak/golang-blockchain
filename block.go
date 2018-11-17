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

fun (b *Block) SetHash() {
	Timestamp := []byte(strconv.FormatInt(b.Timestamp,10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp})
}

