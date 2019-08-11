package models

import (
	"github.com/anant-sharma/go-blockchain/controller/v1/blockchain"
	"github.com/go-bongo/bongo"
)

type transactionModel struct {
	bongo.DocumentBase `bson:",inline"`
	Checksum           string
	Data               string
	DataCategory       string
	Recipient          string
	Sender             string
	TransactionID      string
}

// SaveTransaction - Save New Transaction
func SaveTransaction(transaction blockchain.Transaction) {

	tx := &transactionModel{
		Checksum:      transaction.Checksum,
		Data:          transaction.Data,
		DataCategory:  transaction.DataCategory,
		Recipient:     transaction.Recipient,
		Sender:        transaction.Sender,
		TransactionID: transaction.TransactionID,
	}

	connection.Collection("transactions").Save(tx)
}
