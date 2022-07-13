# 一、编译 chainsql_depolyer

## 1、准备

```bash
> go mod tidy
```

## 2、编译

```bash
> go build -o build/chainsql_depolyer cmd/chainsql_deployer/run.go
```

## 3、配置

```bash
> cd build
# create chainsql.json
> cat > chainsql.json <<EOF
{
    "URL": "ws://127.0.0.1:6006",
    "ServerName": "",
    "RootCertPath":"",
    "ClientCertPath":"",
    "ClientKeyPath":"",
    "Account": {
        "Address":"zHb9CJAWyB4zj91VRWn96DkukG4bwdtyTh",
        "Secrect":"xnoPBzXtMeMyMHUVTgbuqAfg1SUTb"
    }
}
EOF

# create config.json
# poly_wallet.dat 为 poly 中继链的账户钱包
> cat > config.json <<EOF
{
    "RCWallet": "./poly_wallet.dat",
    "RCWalletPwd": "peersafe",
    "RchainJsonRpcAddress": "http://127.0.0.1:21336",
    "RCEpoch": 0,
    "ReportInterval": 60,
    "ReportDir": "./report",
    "BatchTxNum": 1,
    "BatchInterval": 1,
    "TxNumPerBatch": 1,
    "ChainsqlSdkConfFile":"./chainsql.json",
    "ChainsqlChainID":2000
}
EOF
```

## 4、 部署 eccd/eccm/eccmp 合约

```bash
> ./chainsql_deployer -conf=config.json
2022/07/13 11:51:28.392053 [INFO ] GID 1, eccd_address:zEC26Fq4znnTGrWf9NcXrMT3fP1zdU14fK
2022/07/13 11:51:30.010944 [INFO ] GID 1, eccm_address:zLgs3EUNxV7RySkyiw4R9Z1jREEumWrzKk
2022/07/13 11:51:32.061788 [INFO ] GID 1, eccmp_address:zPNpr5mXvZcMN39T1r5K9vcqstegVANyew
```

# 二、编译 poly-tools 工具

## 1、编译

```bash
> go build -o build/poly-tools cmd/tools/run.go
```

## 2、注册侧链

```bash
./poly-tools --conf=./config.json -tool register_side_chain -pwallets ./poly/node1/wallet/wallet.dat,./poly/node2/wallet/wallet.dat,./poly/node3/wallet/wallet.dat,./poly/node4/wallet/wallet.dat -ppwds peersafe,peersafe,peersafe,peersafe --chainid 2000
```

## 3、同步 CA 证书

```bash
./poly-tools --conf=./config.json -tool sync_chainsql_root_ca -pwallets ./poly/node1/wallet/wallet.dat,./poly/node2/wallet/wallet.dat,./poly/node3/wallet/wallet.dat,./poly/node4/wallet/wallet.dat -ppwds peersafe,peersafe,peersafe,peersafe --chainid 2000 -rootca ./certs/rootCA.crt
```

## 4、同步侧链区块

```bash
./poly-tools --conf=./config.json --tool sync_genesis_header -pwallets ./poly/node1/wallet/wallet.dat,./poly/node2/wallet/wallet.dat,./poly/node3/wallet/wallet.dat,./poly/node4/wallet/wallet.dat -ppwds peersafe,peersafe,peersafe,peersafe --chainid 2000
```

