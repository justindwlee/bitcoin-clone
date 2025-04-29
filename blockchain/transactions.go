package blockchain

import (
	"errors"
	"time"

	"github.com/justindwlee/bitcoinClone/utils"
)

const (
	minerReward int = 50
)

type mempool struct {
	Txs []*Tx
}

var Mempool *mempool = &mempool{}

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
	TxID string `json:"txId"`
	Index int	`json:"index"`
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner string `json:"owner"`
	Amount int	 `json:"amount"`
}

type UTxOut struct {
	TxID string `json:"txId"`
	Index int   `json:"index"`
	Amount int	`json:"amount"`
}

func isOnMempool (uTxOut *UTxOut) bool {
	exists := false
	Outer:
		for _, tx := range Mempool.Txs {
			for _, input := range tx.TxIns {
				if input.TxID == uTxOut.TxID && input.Index == uTxOut.Index {
					exists = true
					break Outer
				}
			}
		}
	return exists
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"", -1, "COINBASE"},
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

func makeTx(from, to string, amount int) (*Tx, error) {
	if BalanceByAddress(from, Blockchain()) < amount {
		return nil, errors.New("not enough balance")
	} 
	var txOuts []*TxOut
	var txIns []*TxIn
	total := 0
	uTxOuts := UTxOutsByAddress(from, Blockchain())

	for _, uTxOut := range uTxOuts {
		if total >= amount {
			break
		}
		txIn := &TxIn{uTxOut.TxID, uTxOut.Index, from}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}
	change := total - amount
	if change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)
	tx := &Tx{
		Id: "",
		Timestamp: int(time.Now().Unix()),
		TxIns: txIns,
		TxOuts: txOuts,
	}
	tx.getId()
	return tx, nil
}

func (m *mempool) AddTx(to string, amount int) error{
	tx, err := makeTx("nico", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

func (m *mempool) TxToConfirm() []*Tx {
	coinbaseTx := makeCoinbaseTx("nico")
	txs := m.Txs
	txs = append(txs, coinbaseTx)
	m.Txs = nil
	return txs
}