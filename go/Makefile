CONTRACT ?= SavingsDai

compile :; cd contracts && solc --bin --abi ${CONTRACT}.sol -o build
abigen :; abigen --bin=contracts/build/${CONTRACT}.bin --abi=contracts/build/${CONTRACT}.abi --pkg=contracts --out=contracts/${CONTRACT}.go
