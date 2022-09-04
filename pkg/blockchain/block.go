package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"

	tx "github.com/noodleslove/blockchain-go/pkg/transaction"
	"github.com/noodleslove/blockchain-go/pkg/utils"
)

type Block struct {
	Timestamp     int64
	Transactions  []*tx.Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// NewBlock creates and returns a block
func NewBlock(transactions []*tx.Transaction, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Transactions:  transactions,
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash[:]

	return block
}

// NewBlock creates and returns a genesis block
func NewGenesisBlock(coinbase *tx.Transaction) *Block {
	return NewBlock([]*tx.Transaction{coinbase}, []byte{})
}

// Serialize serializes a block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	utils.Check(err)

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	utils.Check(err)

	return &block
}
