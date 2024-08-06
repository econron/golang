package user

import (
	"math/big"
)

type Shop struct {
	id uint64
	name string
	postalCode string
	city string
	street string
	phoneNumber string
	latitude big.Float
	longitute big.Float
}