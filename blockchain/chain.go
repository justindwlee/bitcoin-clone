package blockchain

import (
	"sync"

	"github.com/justindwlee/bitcoinClone/db"
	"github.com/justindwlee/bitcoinClone/utils"
)

const (
	defaultDifficulty int = 2
	difficultyInterval int = 5
	blockInterval int = 2
	allowedRange int = 2
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height int `json:"height"`
	CurrentDifficulty int `json:"currentDifficulty"`
}

var b *blockchain
var once sync.Once


func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist(){
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height + 1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval - 1]
	actualTime := (newestBlock.Timestamp/60) - (lastRecalculatedBlock.Timestamp/60)
	expectedTime := blockInterval * difficultyInterval
	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height % difficultyInterval == 0 {
		return b.recalculateDifficulty()
	} else {
		return b.CurrentDifficulty
	}
}

func (b *blockchain) txOuts() []*TxOut{
	txOuts := []*TxOut{}
	blocks :=  b.Blocks()
	for _, block := range blocks {
		for _, tx := range block.Transactions {
			txOuts = append(txOuts, tx.TxOuts...)
		}
	}
	return txOuts
}

func (b *blockchain) TxOutsByAddress(address string) []*TxOut {
	allTxOuts := b.txOuts()
	var filteredTxOuts []*TxOut
	for _, txOut := range allTxOuts {
		if txOut.Owner == address {
			filteredTxOuts = append(filteredTxOuts, txOut)
		}
	}
	return filteredTxOuts
}

func (b *blockchain) BalanceByAddress(address string) int{
	addressTxOuts := b.TxOutsByAddress(address)
	amount := 0
	for _, txOut := range addressTxOuts {
		amount += txOut.Amount
	}
	return amount
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func(){
			b = &blockchain{
				Height: 0,
			}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock()
			} else {
				b.restore(checkpoint)
			}			
		})
	}
	return b
}
