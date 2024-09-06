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

test: build
	@echo "Starting Anvil..."
	@{ \
		anvil --fork-url $(OPT_RPC_ENDPOINT) --port $(ANVIL_PORT) & echo $$! > $(ANVIL_PID_FILE); \
		# Function to clean up Anvil process \
		cleanup() { \
			echo "Stopping Anvil..."; \
			if [ -f $(ANVIL_PID_FILE) ]; then \
				kill $$(cat $(ANVIL_PID_FILE)); \
				rm -f $(ANVIL_PID_FILE); \
			fi \
		}; \
		trap 'cleanup' EXIT; \
		# Wait for Anvil to start \
		sleep 5; \
		echo "Setting balance for test address..."; \
		curl -s http://localhost:$(ANVIL_PORT) -X POST -H "Content-Type: application/json" \
		--data '{"method":"anvil_setBalance","params":["$(TEST_ADDRESS)","0x56BC75E2D63100000"],"id":1,"jsonrpc":"2.0"}'; \
		echo "Running tests..."; \
		OPT_RPC_ENDPOINT=http://localhost:$(ANVIL_PORT) go test github.com/NovaSubDAO/nova-sdk/go/...; \
	}
