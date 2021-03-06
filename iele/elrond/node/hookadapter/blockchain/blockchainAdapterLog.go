package blockchainadapter

import (
	"encoding/hex"
	"fmt"
	"math/big"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
)

// GetInputAccounts yields all account input data that went through the adapter,
// if logging was enabled.
func (b *Blockchain) GetInputAccounts() []*vmi.OutputAccount {
	var result []*vmi.OutputAccount
	for _, v := range b.inputAccounts {
		result = append(result, v)
	}
	return result
}

func (b *Blockchain) logNewAddress(creatorAddress []byte, creatorNonce uint64, result []byte) {
	if !b.LogToConsole {
		return
	}
	fmt.Printf("\n#newAddr(%s, %d) => %s",
		hex.EncodeToString(creatorAddress),
		creatorNonce,
		hex.EncodeToString(result))
}

var balanceZero = big.NewInt(0)

func (b *Blockchain) logBalance(address []byte, balance *big.Int) {
	if !b.LogToConsole {
		return
	}
	if b.inputAccounts == nil {
		b.inputAccounts = make(map[string]*vmi.OutputAccount)
	}
	acc, found := b.inputAccounts[string(address)]
	if !found {
		acc = &vmi.OutputAccount{Address: address}
		b.inputAccounts[string(address)] = acc
	}
	acc.Balance = balance
}

func (b *Blockchain) logNonce(address []byte, nonce uint64) {
	if !b.LogToConsole {
		return
	}
	if b.inputAccounts == nil {
		b.inputAccounts = make(map[string]*vmi.OutputAccount)
	}
	acc, found := b.inputAccounts[string(address)]
	if !found {
		acc = &vmi.OutputAccount{Address: address}
		b.inputAccounts[string(address)] = acc
	}
	acc.Nonce = nonce
}

func (b *Blockchain) logStorage(address []byte, index []byte, data []byte) {
	if !b.LogToConsole {
		return
	}
	if b.inputAccounts == nil {
		b.inputAccounts = make(map[string]*vmi.OutputAccount)
	}
	acc, found := b.inputAccounts[string(address)]
	if !found {
		acc = &vmi.OutputAccount{Address: address}
		b.inputAccounts[string(address)] = acc
	}
	acc.StorageUpdates[string(index)] =  &vmi.StorageUpdate{
		Offset: index,
		Data:   data,
	}
}
