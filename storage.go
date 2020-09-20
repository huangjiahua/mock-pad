package main

import (
	"encoding/json"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

const dbDirName = "mock-pad-db"

var (
	ErrNotFound   = errors.New("db: not found")
	ErrDbInternal = errors.New("db: internal error")
	ErrJson       = errors.New("db: json parser/encoder error")
)

type DB struct {
	prefix string
	domain map[string]bool
	ldb    *leveldb.DB
}

func NewDB(prefix string, domains []string) *DB {
	db := DB{}

	ldb, err := leveldb.OpenFile(prefix+dbDirName, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.prefix = prefix
	db.domain = make(map[string]bool)
	for _, v := range domains {
		db.domain[v] = true
	}
	db.ldb = ldb

	return &db
}

func (db *DB) Close() {
	_ = db.ldb.Close()
}

func (db *DB) Get(key string, value interface{}) error {
	data, err := db.ldb.Get([]byte(key), nil)
	if err != nil {
		return mapError(err)
	}
	err = json.Unmarshal(data, &value)
	if err != nil {
		return ErrJson
	}
	return nil
}

func (db *DB) Has(key string) bool {
	r, _ := db.ldb.Has([]byte(key), nil)
	return r
}

func (db *DB) Put(domain, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return ErrJson
	}
	k := domain + "-" + key
	err = db.ldb.Put([]byte(k), data, nil)
	if err != nil {
		return mapError(err)
	}
	return nil
}

func (db *DB) Write(b *Batch) error {
	err := db.ldb.Write(b.lb, nil)
	if err != nil {
		return ErrDbInternal
	}
	return nil
}

type Batch struct {
	lb *leveldb.Batch
}

func (b *Batch) Put(domain, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return ErrJson
	}
	k := domain + "-" + key
	b.lb.Put([]byte(k), data)
	return nil
}

func mapError(err error) error {
	if errors.Is(err, leveldb.ErrNotFound) {
		return ErrNotFound
	} else {
		return ErrDbInternal
	}
}
