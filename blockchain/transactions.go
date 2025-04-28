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

func makeTx(from, to string, amount int) (*Tx, error) {
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough balance")
	}
	var txIns []*TxIn
	var txOuts []*TxOut
	total := 0
	oldTxOuts := Blockchain().TxOutsByAddress(from)
	for _, TxOut := range oldTxOuts {
		if total >= amount {
			break
		}
		txIn := &TxIn{TxOut.Owner, TxOut.Amount}
		txIns = append(txIns, txIn)
		total += txIn.Amount
	}
	change := total - amount
	if change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	newTxOut := &TxOut{to, amount}
	txOuts = append(txOuts, newTxOut)
	tx := &Tx{
		Id:"",
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