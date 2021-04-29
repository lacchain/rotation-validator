package main

import(
	"os"
    "github.com/spf13/viper"
	"github.com/lacchain/rotation-validator/client/model"
	"github.com/lacchain/rotation-validator/client/service"
	"github.com/lacchain/rotation-validator/client/audit"
)

var config *model.Config
var rotationService *service.RotationService

func main(){
    config = getConfigFromFile()

	rotationService = new(service.RotationService)
	err := rotationService.Init(config)
	if err != nil{
		audit.GeneralLogger.Fatal(err)
		return
	}

	rotationService.ProcessEvents()

}

func getConfigFromFile()(*model.Config){
	v := viper.New()
	v.SetConfigName("config")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		audit.GeneralLogger.Printf("couldn't load config: %s", err)
		os.Exit(1)
	}
	var c model.Config
	if err := v.Unmarshal(&c); err != nil {
		audit.GeneralLogger.Printf("couldn't read config: %s", err)
		os.Exit(1)
	}
	audit.GeneralLogger.Printf("smartContract=%s\n", c.Application.ContractAddress)
	return &c
}