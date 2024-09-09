ANVIL_PORT = 8545
ANVIL_PID_FILE = /tmp/anvil.pid

optimism_artifacts := $(wildcard pkg/sdk/optimism/abis/*.json)
ethereum_artifacts := $(wildcard pkg/sdk/ethereum/abis/*.json)

$(optimism_artifacts:.json=.go): %.go : %.json
	abigen --abi=$< --pkg=optimismContracts --type=$(basename $(notdir $<)) --out=$@

$(ethereum_artifacts:.json=.go): %.go : %.json
	abigen --abi=$< --pkg=ethereumContracts --type=$(basename $(notdir $<)) --out=$@

build: $(optimism_artifacts:.json=.go) $(ethereum_artifacts:.json=.go)

compile: clean build

clean:
	rm -f pkg/sdk/optimism/abis/*.go
	rm -f pkg/sdk/ethereum/abis/*.go

test: run_tests.sh
	@./run_tests.sh
