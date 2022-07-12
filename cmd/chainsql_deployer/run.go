package main

import (
	"flag"
	"github.com/polynetwork/poly-io-test/chains/chainsql"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
)

var (
	fnEth        string
	configPath   string
)

func init() {
	flag.StringVar(&configPath, "conf", "./config.json", "Config of poly-io-test")
	flag.StringVar(&fnEth, "func", "deploy", "choose function to run: deploy or setup")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(configPath)
	if err != nil {
		panic(err)
	}

	switch fnEth {
	case "deploy":
		DeployChainsqlSmartContract()
	case "setup":
		break
	}
}

func DeployChainsqlSmartContract() {

	var (
		eccdAddr  string
		eccmAddr  string
		err       error
	)

	invoker, err := chainsql.NewChainsqlInvoker()
	if err != nil {
		panic(err)
	}

	eccdAddr,err = invoker.DeployCrossChainDataContract()
	if err != nil{
		panic(err)
	}

	log.Infof("eccd_address:%s",eccdAddr)

	eccmAddr,err = invoker.DeployCrossChainManagerContract(eccdAddr,config.DefConfig.ChainsqlChainID)

	result,err := invoker.TransaferOwnershipForECCD(eccdAddr,eccmAddr)

	if err != nil{
		panic(err)
	}
	if result.Status != "validate_success"{
		panic(result.ErrorMessage)
	}

	log.Infof("eccm_address:%s",eccmAddr)
	eccmpAddr,err := invoker.DeployCrossChainManagerProxyContract(eccmAddr)
	if err != nil{
		panic(err)
	}
	result,err = invoker.TransferOwnershipForECCM(eccmAddr,eccmpAddr)
	if result.Status != "validate_success"{
		panic(result.ErrorMessage)
	}

	log.Infof("eccmp_address:%s",eccmpAddr)
}
