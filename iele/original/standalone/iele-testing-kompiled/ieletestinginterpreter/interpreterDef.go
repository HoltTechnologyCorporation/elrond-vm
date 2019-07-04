// File generated by the K Framework Go backend. Timestamp: 2019-07-04 13:18:31.546

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
	krypto "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/hookadapter/krypto"
)

// Interpreter is a container with a reference to model and basic options
type Interpreter struct {
	Model         *m.ModelState
	currentStep   int
	MaxSteps      int
	state         m.KReference
	traceHandlers []traceHandler
	Verbose       bool

	kryptoRef *krypto.Krypto
}
// NewInterpreter creates a new interpreter instance
func NewInterpreter(kryptoRef *krypto.Krypto) *Interpreter {
	model := &m.ModelState{}
	model.Init()

	return &Interpreter {
		Model:         model,
		MaxSteps:      0,
		currentStep:   -1, // meaning that no processing started yet
		state:         m.NullReference,
		traceHandlers: nil,
		Verbose:       false,
		kryptoRef: kryptoRef,
	}
}
