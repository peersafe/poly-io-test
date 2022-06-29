         GOFMT=gofmt
GC=go build
PWD := $(shell pwd)

ARCH=$(shell uname -m)
SRC_FILES = $(shell git ls-files | grep -e .go$ | grep -v _test.go)

cct: $(SRC_FILES)
	CGO_ENABLED=1 $(GC) -o cct  cmd/cctest/main.go

cct-windows:
	GOOS=windows GOARCH=amd64 $(GC) -o cct-windows-amd64.exe cmd/cctest/main.go
cct-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GC) -o cct-linux-amd64 cmd/cctest/main.go
cct-mac:
	GOOS=darwin GOARCH=amd64 $(GC)  -o ccct-darwin-amd64 cmd/cctest/main.go
cct-btc-prepare:
	GOOS=linux GOARCH=amd64 $(GC) -o cct-btc-linux-amd64 cmd/btc_prepare/run.go
cct-eth-deployer:
	GOOS=linux GOARCH=amd64 $(GC)  -o cct-eth-linux-amd64 cmd/eth_deployer/run.go
cct-cosmos-deployer:
	GOOS=linux GOARCH=amd64 $(GC)  -o cct-cosmos-linux-amd64 cmd/cosmos_prepare/run.go
cct-ont-deployer:
	GOOS=linux GOARCH=amd64 $(GC)  -o cct-ont-linux-amd64 cmd/ont_deployer/run.go

format:
	$(GOFMT) -w cmd/cctest/main.go

clean:
	rm -rf *.8 *.o *.out *.6 *exe coverage
	rm -rf cct-*


runGetSideChain:

addRelayer:
	go run cmd/tools/run.go -tool add_relayer -newwallet ../poly/build/wallet/wallet2.dat -newpwd 123 -pwallets ../poly/build/wallet/wallet1.dat,../poly/build/wallet/wallet2.dat,../poly/build/wallet/wallet3.dat,../poly/build/wallet/wallet4.dat -ppwds 123,123,123,123 -chainid 7

registerSideChain:
	go run cmd/tools/run.go -tool register_side_chain -pwallets ../poly/build/wallet/wallet1.dat,../poly/build/wallet/wallet2.dat,../poly/build/wallet/wallet3.dat,../poly/build/wallet/wallet4.dat -ppwds 123,123,123,123 -chainid 7

syncGenesisHeader:
	go run cmd/tools/run.go -tool sync_genesis_header -pwallets ../poly/build/wallet/wallet1.dat,../poly/build/wallet/wallet2.dat,../poly/build/wallet/wallet3.dat,../poly/build/wallet/wallet4.dat -ppwds 123,123,123,123 -chainid 7

syncFabricRootCa:
	go run cmd/tools/run.go -tool sync_fabric_root_ca -pwallets ../poly/build/wallet/wallet1.dat,../poly/build/wallet/wallet2.dat,../poly/build/wallet/wallet3.dat,../poly/build/wallet/wallet4.dat -ppwds 123,123,123,123 -chainid 7 -rootca /Volumes/data/project/golang/blockchain/polynetwork/fabric-relayer/build/node2/peerOrganizations/org2.example.com/ca/ca.org2.example.com-cert.pem