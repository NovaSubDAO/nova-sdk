package constants

type Stablecoin string

const (
	DAI  Stablecoin = "DAI"
	USDC Stablecoin = "USDC"
	USDT Stablecoin = "USDT"
)

var StablecoinAddresses = map[int64]map[Stablecoin]string{
	1: {
		DAI: "0x6B175474E89094C44Da98b954EedeAC495271d0F",
	},
	10: {
		USDC: "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85",
	},
}
