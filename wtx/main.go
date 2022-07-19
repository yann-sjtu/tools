package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/okex/exchain/libs/tendermint/crypto/ed25519"
	"github.com/okex/exchain/libs/tendermint/mempool"
	"github.com/tendermint/go-amino"
)

var cdc *amino.Codec
var hexPrivate = "d322864e848a3ebbb88cbd45b163db3c479b166937f10a14ab86a3f860b0b0b64506fc928bd335f434691375f63d0baf97968716a20b2ad15463e51ba5cf49fe"
var privKey ed25519.PrivKeyEd25519

var addrs = []string {
	"0xbbE4733d85bc2b90682147779DA49caB38C0aA1F",
	"0x83D83497431C2D3FEab296a9fba4e5FaDD2f7eD0",
	"0x4C12e733e58819A1d3520f1E7aDCc614Ca20De64",
	"0x2Bd4AF0C1D0c2930fEE852D07bB9dE87D8C07044",
}

func init() {
	cdc = amino.NewCodec()
	mempool.RegisterMessages(cdc)

	b, _ := hex.DecodeString(hexPrivate)
	copy(privKey[:], b)
}

func main() {
	for _, addr := range addrs {
		//convertTxMessage(addr)
		convertWtxMessage(addr)
	}
}

func convertTxMessage(address string) {
	f, err := os.OpenFile(fmt.Sprintf("%s.txt", address), os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)

	f2, err := os.OpenFile(fmt.Sprintf("TxMessage-%s.txt", address), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	w := bufio.NewWriter(f2)


	for {
		data, _, err := r.ReadLine()
		if err != nil {
			return
		}
		raw, _ := hex.DecodeString(string(data))
		msg := mempool.TxMessage{
			Tx: raw,
		}

		if _, err = w.WriteString(hex.EncodeToString(cdc.MustMarshalBinaryBare(&msg))); err != nil {
			panic(err)
		}
		if err = w.WriteByte('\n'); err != nil {
			panic(err)
		}
	}
}

func convertWtxMessage(address string) {
	f, err := os.OpenFile(fmt.Sprintf("%s.txt", address), os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)

	f2, err := os.OpenFile(fmt.Sprintf("WtxMessage-%s.txt", address), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f2.Close()
	w := bufio.NewWriter(f2)


	for {
		data, _, err := r.ReadLine()
		if err != nil {
			return
		}
		raw, _ := hex.DecodeString(string(data))
		wtx := mempool.WrappedTx{
			Payload: raw,
			From: address,
			NodeKey: privKey.PubKey().Bytes(),
		}

		sig, err := privKey.Sign(append(wtx.Payload, wtx.From...))
		if err != nil {
			panic(err)
		}
		wtx.Signature = sig

		msg := mempool.WtxMessage{
			Wtx: &wtx,
		}

		if _, err = w.WriteString(hex.EncodeToString(cdc.MustMarshalBinaryBare(&msg))); err != nil {
			panic(err)
		}
		if err = w.WriteByte('\n'); err != nil {
			panic(err)
		}
	}
}