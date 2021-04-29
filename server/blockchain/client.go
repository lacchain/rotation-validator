package blockchain

import (
	"fmt"
	"context"
	"crypto/ecdsa"
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

//ConfigTransaction from ethereum address contract
func (ec *Client) ConfigTransaction(key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	auth := bind.NewKeyedTransactor(key)

	nonce, err := ec.client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		msg := fmt.Sprintf("can't get pending nonce for:%s", auth.From)
		err = errors.FailedConfigTransaction.Wrapf(err, msg, -32603)
		return nil, err
	}

	gasPrice, err := ec.client.SuggestGasPrice(context.Background())
	if err != nil {
		msg := "can't get gas price suggested"
		err = errors.FailedConfigTransaction.Wrapf(err, msg, -32603)
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                // in wei
	auth.GasLimit = gasPrice.Uint64() * 2 // in units
	auth.GasPrice = gasPrice

	audit.GeneralLogger.Printf("OptionsTransaction=[From:0x%x,nonce:%d,gasPrice:%s,gasLimit:%d", auth.From, nonce, gasPrice, auth.GasLimit)

	return auth, nil
}

//InitRotation
func (ec *Client) InitRotation(contractAddress common.Address, options *bind.TransactOpts)(error){
	contract, err := relay.NewContract(contractAddress, ec.client)
	if err != nil {
		msg := fmt.Sprintf("can't instance Rotation contract %s", contractAddress)
		err = errors.FailedContract.Wrapf(err, msg, -32603)
		return err
	}

	audit.GeneralLogger.Println("Rotation Contract instanced:", contractAddress.Hex())

	transaction, err := contract.ExecuteRotation(options)
	if err != nil {
		msg := "failed execute rotation"
		err = errors.CallBlockchainFailed.Wrapf(err, msg, -32603)
		return err
	}

	if len(transaction.Hash()) == 0 {
		msg := "failed executing transaction relay Metatransaction"
		err = errors.FailedTransaction.Wrapf(err, msg, -32603)
		return err
	}
	audit.GeneralLogger.Printf("Transaction sent: %s", transaction.Hash().Hex())

	return nil
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
