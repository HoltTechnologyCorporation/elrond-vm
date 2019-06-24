// File generated by the K Framework Go backend. Timestamp: 2019-06-25 00:00:28.701

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
)

func (i *Interpreter) freshFunction (s m.Sort, config m.K, counter int) (m.K, error) {
	switch s {
		case m.SortID:
			return i.evalFreshID(m.NewIntFromInt(counter), config, -1)
		case m.SortInt:
			return i.evalFreshInt(m.NewIntFromInt(counter), config, -1)
		default:
			panic("Cannot find fresh function for sort " + s.Name())
	}
}

