CONTRACT ?= SavingsDai

compile :; cd pkg/contracts && solc --bin --abi ${CONTRACT}.sol -o build
abigen :; abigen --bin=pkg/contracts/build/${CONTRACT}.bin --abi=pkg/contracts/build/${CONTRACT}.abi --pkg=contracts --out=pkg/contracts/${CONTRACT}.go
