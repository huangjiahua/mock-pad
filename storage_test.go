package main

import (
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestDB_Get(t *testing.T) {
	type fields struct {
		prefix string
		domain map[string]bool
		ldb    *leveldb.DB
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				prefix: tt.fields.prefix,
				domain: tt.fields.domain,
				ldb:    tt.fields.ldb,
			}
			if err := db.Get(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("DB.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
