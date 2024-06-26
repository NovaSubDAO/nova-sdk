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
	SDai          string
}

var ConfigDetails = map[int64]ConfigData{
	1: {
		VaultAddress:  "0x83F20F44975D03b1b09e64809B757c47f942BEeA",
		SDai:          "0x83F20F44975D03b1b09e64809B757c47f942BEeA",
		VaultDecimals: 18,
	},
	10: {
		VaultAddress:  "0x7A8F265F2d1362ED8b6D5dd52E82741217BE8D3C",
		SDai:          "0x2218a117083f5B482B0bB821d27056Ba9c04b1D3",
		VaultDecimals: 18,
	},
}
