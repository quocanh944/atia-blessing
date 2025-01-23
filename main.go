package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"main/utils/rpc"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://api.roninchain.com/rpc")
	if err != nil {
		log.Fatal(err)
	}

	fromAddress, privateKey := ImportWallet("")

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(2020))
	if err != nil {
		log.Fatal(err)
	}

	auth.GasLimit = uint64(100584 * 2)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice

	contractAddress := common.HexToAddress("0x9d3936dbd9a794ee31ef9f13814233d435bd806c")
	instance, err := rpc.NewAtiaBlessingTransactor(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	addressGain := common.HexToAddress("0xe3e012D5c02EA2F2DC2Aa75E101f07CDb623ba27")
	tx, err := instance.ActivateStreak(auth, addressGain)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

// Import Wallet
func ImportWallet(privateKey string) (common.Address, *ecdsa.PrivateKey) {
	importedPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := importedPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, importedPrivateKey
}
