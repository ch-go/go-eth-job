package task1

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

    


func Task1() {
    client, err := ethclient.Dial("https://sepolia.infura.io/v3/ab2dcdfe5bc44813a9586d5b2cfb1bdd")
	if err != nil {
		log.Fatal(err)
	}

    privateKey, err := crypto.HexToECDSA("ca90c97c3f33ae0fce7e65df6def122d938811a86e82d762d8f7227e55c0f554")
    if err != nil {
        log.Fatal(err)
    }

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal(err)
    }

    value := big.NewInt(1000000000000000000) // in wei (1 eth)
    gasLimit := uint64(300000)                // in units
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    toAddress := common.HexToAddress("0xB18d678555129b4AEa2F0E8480BF7b3fF426556a")
    var data []byte
    tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

    chainID, err := client.ChainID(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
    if err != nil {
        log.Fatal(err)
    }
    // 发送交易
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

    // 查询区块
    balance, err := client.BalanceAt(context.Background(), fromAddress, nil)

	if err != nil {
		log.Fatal(err)
	}
	eth := balance.Div(balance, big.NewInt(1e18)) // Convert Wei to Ether
	fmt.Printf("Balance %v/n", balance)
	fmt.Printf("Balance: %s\n", eth.String())
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Header: %+v\n   txHash: %v\n", header, header.TxHash.Hex())

}

