package eth

import (
	"github.com/jonkofee/go-ethereum/core"
	"github.com/jonkofee/go-ethereum/core/state"
	"github.com/jonkofee/go-ethereum/core/types"
	"github.com/jonkofee/go-ethereum/core/vm"
	"github.com/jonkofee/go-ethereum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, error) {
	return eth.stateAtTransaction(block, txIndex, reexec)
}
