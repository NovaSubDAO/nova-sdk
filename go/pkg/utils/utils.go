package utils

import (
    "math/big"
)

var RAY *big.Int = big.NewInt(0).Exp(big.NewInt(10), big.NewInt(27), nil) // Initialize RAY as a big.Int

func Rpow(x, n *big.Int) *big.Int {
    z := new(big.Int)
    if x.Cmp(big.NewInt(0)) == 0 {
        if n.Cmp(big.NewInt(0)) == 0 {
            z.Set(RAY) // z = RAY if x == 0 && n == 0
        } else {
            z.SetInt64(0) // z = 0 if x == 0 && n != 0
        }
    } else {
        half := new(big.Int).Div(RAY, big.NewInt(2)) // half = RAY / 2 for rounding
        if new(big.Int).Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
            z.Set(RAY) // z = RAY if n % 2 == 0
        } else {
            z.Set(x) // z = x if n % 2 != 0
        }

        for n = new(big.Int).Div(n, big.NewInt(2)); n.Sign() > 0; n = new(big.Int).Div(n, big.NewInt(2)) {
            xx := new(big.Int).Mul(x, x)
            xxRound := new(big.Int).Add(xx, half)
            x = new(big.Int).Div(xxRound, RAY)

            if new(big.Int).Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) != 0 {
                zx := new(big.Int).Mul(z, x)
                zxRound := new(big.Int).Add(zx, half)
                z = new(big.Int).Div(zxRound, RAY)
            }
        }
    }
    return z
}
