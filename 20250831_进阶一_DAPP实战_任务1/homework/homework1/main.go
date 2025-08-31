package main

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

/**
任务 1：区块链读写 任务目标
使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。
具体任务
环境搭建
安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
查询区块
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
输出查询结果到控制台。
发送交易
准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
对交易进行签名，并将签名后的交易发送到网络。
输出交易的哈希值。
*/

func main() {
	SubscribeNewHead()
	// newNFTTRansaction()
}

func SubscribeNewHead() {

	fmt.Println("subscribe Start-------------------------------------------------------------------")
	//订阅区块需要 websocket RPC URL,这里用了一个 公共的 WSS URL
	client, err := ethclient.Dial("wss://ethereum.publicnode.com/?sepolia")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个新的通道，用于接收最新的区块头
	headers := make(chan *types.Header)
	//调用客户端的 SubscribeNewHead 方法，它接收我们刚创建的区块头通道，该方法将返回一个订阅对象。
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	//订阅将推送新的区块头事件到我们的通道，因此我们可以使用一个 select 语句来监听新消息。
	//订阅对象还包括一个 error 通道，该通道将在订阅失败时发送消息。
	//同时可以订阅到 多个 区块信息
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)

		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			//要获得该区块的完整内容，我们可以将区块头的摘要传递给客户端的 BlockByHash 函数。
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			//区块hash
			fmt.Println(block.Hash().Hex())
			//区块号转64位
			fmt.Println(block.Number().Uint64())
			//时间戳
			fmt.Println(block.Time())
			//nonce
			fmt.Println(block.Nonce())
			//交易数量
			fmt.Println(len(block.Transactions()))

		}

	}
	fmt.Println("subscribe Done-------------------------------------------------------------------")
}

func newNFTTRansaction() {

	fmt.Println("NFT 交易开始------------------------------------------------------------------")

	//连接客户端
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/<your api key>")
	if err != nil {
		log.Fatal(err)
	}

	//加载私钥
	privateKey, err := crypto.HexToECDSA("<your test private key>")
	if err != nil {
		log.Fatal(err)
	}

	//生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//生成 from address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 我们应该用于帐户交易的随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 设置要转移的 ETH 数量
	value := big.NewInt(10) // in wei (1 eth)
	// 设置 ETH 转账的燃气应设上限
	gasLimit := uint64(21000) // in units
	//SuggestGasPrice 函数，用于根据'x'个先前块来获得平均燃气价格。
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//设置 to address
	toAddress := common.HexToAddress("your test to address")
	var data []byte
	//调用 NewTransaction 来生成我们的未签名以太坊事务，这个函数需要接收 nonce，地址，值，燃气上限值，燃气价格和可选发的数据。
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	//从客户端拿到链 ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//使用 SignTx 来对事务进行签名，这个函数需要接收事务，签名器，私钥。
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//调用 SendTransaction 来将已签名的事务广播到整个网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

	fmt.Println("NFT 交易完成------------------------------------------------------------------")

}
