/*
Copyright © 2019 Adrian Pareja <adriancc5.5@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"os"
	"math/big"
	"sync"
	"fmt"
	"github.com/spf13/viper"
	"github.com/lacchain/rotation-validator/server/model"
	"github.com/lacchain/rotation-validator/server/service"
	log "github.com/lacchain/rotation-validator/server/audit"
)


var config *model.Config

var blockNumber int64 
var mappingSender map[string]int

var lock sync.Mutex
var rotationService *service.RotationService

func main() {
	config = getConfigFromFile()

	rotationService = new(service.RotationService)
	err := rotationService.Init(config)
	if err != nil{
		log.GeneralLogger.Fatal(err)
		return
	}

	header, err := rotationService.GetHeaderBlock()
	if err != nil {
		handleError(err)
	}

	blockNumber = header.Int64()

	fmt.Println("header:",blockNumber)

	finish := make(chan interface{})

	processBlocks(finish)

	<-finish
}

func getConfigFromFile()(*model.Config){
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		log.GeneralLogger.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	var c model.Config
	if err := v.Unmarshal(&c); err != nil {
		log.GeneralLogger.Printf("couldn't read config: %s", err)
		os.Exit(1)
	}
	log.GeneralLogger.Printf("nodeURL=%s", c.Application.WSURL)
	return &c
}

func processBlocks(finish chan<- interface{}){
	go func() {
		for {
			existBlock, _ := rotationService.GetBlockNumber(big.NewInt(blockNumber))
			if existBlock {
				blockNumber++
				if (blockNumber % 30000 == 0){
					rotationService.ExecuteRotation()
				}
			}
		}
		close(finish)
	}()


}

func handleError(err error){
	log.GeneralLogger.Fatal(err.Error())	
}
