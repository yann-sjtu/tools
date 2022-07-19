package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

const (
	dbpath    = "/Users/xzavier/go/src/github.com/okex/exchain/dev/_cache_evm/data/application.db"
	watchpath = "/Users/xzavier/go/src/github.com/okex/exchain/dev/_cache_evm/data/wasm-watcher.db"
	addr      = "ade4a5f5803a439835c636395a8d648dee57b2fc90d98dc17fa887159b69638b"
)

func main() {
	baddr, _ := hex.DecodeString(addr)
	fmt.Println(baddr)
	db, err := leveldb.OpenFile(dbpath, nil)
	if err != nil {
		panic(err)
	}
	it := db.NewIterator(&util.Range{}, nil)
	var index int
	_ = index
	fmt.Println([]byte("wasm"))
	for it.First(); it.Valid(); it.Next() {
		if bytes.Contains(it.Key(), []byte("s/k:wasm/")) {
			fmt.Println("index:", index)
			index++
			fmt.Println(it.Key())
		}
		if bytes.Contains(it.Key(), baddr) {
			fmt.Println("found:", it.Key())
			return
		}
		//continue
		//if !bytes.Contains(it.Key(), []byte("balanceex")) {
		//	continue
		//}

		//
		//fmt.Println("key:", string(it.Key()), it.Key())
		//fmt.Println("value:", string(it.Value()), it.Value())
	}
}
