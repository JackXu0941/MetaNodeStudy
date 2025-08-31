package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"learn1/store"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**
任务 2：合约代码生成 任务目标
使用 abigen 工具自动生成 Go 绑定代码，用于与 Sepolia 测试网络上的智能合约进行交互。
 具体任务
编写智能合约
使用 Solidity 编写一个简单的智能合约，例如一个计数器合约。
编译智能合约，生成 ABI 和字节码文件。
使用 abigen 生成 Go 绑定代码
安装 abigen 工具。
使用 abigen 工具根据 ABI 和字节码文件生成 Go 绑定代码。
使用生成的 Go 绑定代码与合约交互
编写 Go 代码，使用生成的 Go 绑定代码连接到 Sepolia 测试网络上的智能合约。
调用合约的方法，例如增加计数器的值。
输出调用结果。
*/

func main() {
	// deployContract()
	executeContract()
}

func deployContract() {

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/<your api key>")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("<your test private key>")

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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	// auth.GasLimit = uint64(81267177000000) // in units
	auth.GasPrice = gasPrice

	// input := "1.0"
	// address, tx, instance, err := store.DeployStore(auth, client, input)
	address, tx, instance, err := store.DeployStore(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract address:", address.Hex())
	fmt.Println("tranction hash:", tx.Hash().Hex())

	_ = instance

}

func executeContract() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/<your api key>")
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := store.NewStore(common.HexToAddress("your contract address"), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("<your test private key>")
	if err != nil {
		log.Fatal(err)
	}

	// var key [32]byte
	// var value [32]byte

	// copy(key[:], []byte("demo_save_key"))
	// copy(value[:], []byte("demo_save_value11111"))

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.SetCount(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Count(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract count:", valueInContract)

}
