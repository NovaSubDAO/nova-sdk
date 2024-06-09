CONTRACT ?= SavingsDai
optimism_artifacts := $(wildcard pkg/sdk/optimism/abis/*.json)
ethereum_artifacts := $(wildcard pkg/sdk/ethereum/abis/*.json)

$(optimism_artifacts:.json=.go): %.go : %.json
	abigen --abi=$< --pkg=optimism$(basename $(notdir $<)) --out=$@

$(ethereum_artifacts:.json=.go): %.go : %.json
	abigen --abi=$< --pkg=ethereum$(basename $(notdir $<)) --out=$@

build: $(optimism_artifacts:.json=.go) $(ethereum_artifacts:.json=.go)

compile: clean build

clean:
	rm -f pkg/sdk/optimism/abis/*.go
	rm -f pkg/sdk/ethereum/abis/*.go

test :; go test github.com/NovaSubDAO/nova-sdk/go/...
