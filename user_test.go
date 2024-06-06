package main

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"testing"
)


func TestPrivateKeyToAddress(t *testing.T) {
	_,addr := PrivateKeyToAddress(PrivateKey)
	println("addr",addr.Hex())
}

func TestGenerateDAAndCommit(t *testing.T) {
	index := 0
	s := strconv.Itoa(index)
	data := bytes.Repeat([]byte(s), 1024)
	str := GenerateDAAndCommit(data)
	println("commit str",str)
}

func TestGetL1LatestBlockNum(t *testing.T) {
	num := GetL1LatestBlockNum()
	println("current block num",num)
}

func TestReadKeystoreFile(t *testing.T) {
	 priv,addr := ReadKeystoreFile("./keystoreFile","123456")
	 privStr := TransPrivateKeyToString(priv)
	 println("addr -----",addr.String())
	 println("privStr-----",privStr)
}

func TestReadKeystoreFile1(t *testing.T) {
	priv,addr := ReadKeystoreFile("./keystoreFileS","123456")
	privStr := TransPrivateKeyToString(priv)
	println("addr -----",addr.String())
	println("privStr-----",privStr)
}

func TestApproveToTheContract(t *testing.T) {

	ApproveToTheContract("0xEA2bA414b6BD12c682aDA853b36e4C3C7483ed5f",0)
}

func TestSendCommitToL1(t *testing.T) {
	daskey := common.HexToHash("bd5064c5be5c91b2c22c616f33d66f6c0f83b93e8c4748d8dfaf37cb9f00d622")
	sgin := common.Hex2Bytes("9d48eb0901fec77298bb5508a4fb9214a2ee6a0d1c06650d909848c7b965af213f79e80d5479ec87e5853ded01893cff9999aab2231acb6cabfa593ff91889561c")
      commit := "13203ac5d527e264a3883339079297b4353a3f43418bbd7f1486dd5a4d6dcfc710fa963ab7fef1ff62c37d0ae3ff2b74b095c33dce34243aae6bd71bcdcff5f2"
	SendCommitToL1(8,1024,daskey,[][]byte{sgin},commit)
}

func TestFilterQuery(t *testing.T) {

	FilterQuery()
}

func TestEventSignatureHash(t *testing.T)  {
	eventName := "SendDACommitment(address,(uint256,uint256),uint256,uint256,uint256,uint256,bytes32,bytes32,bytes[])"
	eventHash := getEventSignatureHash(eventName)
	fmt.Printf("Event Signature Hash: %s\n", eventHash.Hex())
}
