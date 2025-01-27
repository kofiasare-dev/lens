package db

import (
	"log"
	"sync"

	"github.com/dgraph-io/badger/v4"
)

type BadgerClient struct{ *badger.DB }

var badgerClient *BadgerClient
var setup sync.Once

func GetBadgerClient() *BadgerClient {
	setup.Do(func() {
		client, err := badger.Open(badger.DefaultOptions("tmp/badger"))
		if err != nil {
			log.Fatal(err)
		}

		badgerClient = &BadgerClient{client}
	})

	return badgerClient
}

func (bc *BadgerClient) Set(id string, data []byte) (err error) {
	err = bc.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(id), data)
	})

	return
}

func (bc *BadgerClient) Get(key string) (d []byte, err error) {
	err = bc.View(func(txn *badger.Txn) error {
		i, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		return i.Value(func(v []byte) error {
			d = append([]byte{}, v...)
			return nil
		})
	})

	return d, err
}

func (bc *BadgerClient) Delete(key string) error {
	return bc.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}
