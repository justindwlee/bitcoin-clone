package blockchain

import (
	"time"

	"github.com/justindwlee/bitcoinClone/utils"
)

const (
	minerReward int = 50
)

type Tx struct {
	Id string
	Timestamp int
	TxIns []*TxIn
	TxOuts []*TxOut
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	Owner string
	Amount int
}

type TxOut struct {
	Owner string
	Amount int	
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		Id: "",
		Timestamp: int(time.Now().Unix()),
		TxIns: txIns,
		TxOuts: txOuts,
	}
	tx.getId()
	return &tx
}