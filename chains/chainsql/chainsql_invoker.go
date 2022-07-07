package chainsql

import (
	"encoding/json"
	"fmt"
	"github.com/ChainSQL/go-chainsql-api/core"
	config2 "github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"io/ioutil"
	"os"
)
type ChainsqlInvoker struct {
	ChainsqlSdk *core.Chainsql
	TransOpts *core.TransactOpts
}

type Account struct {
	Address string
	Secrect string
}
type Config struct {
	URL            string
	ServerName     string
	RootCertPath   string
	ClientCertPath string
	ClientKeyPath  string
	Account        Account
}


func NewConfig(configFilePath string) *Config {

	fileContent, err := ReadFile(configFilePath)
	if err != nil {
		log.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	config := &Config{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		log.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}

	return config
}

// Dial connects a client to the given URL and groupID.
func Dial(config *Config) (*core.Chainsql, error) {
	node := core.NewChainsql()
	node.Connect(
		config.URL,
		config.RootCertPath,
		config.ClientCertPath,
		config.ClientKeyPath,
		config.ServerName)

	node.As(config.Account.Address, config.Account.Secrect)
	return node, nil
}

func NewChainsqlInvoker() (*ChainsqlInvoker, error) {
	instance := &ChainsqlInvoker{}
	cfg := NewConfig(config2.DefConfig.ChainsqlSdkConfFile)

	chainsql, err := Dial(cfg)
	if err != nil {
		return nil, err
	}
	instance.ChainsqlSdk = chainsql
	instance.TransOpts = &core.TransactOpts{
		ContractValue: 0,
		Gas:           30000000,
		Expectation:   "validate_success",
	}
	return instance, nil
}


func ReadFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: open file %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Errorf("ReadFile: File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}