package blockchain

import (
	"fmt"
	"context"
	"math/big"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lacchain/rotation-validator/server/errors"
	"github.com/lacchain/rotation-validator/server/audit"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	relay "github.com/lacchain/rotation-validator/server/blockchain/contracts"
)

//Client to manage Connection to Ethereum
type Client struct {
	client *ethclient.Client
}

//Connect to Ethereum
func (ec *Client) Connect(nodeURL string) error {
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		msg := fmt.Sprintf("Can't connect to node %s", nodeURL)
		err = errors.FailedConnection.Wrapf(err,msg)
		return err
	}

	fmt.Sprintf("Connected to Ethereum Node:", nodeURL)
	ec.client = client
	return nil
}

//Close ethereum connection
func (ec *Client) Close() {
	ec.client.Close()
}

//GetEthclient ...
func (ec *Client) GetEthclient()(*ethclient.Client) {
	return ec.client
}

//GetOldValidators
func (ec *Client) GetOldValidators(contractAddress common.Address)([]string, error){
	contract, err := relay.NewRelay(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance Rotation contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err, msg, -32603)
		return nil, err
	}

	audit.GeneralLogger.Println("Rotation Contract instanced:", contractAddress.Hex())

	oldValidators, err := contract.GetOldValidators(&bind.CallOpts{})
	if err != nil {
		msg := "failed get old Validators"
		err = errors.CallBlockchainFailed.Wrapf(err, msg, -32603)
		return nil, err
	}

	fmt.Println("oldValidators:",len(oldValidators))

	validators := make([]string, len(oldValidators))

	for i,validator := range oldValidators{
		validators[i] = validator.Hex()[2:]
	}

	return validators, nil
}

//GetOldValidators
func (ec *Client) GetNewValidators(contractAddress common.Address)([]string, error){
	contract, err := relay.NewRelay(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance Rotation contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err, msg, -32603)
		return nil, err
	}

	audit.GeneralLogger.Println("Rotation Contract instanced:", contractAddress.Hex())

	newValidators, err := contract.GetNewValidators(&bind.CallOpts{})
	if err != nil {
		msg := "failed get old Validators"
		err = errors.CallBlockchainFailed.Wrapf(err, msg, -32603)
		return nil, err
	}

	fmt.Println("newValidators:",len(newValidators))

	validators := make([]string, len(newValidators))

	for i,validator := range newValidators{
		validators[i] = validator.Hex()[2:]
	}

	return validators, nil
}

//GetBlockNumber ...
func (ec *Client) GetBlockNumber(blockNumber *big.Int) (bool, *types.Block){
    block, err := ec.client.BlockByNumber(context.Background(), blockNumber)
    if err != nil {
        audit.GeneralLogger.Println(err)
    }

	if block != nil{
    	fmt.Println(block.Number().Uint64())     // 5671744
    	fmt.Println(block.Time())       // 1527211625
    	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
    	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
    	fmt.Println(len(block.Transactions()))   // 144

		count, err := ec.client.TransactionCount(context.Background(), block.Hash())
    	if err != nil {
        	audit.GeneralLogger.Println(err)
    	}

    	fmt.Println(count) // 144
		return true, block
	}

	return false,nil  
}

//GetHeaderBlock
func (ec *Client) GetHeaderBlock() (*big.Int,error){
	header, err := ec.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		audit.GeneralLogger.Fatal(err)
	}

	return header.Number,nil
}
