package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/justindwlee/bitcoinClone/db"
	"github.com/justindwlee/bitcoinClone/utils"
)

type Block struct {
	Hash string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height int `json:"height"`
	Difficulty int `json:"difficulty"`
	Nonce int `json:"nonce"`
	Timestamp int `json:"timestamp"`
	Transactions []*Tx `json:"transactions"`
}



func persistBlock(b *Block){
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine(){
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		// fmt.Printf("\n\n\nTarget: %s\nHash: %s\nNonce: %d\nTimestamp: %d\n\n\n", target, hash, b.Nonce, b.Timestamp)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break;
		} else {
			b.Nonce++;
		}
	}
}

func createBlock (prevHash string, height int, diff int) *Block{
	block := Block{
		Hash: "",
		PrevHash: prevHash,
		Height: height,
		Difficulty: diff,
		Nonce: 0,
	}
	block.mine()
	block.Transactions = Mempool().TxToConfirm()
	persistBlock(&block)
	return &block
}