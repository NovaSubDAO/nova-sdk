package sdk

import (
    "log"
    "math/big"
    "context"
    "fmt"
    "os"
    "strings"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/NovaSubDAO/nova-sdk/go/pkg/config"
	"github.com/NovaSubDAO/nova-sdk/go/pkg/contracts"
)

type NovaSDK struct {
    Contract *contracts.ContractsCaller
    Client   *ethclient.Client
}

func NewNovaSDK(contractAddress string) (*NovaSDK, func()) {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
    }

    // Initialize the Ethereum client
    client, err := ethclient.Dial(cfg.RpcEndpoint)
    if err != nil {
        log.Fatalf("Failed to connect to Ethereum client: %v", err)
    }

    contract, err := contracts.NewContractsCaller(common.HexToAddress(contractAddress), client)
    if err != nil {
        log.Fatalf("Failed to instantiate contract caller: %v", err)
    }

    return &NovaSDK{
        Contract: contract,
        Client:   client,
    }, func() {
        client.Close()
    }
}

func (sdk *NovaSDK) GetPrice() (*big.Int, error) {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading configuration:", err)
    }

    factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(cfg.VaultDecimals)), nil)

    assets, err := sdk.Contract.ConvertToAssets(nil, factor)
    if err != nil {
        return big.NewInt(0), err
    }

    return assets, nil
}

func (sdk *NovaSDK) GetPosition(address common.Address) (*big.Int, error) {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading configuration:", err)
    }

    balance, err := sdk.Contract.BalanceOf(nil, address)
    if err != nil {
        return big.NewInt(0), err
    }

    price, err := sdk.GetPrice()
    if err != nil {
        return big.NewInt(0), err
    }

    value := new(big.Int).Mul(balance, price)
    factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(cfg.VaultDecimals)), nil)
    valueNormalized := new(big.Int).Div(value, factor)

    return valueNormalized, nil
}

func (sdk *NovaSDK) GetTotalValue() (*big.Int, error) {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Error loading configuration:", err)
    }

    totalSupply, err := sdk.Contract.TotalSupply(nil)
    if err != nil {
        return big.NewInt(0), err
    }

    price, err := sdk.GetPrice()
    if err != nil {
        return big.NewInt(0), err
    }

    totalValue := new(big.Int).Mul(totalSupply, price)
    factor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(cfg.VaultDecimals)), nil)
    totalValueNormalized := new(big.Int).Div(totalValue, factor)

    return totalValueNormalized, nil
}

func (sdk *NovaSDK) GetSlippage(amount *big.Int) (*big.Int, error) {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
    }

    client, err := ethclient.Dial(cfg.RpcEndpoint)
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress("0xa4ac92a0F54f1a447c55a4082c90742F5E76Df62")

    abiPath := "abis/MixedRouteQuoterV1.json"
    file, err := os.ReadFile(abiPath)
    if err != nil {
        log.Fatalf("Failed to read ABI file: %v", err)
    }

    parsedABI, err := abi.JSON(strings.NewReader(string(file)))
    if err != nil {
        log.Fatal(err)
    }

    // Define the parameters for the function call
    tokenIn := common.HexToAddress("0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85")
    tokenOut := common.HexToAddress("0x2218a117083f5B482B0bB821d27056Ba9c04b1D3")
    amountIn := big.NewInt(100) // example value; adjust as needed
    fee := big.NewInt(3000)     // Fee tier, adjust according to your scenario

    // Packing the input arguments
    data, err := parsedABI.Pack("quoteExactInputSingleV2", tokenIn, tokenOut, fee, amountIn, false)
    if err != nil {
        log.Fatal(err)
    }

    // Setting up the call message
    msg := ethereum.CallMsg{
        To:   &contractAddress,
        Data: data,
    }
    ctx := context.Background()

    // Making the call
    output, err := client.CallContract(ctx, msg, nil)
    if err != nil {
        log.Fatal(err)
    }

    // Process the output data (this depends on what the function returns)
    results, err := parsedABI.Unpack("quoteExactInputSingleV2", output)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Output from contract call:", results)

    // Type assert the result as *big.Int
    result, ok := results[0].(*big.Int)
    if !ok {
        return nil, fmt.Errorf("result type assertion to *big.Int failed")
    }

    return result, nil
}
