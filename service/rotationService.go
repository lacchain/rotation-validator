/*
	Rotation Service
	version 0.9
	author: Adrian Pareja Abarca
	email: adriancc5.5@gmail.com
*/

package service

import (
	"fmt"
	"sync"
	"context"
	"strconv"
	"io/ioutil"
	"encoding/hex"
	"crypto/ecdsa"
	"github.com/lacchain/rotation-validators/model"
	"github.com/lacchain/rotation-validators/errors"
	"github.com/lacchain/rotation-validators/audit"
	"github.com/lacchain/rotation-validators/rpc"
	bl "github.com/lacchain/rotation-validators/blockchain"
	sha "golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

//RelaySignerService is the main service
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
		err = errors.FailedReadFile.Wrapf(err,msg, -32602)
        return err
	}

	privateKey, err := crypto.HexToECDSA(string(key[2:66]))
    if err != nil {
        audit.GeneralLogger.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        audit.GeneralLogger.Fatal("error casting public key to ECDSA")
	}
	
	service.Config.Application.NodeAddress = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()[2:]

	return nil
}

//ProcessEvents from blockchain
func (service *RotationService) ProcessEvents(){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(err)
	}
	defer client.Close()
	contractAddress := common.HexToAddress(service.Config.Application.ContractAddress)
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := client.GetEthclient().SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        audit.GeneralLogger.Fatal(err)
    }

	d := sha.NewLegacyKeccak256()
	d.Write([]byte("RotationStarted(uint256)"))

	eventRotationStarted := hex.EncodeToString(d.Sum(nil))

    for {
        select {
        case err := <-sub.Err():
            audit.GeneralLogger.Fatal(err)
        case vLog := <-logs:
            audit.GeneralLogger.Println(vLog)
			audit.GeneralLogger.Println("event:",vLog.Topics[0].Hex())
			if vLog.Topics[0].Hex() == "0x"+eventRotationStarted {
				var wg sync.WaitGroup
				results := make(chan *rpc.JsonrpcMessage)
				done := make(chan interface{})
				defer close(done)
				wg.Add(2)
				go service.removeValidators(done, &wg, results)
				go service.voteByValidators(done, &wg, results)
				wg.Wait()
			}
        }
    }
}

func (service *RotationService) removeValidators(done <-chan interface{}, wg *sync.WaitGroup, results chan<- *rpc.JsonrpcMessage){
	defer wg.Done()
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(err)
	}
	defer client.Close()
	oldValidators,err := client.GetOldValidators(common.HexToAddress(service.Config.Application.ContractAddress))
	if err != nil {
		HandleError(err)
	}
	for id, validator := range oldValidators{
		fmt.Println("index:",id+10)
		resp := vote(strconv.Itoa(id+1), validator, false)
		
		select {            
			case <-done: 
				return            
			case results <- resp:            
		}
	}
}

func (service *RotationService) voteByValidators(done <-chan interface{}, wg *sync.WaitGroup, results chan<- *rpc.JsonrpcMessage){
	defer wg.Done()
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.NodeURL)
	if err != nil {
		HandleError(err)
	}
	defer client.Close()
	newValidators,err := client.GetNewValidators(common.HexToAddress(service.Config.Application.ContractAddress))
	if err != nil {
		HandleError(err)
	}
	for id, validator := range newValidators{

		resp := vote(string(id+100), validator, true)
		
		select {            
			case <-done: 
				return            
			case results <- resp:            
		}
	} 
}

//HandleError ...
func HandleError(err error){
	audit.GeneralLogger.Println(err.Error())
}