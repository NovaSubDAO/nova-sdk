CONTRACT ?= SavingsDai

compile :; cd pkg/contracts && solc --bin --abi ${CONTRACT}.sol -o build --overwrite && abigen --bin=build/${CONTRACT}.bin --abi=build/${CONTRACT}.abi --pkg=contracts --out=${CONTRACT}.go
