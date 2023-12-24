package day24

import (
	"math/big"
)

type Coord struct {
	x, y, z int
}

func NewCoord(x, y, z int) Coord {
	return Coord{x: x, y: y, z: z}
}

func (c Coord) Length() float64 {
	xbig := big.NewFloat(float64(c.x))
	ybig := big.NewFloat(float64(c.y))
	zbig := big.NewFloat(float64(c.z))

	ret, _ := big.NewFloat(0).Sqrt(big.NewFloat(0).Add(big.NewFloat(0).Add(big.NewFloat(0).Mul(xbig, xbig), big.NewFloat(0).Mul(ybig, ybig)), big.NewFloat(0).Mul(zbig, zbig))).Float64()
	return ret
}
