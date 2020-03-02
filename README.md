# KICKTOKEN BALANCES
This script generates 2 csv files: all holders liquid and frozen balances and total liquid and frozen tokens amount

### Install
```
go mod download
```

### Local run
```
go build
./kicktoken_viewer -l ETHEREUM_NODE_URL
```

Change ETHEREUM_NODE_URL on your ethereum node url like this: `ws://128.0.0.1:8546` or `https://mainnet.infura.io/v3/PROJECT_ID`