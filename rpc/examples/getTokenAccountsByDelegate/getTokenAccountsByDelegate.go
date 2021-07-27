package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	pubKey := solana.MustPublicKeyFromBase58("AfkALUPjQp8R1rUwE6KhT38NuTYWCncwwHwcJu7UtAfV")
	out, err := client.GetTokenAccountsByDelegate(
		context.TODO(),
		pubKey,
		&rpc.GetTokenAccountsConfig{
			Mint: solana.MustPublicKeyFromBase58("So11111111111111111111111111111111111111112"),
		},
		nil,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(out)
}