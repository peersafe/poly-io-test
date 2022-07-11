package chainsql_deployer

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
	eccmAddr,err = invoker.DeployCrossChainManagerContract(eccdAddr,config.DefConfig.ChainsqlChainID)
	if err != nil{
		panic(err)
	}
	log.Infof("eccd_address:%s",eccdAddr)
	log.Infof("eccm_address:%s",eccmAddr)
}
