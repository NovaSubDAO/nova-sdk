#!/bin/bash

ANVIL_PORT=8545
ANVIL_PID_FILE=/tmp/anvil.pid

echo "Starting Anvil..."

anvil --fork-url "$OPT_RPC_ENDPOINT" --port "$ANVIL_PORT" & echo $! > "$ANVIL_PID_FILE"

cleanup() {
    echo "Stopping Anvil..."
    if [ -f "$ANVIL_PID_FILE" ]; then
        kill $(cat "$ANVIL_PID_FILE")
        rm -f "$ANVIL_PID_FILE"
    fi
}

trap 'cleanup' EXIT

sleep 5

echo "Setting balance for test address..."
curl -s http://localhost:"$ANVIL_PORT" -X POST -H "Content-Type: application/json" \
--data '{"method":"anvil_setBalance","params":["'"$TEST_ADDRESS"'","0x56BC75E2D63100000"],"id":1,"jsonrpc":"2.0"}'

echo "Running tests..."
OPT_RPC_ENDPOINT=http://localhost:"$ANVIL_PORT" go test github.com/NovaSubDAO/nova-sdk/go/...
