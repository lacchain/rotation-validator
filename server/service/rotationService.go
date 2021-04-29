/*
	Rotation Service
	version 0.9
	author: Adrian Pareja Abarca
	email: adriancc5.5@gmail.com
*/

package service

import (
	"io/ioutil"
	"math/big"
	"crypto/ecdsa"
	"github.com/lacchain/rotation-validator/server/model"
	"github.com/lacchain/rotation-validator/server/errors"
	"github.com/lacchain/rotation-validator/server/audit"
	bl "github.com/lacchain/rotation-validator/server/blockchain"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//RotationService is the main service
type RotationService struct {
	// The service's configuration
	Config *model.Config
}

//Init configuration parameters
func (service *RotationService) Init(_config *model.Config) error{
	service.Config = _config
	key, err := ioutil.ReadFile(service.Config.Application.NodeKeyPath)
    if err != nil {
		msg := "fail to read key file"
		err = errors.FailedReadFile.Wrapf(err,msg, -32600)
        return err
	}

	privateKey, err := crypto.HexToECDSA(string(key[2:66]))
    if err != nil {
		msg := "fail to convert Hex to ECDSA"
		err = errors.FailedKeystore.Wrapf(err,msg, -32601)
        return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
		msg := "error casting public key to ECDSA"
		err = errors.FailedKeystore.Wrapf(err,msg, -32601)
        return err
	}
	
	service.Config.Application.NodeAddress = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()[2:]

	return nil
}

//ExecuteRotation of validators
func (service *RotationService) ExecuteRotation() {
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.WSURL)
	if err != nil {
		HandleError(err)
	}
	defer client.Close()
	
	privateKey, err := crypto.HexToECDSA(service.Config.Application.Key)
    if err != nil {
        audit.GeneralLogger.Fatal(err)
	}
	
	options, err := client.ConfigTransaction(privateKey)
	if err != nil {
		HandleError(err)
	}
	contractAddress := common.HexToAddress(service.Config.Application.ContractAddress)

	err = client.InitRotation(contractAddress, options)
	if err != nil {
		HandleError(err)
	}
}

//GetHeaderBlock of blockchain
func (service *RotationService) GetHeaderBlock() (*big.Int,error){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.WSURL)
	if err != nil {
		HandleError(err)
	}
	defer client.Close()
	
	header, err := client.GetHeaderBlock()

	if err!=nil{
		return nil, err
	}

	return header, nil
}

//GetBlockNumber of blockchain
func (service *RotationService) GetBlockNumber(blockNumber *big.Int) (bool,*types.Block){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.WSURL)
	if err != nil {
		HandleError(err)
	}
	defer client.Close()
	
	return client.GetBlockNumber(blockNumber)
}

//HandleError ...
func HandleError(err error){
	audit.GeneralLogger.Println(err.Error())
}