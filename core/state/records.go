package state

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// BalanceDelta represents a change to a balance.
type BalanceDelta struct {
	Account common.Address
	Delta   [32]byte
}

// StorageDelta represents a change in the storage of an account's storage slot at the given key.
type StorageDelta struct {
	Account common.Address
	Key     common.Hash
	Delta   common.Hash
}

type JournalRecords struct {
	BalanceDeltas []*BalanceDelta
	StorageDeltas []*StorageDelta
}

func newJournalRecords() *JournalRecords {
	return &JournalRecords{
		BalanceDeltas: make([]*BalanceDelta, 0),
		StorageDeltas: make([]*StorageDelta, 0),
	}
}

// adds a new balance delta
func (r *JournalRecords) addBalanceDelta(account common.Address, delta *uint256.Int) {
	r.BalanceDeltas = append(r.BalanceDeltas, &BalanceDelta{
		Account: account,
		Delta:   delta.Bytes32(),
	})
}

// adds a new storage delta
func (r *JournalRecords) addStorageDelta(account common.Address, key, delta common.Hash) {
	r.StorageDeltas = append(r.StorageDeltas, &StorageDelta{
		Account: account,
		Key:     key,
		Delta:   delta,
	})
}

// EncodeBalanceDeltas encodes balance deltas into packed bytes.
// Each delta is 64 bytes: address (32) + value (32).
func (r *JournalRecords) EncodeBalanceDeltas() []byte {
	encoded := make([]byte, 0, len(r.BalanceDeltas)*64)
	for _, delta := range r.BalanceDeltas {
		encoded = append(encoded, common.LeftPadBytes(delta.Account.Bytes(), 32)...)
		encoded = append(encoded, delta.Delta[:]...)
	}
	return encoded
}

// EncodeStorageDeltas encodes storage deltas into packed bytes.
// Each delta is 96 bytes: address (32) + key (32) + value (32).
func (r *JournalRecords) EncodeStorageDeltas() []byte {
	encoded := make([]byte, 0, len(r.StorageDeltas)*96)
	for _, delta := range r.StorageDeltas {
		encoded = append(encoded, common.LeftPadBytes(delta.Account.Bytes(), 32)...)
		encoded = append(encoded, delta.Key.Bytes()...)
		encoded = append(encoded, delta.Delta.Bytes()...)
	}
	return encoded
}
