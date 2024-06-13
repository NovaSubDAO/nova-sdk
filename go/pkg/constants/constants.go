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
