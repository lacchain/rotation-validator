package blockchain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lacchain/rotation-validator/client/errors"
	"github.com/lacchain/rotation-validator/client/audit"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	relay "github.com/lacchain/rotation-validator/client/blockchain/contracts"
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

	audit.GeneralLogger.Println("Connected to Ethereum Node:", nodeURL)
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

	oldValidators, err := contract.GetRemovedValidators(&bind.CallOpts{})
	if err != nil {
		msg := "failed get validators to remove"
		err = errors.CallBlockchainFailed.Wrapf(err, msg, -32603)
		return nil, err
	}

	audit.GeneralLogger.Println("Number of validators to remove:",len(oldValidators))

	validators := make([]string, len(oldValidators))

	for i,validator := range oldValidators{
		validators[i] = validator.Id.Hex()[2:]
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
		msg := "failed get new Validators"
		err = errors.CallBlockchainFailed.Wrapf(err, msg, -32603)
		return nil, err
	}

	audit.GeneralLogger.Println("Number of validators to add:",len(newValidators))

	validators := make([]string, len(newValidators))

	for i,validator := range newValidators{
		validators[i] = validator.Hex()[2:]
	}

	return validators, nil
}
