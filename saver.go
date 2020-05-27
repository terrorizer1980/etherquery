package main

import (
	"encoding/json"
	log "github.com/cihub/seelog"
)

type Saver interface {
	SaveTransactionList(transactionList []Transaction) (int64, error)
}

type DummySaver struct {
	appConfig *AppConfig
	done      bool
}

func (s *DummySaver) SaveTransactionList(transactionList []Transaction) (int64, error) {
	if !s.done && len(transactionList) == 11 {
		marshal, _ := json.Marshal(transactionList)
		log.Infof("%v", string(marshal))
	}
	return int64(len(transactionList)), nil
}
