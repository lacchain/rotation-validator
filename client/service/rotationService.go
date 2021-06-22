/*
	Rotation Service
	version 0.9
	author: Adrian Pareja Abarca
	email: adriancc5.5@gmail.com
*/

package service

import (
	"sync"
	"context"
	"strconv"
	"io/ioutil"
	"encoding/hex"
	"crypto/ecdsa"
	"github.com/lacchain/rotation-validator/client/model"
	"github.com/lacchain/rotation-validator/client/errors"
	"github.com/lacchain/rotation-validator/client/audit"
	"github.com/lacchain/rotation-validator/client/rpc"
	bl "github.com/lacchain/rotation-validator/client/blockchain"
	sha "golang.org/x/crypto/sha3"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

const rotationStartedEvent = "RotationStarted(uint256)"

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

//ProcessEvents from blockchain
func (service *RotationService) ProcessEvents(){
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.WSURL)
	if err != nil {
		audit.GeneralLogger.Fatal(err)
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
	d.Write([]byte(rotationStartedEvent))

	eventRotationStarted := hex.EncodeToString(d.Sum(nil))

    for {
        select {
        case err := <-sub.Err():
			audit.GeneralLogger.Println("WebSocket Failed")
            audit.GeneralLogger.Fatal(err)
        case vLog := <-logs:
			audit.GeneralLogger.Println("event emited:",vLog.Topics[0].Hex())
			if vLog.Topics[0].Hex() == "0x"+eventRotationStarted {
				var wg sync.WaitGroup
				results := make(chan *rpc.JsonrpcMessage)
				done := make(chan interface{})
				wg.Add(2)
				go service.removeValidators(done, &wg, results)
				go service.voteByValidators(done, &wg, results)
				go service.checkResults(results)
				wg.Wait()
				close(done)
			}
        }
    }
}

func (service *RotationService) removeValidators(done <-chan interface{}, wg *sync.WaitGroup, results chan<- *rpc.JsonrpcMessage){
	defer wg.Done()
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.RPCURL)
	if err != nil {
		audit.GeneralLogger.Fatal(err)
	}
	defer client.Close()
	oldValidators,err := client.GetOldValidators(common.HexToAddress(service.Config.Application.ContractAddress))
	if err != nil {
		HandleError(err)
	}
	for id, validator := range oldValidators{
		audit.GeneralLogger.Println("removing validator:",validator)	
		resp := vote(service.Config.Application.RPCURL, strconv.Itoa(id+10), validator, false, service.Config.Application.Timeout)
		
		select {            
			case <-done: 
				audit.GeneralLogger.Println("quit signal received...exiting from vote to removeValidators")
				return            
			case results <- resp:            
		}
	}
}

func (service *RotationService) voteByValidators(done <-chan interface{}, wg *sync.WaitGroup, results chan<- *rpc.JsonrpcMessage){
	defer wg.Done()
	client := new(bl.Client)
	err := client.Connect(service.Config.Application.RPCURL)
	if err != nil {
		audit.GeneralLogger.Fatal(err)
	}
	defer client.Close()
	newValidators,err := client.GetNewValidators(common.HexToAddress(service.Config.Application.ContractAddress))
	if err != nil {
		HandleError(err)
	}
	for id, validator := range newValidators{
		audit.GeneralLogger.Println("adding validator:",validator)
		resp := vote(service.Config.Application.RPCURL, strconv.Itoa(id+100), validator, true, service.Config.Application.Timeout)
		select {            
			case <-done: 
				audit.GeneralLogger.Println("quit signal received...exiting from voteByValidators")
				return            
			case results <- resp:            
		}
	} 
}

func (service *RotationService) checkResults(results <-chan *rpc.JsonrpcMessage){
	for result := range results{
		executed,err:=strconv.ParseBool(string(result.Result))
		if err != nil{
			HandleError(err)
			executed = false
		}
		if !executed{
			audit.GeneralLogger.Println("Error voting:",result.String())	
		}
	}
}

//HandleError ...
func HandleError(err error){
	audit.GeneralLogger.Println(err.Error())
}