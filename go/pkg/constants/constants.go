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
	VaultAddressV2  string
	VaultDecimals int64
	SDai          string
}

var ConfigDetails = map[int64]ConfigData{
	1: {
		VaultAddressV2:  "0x83F20F44975D03b1b09e64809B757c47f942BEeA",
		SDai:          "0x83F20F44975D03b1b09e64809B757c47f942BEeA",
		VaultDecimals: 18,
	},
	10: {
		VaultAddressV2:  "0x04b12a2590BD808F7aC01f066aae0e2f48A3991C",
		SDai:          "0x2218a117083f5B482B0bB821d27056Ba9c04b1D3",
		VaultDecimals: 18,
	},
}

const (
	OptSDaiOracleAddress string = "0x33a3aB524A43E69f30bFd9Ae97d1Ec679FF00B64"
	OptQuoterV2Address   string = "0x89D8218ed5fF1e46d8dcd33fb0bbeE3be1621466"
)
