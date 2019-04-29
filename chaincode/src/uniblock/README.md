# UniBlock Chaincode

## Dependencies

### Using `go mod`

Run the following `go mod` commands in this directory:

```sh
go mod download
go mod vendor
```

`go mod download` will download all modules listed in `go.mod` to the local cache.
`go mod vendor` will make vendored copy of the dependencies. This is use for vendoring the dependencies into the chaincode binary for installation of the chaincode on the Hyperledger Fabric network.


### Not using `go mod`

To install additional dependencies in your chaincode use [govendor](https://github.com/kardianos/govendor).

Run the following `govendor fetch` commands in this directory:

```sh
govendor fetch github.com/thanakritlee/dockerised-fabric-app/chaincode/src/uniblock/controllers
govendor fetch github.com/thanakritlee/dockerised-fabric-app/chaincode/src/uniblock/utils
```