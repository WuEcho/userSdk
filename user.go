package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	kzgSdk "github.com/domicon-labs/kzg-sdk"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/kzg"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

var (
	Url           string = "https://eth-sepolia.g.alchemy.com/v2/-t67_L9EE802yd-RZYxsZ38XRcJOCHfq"
	PrivateKey    string = "3180b6cc1ef8d68c00dc30c83b9f00321a60dbeeac202e7671312dc0cd9707b9"
	TokenContractAddress string = "0x3e0F3B6235fC7f8F0Dd5d74e026f7e08c07D0557"
	StorageManagementAddress string = "0x339D9a194f1f6882Fa5CD91538e909823A290A7f"
	ApproveValue   uint64 = 1000000000000000
	DomiconCommitmentContractAddress string = "0xC02C16bb85DF4290882A2587fBE58C47F14abDd5"
	DomiconNodeContractAddress  string = "0xd49912EEa34a44BaC11464fFCc7AB0679018E2C6"
	ChainID        int64 = 11155111
	Daskey            string = "0xa54af72a7b9f92d8d3a7c2ad0a6b4f84275d1fef612ab3b1297fbf8a31815ba3"
	RegistNodeAddr    string = ""
	BroadCastKeyStorePth string = "./keystoreFile"
	StorageKeyStorePth   string = "./keystoreFileS"
	PWD                   string = "123456"
	BroadCastURL           string = "43.203.215.230:8545"
	StorageURL              string = "43.201.29.243:8545"
)
const dSrsSize = 1 << 16
func PrivateKeyToAddress(key string) (*ecdsa.PrivateKey,common.Address) {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	// 创建一个签名交易的发送者
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return privateKey,fromAddress
}

type UserService struct {
	ctx  context.Context
	client  *ethclient.Client
	priv *ecdsa.PrivateKey
	localFile  *os.File
}

func NewUserService(url string,priv string) *UserService {
	client,_ := ethclient.Dial(url)
	private,_ := PrivateKeyToAddress(PrivateKey)
	// 创建文件
	file, err := os.Create("backendInfo.txt")
	if err != nil {
		fmt.Println("无法创建文件:", err)
	}

	return &UserService{
		ctx: context.Background(),
		client: client,
		priv: private,
		localFile: file,
	}
}

func (u *UserService) SendDAToExec(index uint64,length uint64,data []byte,sign [][]byte,commit string,dasKey [32]byte) {
	contractAddress := common.HexToAddress(DomiconCommitmentContractAddress)
	instance,err := NewDomicCommit(contractAddress,u.client)
	if err != nil {
		errStr := fmt.Sprintf("cant create contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(u.priv, big.NewInt(ChainID)) // For Mainnet
	if err != nil {
		log.Fatal(err)
	}
	// dasKey  签名所在组织的hash key
	var digest kzg.Digest
	str, _ := hex.DecodeString(commit)
	digest.SetBytes(str)
	commitData := PairingG1Point{
		X: new(big.Int).SetBytes(digest.X.Marshal()),
		Y: new(big.Int).SetBytes(digest.Y.Marshal()),
	}
	tx,err := instance.SubmitCommitment(auth,index,length,dasKey,sign,commitData)
	if err != nil {
		log.Fatal(err)
	}
	// 等待交易被打包并获取交易哈希
	fmt.Println("交易哈希:", tx.Hash().Hex())


	// 等待交易被确认
	receipt, err := bind.WaitMined(u.ctx, u.client, tx)
	if err != nil {
		errStr := fmt.Sprintf("cant WaitMined by contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	// 检查交易状态SignData
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("交易失败")
	}
	fmt.Println("交易成功!")

}

func FilterQuery()  {
	client, err := ethclient.Dial(Url)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(60000), // 起始区块编号
		ToBlock:   nil,                             // 终止区块为nil表示最新区块
		Addresses: []common.Address{
			common.HexToAddress(DomiconCommitmentContractAddress), // 合约地址
		},
		Topics:    [][]common.Hash{{common.HexToHash("0xfb3a62c71e56d41dd4a93b72df383612d4377bd1d2c2466402c09092931815a9")}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to retrieve logs: %v", err)
	}
	for _, vLog := range logs {
		//event := new(SendDACommitment)
		//err := event.Unpack(vLog.Data)
		//if err != nil {
		//	log.Fatalf("Failed to unpack log data: %v", err)
		//}
		log.Printf("Log: %+v", vLog) // 处理日志
	}

}



func getEventSignatureHash(eventSignature string) common.Hash {
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(eventSignature))
	var buf []byte
	return common.BytesToHash(hash.Sum(buf))
}

func ReadKeystoreFile(file string, password string) (*ecdsa.PrivateKey,common.Address){
	// 读取keystore文件
	keyJson, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("failed to read keystore file: %v", err)
	}

	// 使用go-ethereum库中的keystore解密
	key, err := keystore.DecryptKey(keyJson, password)
	if err != nil {
		log.Fatalf("failed to decrypt key: %v", err)
	}

	return key.PrivateKey,key.Address
}

func TransPrivateKeyToString(key *ecdsa.PrivateKey) string {
	return hex.EncodeToString(crypto.FromECDSA(key))
}

func ApproveToTheContract(addr string,nodeType uint64)  {
	client, err := ethclient.Dial(Url)
	if err != nil {
		errStr := fmt.Sprintf("ethclient dial failed:%s",err.Error())
		log.Fatal(errStr)
	}
	// 创建一个上下文，用于取消交易
	ctx := context.Background()

	private,_ := PrivateKeyToAddress(PrivateKey)
	// 通过合约地址实例化合约
	contractAddress := common.HexToAddress(TokenContractAddress)
	instance,err := NewToken(contractAddress,client)
	if err != nil {
		errStr := fmt.Sprintf("NewToken new failed:%s",err.Error())
		log.Fatal(errStr)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(private, big.NewInt(ChainID)) // For Mainnet
	if err != nil {
		log.Fatal(err)
	}
	approveToAddr := common.HexToAddress(addr)
	tx, err := instance.Approve(auth,approveToAddr,new(big.Int).SetUint64(ApproveValue))
	if err != nil {
		errStr := fmt.Sprintf("domic node NewToken new failed:%s",err.Error())
		log.Fatal(errStr)
	}
	// 等待交易被打包并获取交易哈希
	fmt.Println("交易哈希:", tx.Hash().Hex())
	// 等待交易被确认
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatal(err)
	}
	// 检查交易状态
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("交易失败")
	}
	fmt.Println("交易成功!")
}

func MintToken()  {
	client, err := ethclient.Dial(Url)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个上下文，用于取消交易
	ctx := context.Background()
	private,addr := PrivateKeyToAddress(PrivateKey)
	// 通过合约地址实例化合约
	contractAddress := common.HexToAddress(TokenContractAddress)
	instance,err := NewToken(contractAddress,client)
	if err != nil {
		errStr := fmt.Sprintf("cant create contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(private, big.NewInt(ChainID)) // For Mainnet
	if err != nil {
		log.Fatal(err)
	}
	tx,err := instance.Mint(auth,addr,new(big.Int).SetUint64(ApproveValue))
	if err != nil {
		errStr := fmt.Sprintf("cant Mint by contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	// 等待交易被打包并获取交易哈希
	fmt.Println("交易哈希:", tx.Hash().Hex())
	// 等待交易被确认
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		errStr := fmt.Sprintf("cant WaitMined by contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	// 检查交易状态
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("交易失败")
	}
	fmt.Println("交易成功!")
}

func SendCommitToL1(index uint64,length uint64,dasKey [32]byte,sign [][]byte,commit string)  {
	client, err := ethclient.Dial(Url)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个上下文，用于取消交易
	ctx := context.Background()
	private,_ := PrivateKeyToAddress(PrivateKey)
	// 通过合约地址实例化合约
	contractAddress := common.HexToAddress(DomiconCommitmentContractAddress)
	instance,err := NewDomicCommit(contractAddress,client)
	if err != nil {
		errStr := fmt.Sprintf("cant create contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(private, big.NewInt(ChainID)) // For Mainnet
	if err != nil {
		log.Fatal(err)
	}
	// dasKey  签名所在组织的hash key
	var digest kzg.Digest
	str, _ := hex.DecodeString(commit)
	digest.SetBytes(str)
	commitData := PairingG1Point{
		X: new(big.Int).SetBytes(digest.X.Marshal()),
		Y: new(big.Int).SetBytes(digest.Y.Marshal()),
	}
	tx,err := instance.SubmitCommitment(auth,index,length,dasKey,sign,commitData)
	if err != nil {
		log.Fatal(err)
	}
	// 等待交易被打包并获取交易哈希
	fmt.Println("交易哈希:", tx.Hash().Hex())
	// 等待交易被确认
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		errStr := fmt.Sprintf("cant WaitMined by contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	// 检查交易状态SignData
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("交易失败")
	}
	fmt.Println("交易成功!")
}

func GenerateDAAndCommit(data []byte) string {
	currentPath, _ := os.Getwd()
	path := currentPath + "/srs"
	domiconSDK,err := kzgSdk.InitDomiconSdk(dSrsSize,path)
	if err != nil {
		println("kzg init domicon sdk err",err.Error())
	}
	digst,err := domiconSDK.GenerateDataCommit(data)
	if err != nil {
		println("GenerateDataCommit ---ERR",err.Error())
	}
	digstData := digst.Marshal()
	return common.Bytes2Hex(digstData)
}

func GetL1LatestBlockNum() uint64 {
	client, err := ethclient.Dial(Url)
	if err != nil {
		log.Fatal(err)
	}

	num,err := client.BlockNumber(context.Background())
	if err != nil {

	}
	return num
}


func RegisterBroadCastNode()  {
	client, err := ethclient.Dial(Url)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个上下文，用于取消交易
	ctx := context.Background()
	private,addr := ReadKeystoreFile(BroadCastKeyStorePth,PWD)
	contractAddress := common.HexToAddress(DomiconNodeContractAddress)
	instance,err := NewDomicNode(contractAddress,client)
	if err != nil {
		errStr := fmt.Sprintf("cant create contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(private, big.NewInt(ChainID)) // For Mainnet
	if err != nil {
		log.Fatal(err)
	}

	nodeInfo := DomiconNodeNodeInfo{
		Url: BroadCastURL,
		Name: "BroadCastNode",
		StakedTokens: new(big.Int).SetUint64(1000),
		Location: "-",
		MaxStorageSpace: new(big.Int).SetUint64(100),
		Addr: addr,
	}

	tx,err := instance.RegisterBroadcastNode(auth,nodeInfo)
	if err != nil {
		log.Fatal(err)
	}
	// 等待交易被打包并获取交易哈希
	fmt.Println("交易哈希:", tx.Hash().Hex())
	// 等待交易被确认
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		errStr := fmt.Sprintf("cant WaitMined by contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	// 检查交易状态
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("交易失败")
	}
	fmt.Println("交易成功!")
}

func RegistStorageNode()  {
	client, err := ethclient.Dial(Url)
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个上下文，用于取消交易
	ctx := context.Background()
	private,addr := ReadKeystoreFile(StorageKeyStorePth,PWD)
	contractAddress := common.HexToAddress(DomiconNodeContractAddress)
	instance,err := NewDomicNode(contractAddress,client)
	if err != nil {
		errStr := fmt.Sprintf("cant create contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(private, big.NewInt(ChainID)) // For Mainnet
	if err != nil {
		log.Fatal(err)
	}
	nodeInfo := DomiconNodeNodeInfo{
		Url: StorageURL,
		Name: "StorageNode",
		StakedTokens: new(big.Int).SetUint64(1000),
		Location: "-",
		MaxStorageSpace: new(big.Int).SetUint64(100),
		Addr: addr,
	}

	tx,err := instance.RegisterStorageNodeList(auth,nodeInfo)
	if err != nil {
		log.Fatal(err)
	}
	// 等待交易被打包并获取交易哈希
	fmt.Println("交易哈希:", tx.Hash().Hex())
	// 等待交易被确认
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		errStr := fmt.Sprintf("cant WaitMined by contract address err:%s",err.Error())
		log.Fatal(errStr)
	}
	// 检查交易状态
	if receipt.Status == types.ReceiptStatusFailed {
		log.Fatal("交易失败")
	}
	fmt.Println("交易成功!")
}