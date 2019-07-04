// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:18:31.546

package ieletestingmodel

import (
	"fmt"
	"strings"
)

// KObject defines a K item object that is managed by the model
type KObject interface {
	referenceType() kreferenceType
	equals(ms *ModelState, other KObject) bool
	deepCopy(ms *ModelState) KObject
	prettyPrint(ms *ModelState, sb *strings.Builder, indent int)
	kprint(ms *ModelState, sb *strings.Builder)
	collectionsToK(ms *ModelState) KReference
	increaseUsage(ms *ModelState)
	decreaseUsage(ms *ModelState)
	recycleUnused(ms *ModelState)
	preserve(ms *ModelState)
}

type objectReuseStatus int

const (
    active objectReuseStatus = iota
    inRecycleBin
    preserved
)

// ModelState holds the state of the executor at a certain moment
type ModelState struct {
	initialized bool

	// allKs keeps all KSequences into one large structure
	// all KSequences point to this structure
	allKs *ksequenceSliceContainer

	// keeps big int objects, big int references point here
	bigInts []*bigInt

	// recycle bin for big ints
	// works as a stack
	bigIntRecycleBin []KReference

	// keeps object types mixed together
	allObjects []KObject

	// memoTables is a structure containing all memoization maps.
	// Memoization tables are implemented as maps of maps of maps of ...
	memoTables map[MemoTable]interface{}
}

// constantsModel is another instance of the model, but which only contains a few constants.
var constantsModel = NewModel()

func (ms *ModelState) getReferencedObject(ref KReference) KObject {
	index := int(ref.value1)
	if ref.constantObject {
		return constantsModel.allObjects[index]
	}
	if index >= len(ms.allObjects) {
		panic("trying to reference object beyond allocated objects")
	}
	return ms.allObjects[index]
}

func (ms *ModelState) addObject(obj KObject) KReference {
	newIndex := len(ms.allObjects)
	ms.allObjects = append(ms.allObjects, obj)
	return KReference{refType: obj.referenceType(), constantObject: false, value1: uint32(newIndex), value2: 0}
}

func addConstantObject(obj KObject) KReference {
	newIndex := len(constantsModel.allObjects)
	constantsModel.allObjects = append(constantsModel.allObjects, obj)
	return KReference{refType: obj.referenceType(), constantObject: true, value1: uint32(newIndex), value2: 0}
}

// NewModel creates a new blank model.
func NewModel() *ModelState {
	ms := &ModelState{}
	ms.Init()
	return ms
}

// Init prepares model for execution
func (ms *ModelState) Init() {
	if ms.initialized {
		return
	}
	ms.initialized = true
	ms.allKs = &ksequenceSliceContainer{}
	ms.allObjects = nil
	ms.memoTables = nil
}

// ClearModel ... clean up any data left from previous executions, to save memory
func (ms *ModelState) ClearModel() {
	ms.initialized = false
	ms.Init()
}

// PrintStats simply prints some statistics to the console.
// Useful for checking the size of the model data.
func (ms *ModelState) PrintStats() {
	fmt.Printf("Nr. BigInt objects: %d\n", len(ms.bigInts))
	fmt.Printf("Nr. objects: %d\n", len(ms.allObjects))
	fmt.Printf("Nr. K sequence slices: %d\n", len(ms.allKs.allSlices))
	fmt.Printf("Recycle bin\n")
	fmt.Printf("     BigInt    %d\n", len(ms.bigIntRecycleBin))
}
