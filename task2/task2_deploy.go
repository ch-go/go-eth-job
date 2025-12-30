package task2

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"task/v/task2/contract"
)

// Deploy deploys the Counter contract to Sepolia and prints the address/tx hash.
func Deploy() {
	rpcURL := os.Getenv("SEPOLIA_RPC_URL")
	privateKeyHex := os.Getenv("SEPOLIA_PRIVATE_KEY")

	if rpcURL == "" || privateKeyHex == "" {
		log.Fatal("missing env: SEPOLIA_RPC_URL, SEPOLIA_PRIVATE_KEY")
	}

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	address, tx, _, err := contract.DeployCounter(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deploy tx: %s\n", tx.Hash().Hex())
	fmt.Printf("contract address: %s\n", address.Hex())
}
