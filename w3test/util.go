package w3test

import (
	"crypto/rand"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	weth9BalancePos   = big.NewInt(3)
	weth9AllowancePos = big.NewInt(4)
)

// WETH9BalanceSlot returns the storage slot of the WETH9 balance of the given
// address.
func WETH9BalanceSlot(addr common.Address) common.Hash {
	data := make([]byte, 64)
	copy(data[12:32], addr[:])
	weth9BalancePos.FillBytes(data[32:])

	return crypto.Keccak256Hash(data)
}

// WETH9AllowanceSlot returns the storage slot of the WETH9 allowance of the
// given owner for the given spender.
func WETH9AllowanceSlot(owner, spender common.Address) common.Hash {
	data := make([]byte, 64)
	copy(data[12:32], owner[:])
	weth9AllowancePos.FillBytes(data[32:])

	copy(data[32:], crypto.Keccak256(data))
	copy(data[12:32], spender[:])

	return crypto.Keccak256Hash(data)
}

// RandA returns a random address.
func RandA() (addr common.Address) {
	rand.Read(addr[:])
	return
}
