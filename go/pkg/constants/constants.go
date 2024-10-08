package constants

type Stablecoin string

const (
	DAI  Stablecoin = "DAI"
	USDC Stablecoin = "USDC"
	USDT Stablecoin = "USDT"
)

type StablecoinData struct {
	Address  string
	Decimals int
}

var StablecoinDetails = map[int64]map[Stablecoin]StablecoinData{
	1: {
		DAI: StablecoinData{
			Address:  "0x6B175474E89094C44Da98b954EedeAC495271d0F",
			Decimals: 18,
		},
	},
	10: {
		USDC: StablecoinData{
			Address:  "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85",
			Decimals: 6,
		},
	},
}

type ConfigData struct {
	VaultAddress  string
	VaultDecimals int64
	Savings       string
	LiquidityPool string
}

var ConfigDetails = map[int64]ConfigData{
	1: {
		VaultAddress:  "0x83F20F44975D03b1b09e64809B757c47f942BEeA",
		Savings:       "0x83F20F44975D03b1b09e64809B757c47f942BEeA",
		LiquidityPool: "0x83F20F44975D03b1b09e64809B757c47f942BEeA",
		VaultDecimals: 18,
	},
	10: {
		VaultAddress:  "0xbf3ccf927eD469229ed834FD67004533f37a7291", // NovaVaultV1
		Savings:       "0x2218a117083f5B482B0bB821d27056Ba9c04b1D3",
		LiquidityPool: "0x131525f3FA23d65DC2B1EB8B6483a28c43B06916",
		VaultDecimals: 18,
	},
}

const (
	OptSavingsOracleAddress string = "0x33a3aB524A43E69f30bFd9Ae97d1Ec679FF00B64"
	OptQuoterV2Address      string = "0x89D8218ed5fF1e46d8dcd33fb0bbeE3be1621466"
)
