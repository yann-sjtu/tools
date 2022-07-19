package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"strings"
	"sync/atomic"
)

func main() {
	input := []byte(`stakedex1h0j8x0v9hs4eq6ppgamemfyu4vuvp2sl0q9p3v`)
	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)
	fmt.Println(hex.EncodeToString(input), hex.EncodeToString([]byte(encodeString)))
	b, err := base64.StdEncoding.DecodeString("Ijk5OTk5OTAwIg==")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	b, err = hex.DecodeString("636f6e74726163745f696e666f")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	coin()

	a := " "
	ss := strings.Split(a, ",")
	fmt.Println(len(ss), ss)

	fmt.Println(strings.ToLower("0xB8FBB45bc2D8151EF6149ebE49E910866a72E6aA"))

	var tmp []int
	tmp2 := tmp[0:0]
	fmt.Println(tmp == nil, tmp2 == nil)

	test1()
}

func coin() {
	c := sdk.Coin{
		Denom:  "okt",
		Amount: sdk.NewDec(2800000000),
	}
	fmt.Println(c.String())
}

func test1() {
	var a int64
	var ss []int64
	defer atomic.AddInt64(&a, int64(len(ss)))
	ss = make([]int64, 10)
	fmt.Println(len(ss))
}
