CONTRACT ?= SavingsDai

compile :
	cd pkg/sdk/ethereum/abis && \
		solc --bin --abi ${CONTRACT}.sol \
			-o /tmp --overwrite && \
			abigen --bin=/tmp/${CONTRACT}.bin --abi=/tmp/${CONTRACT}.abi \
			--pkg=abis --out=${CONTRACT}.go
	#cd pkg/sdk/optimism/abis && \
		solc --bin --abi ${CONTRACT}.sol \
			-o build --overwrite && \
			abigen --bin=build/${CONTRACT}.bin --abi=build/${CONTRACT}.abi \
			--pkg=abis --out=${CONTRACT}.go
test :; go test github.com/NovaSubDAO/nova-sdk/go/...
