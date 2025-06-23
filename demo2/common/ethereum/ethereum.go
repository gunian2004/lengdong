package ethereum

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/coldchain-traceability-system/common/config"
	"github.com/coldchain-traceability-system/common/contracts" // 引入合约绑定代码
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Client          *ethclient.Client
	Contract        *contracts.Contracts
	ContractAddress common.Address
	Auth            *bind.TransactOpts
)

func InitEthereum() {
	// 获取以太坊节点 URL
	nodeURL := config.AppConfig.Ethereum.NodeURL
	var err error
	log.Println("正在连接以太坊节点...")
	Client, err = ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatalf("连接以太坊客户端失败: %v", err)
	}

	// 获取合约地址
	contractAddrStr := config.AppConfig.Ethereum.ContractAddress
	ContractAddress = common.HexToAddress(contractAddrStr)
	log.Printf("正在实例化合约 %s...", contractAddrStr)
	Contract, err = contracts.NewContracts(ContractAddress, Client)
	if err != nil {
		log.Fatalf("实例化合约失败: %v", err)
	}

	// 获取私钥
	privateKeyStr := config.AppConfig.Ethereum.PrivateKey
	// 去除 0x 前缀后传给 HexToECDSA
	log.Println("正在解析私钥..." + privateKeyStr)
	privateKeyStr = strings.TrimPrefix(privateKeyStr, "0x")
	fmt.Println("私钥", len(privateKeyStr))
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}
	Auth = bind.NewKeyedTransactor(privateKey)
	Auth.Value = big.NewInt(0)             // in wei
	Auth.GasLimit = uint64(30000000)       // in units
	Auth.GasPrice = big.NewInt(1000000000) // in wei

	log.Println("以太坊模块初始化成功")
}

