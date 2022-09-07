package testbox

import "sync"

func New() *testBox {
	return &testBox{}
}

type testBox struct {
	values       [100]float64
	valueMutex   sync.Mutex
	valueRWMutex sync.RWMutex
}

func (box *testBox) GetValueM(index int) float64 {
	box.valueMutex.Lock()
	defer box.valueMutex.Unlock()
	return box.values[index]
}

func (box *testBox) SetValueM(index int, value float64) {
	box.valueMutex.Lock()
	box.values[index] = value
	box.valueMutex.Unlock()
}

func (box *testBox) GetValueRW(index int) float64 {
	box.valueRWMutex.RLock()
	defer box.valueRWMutex.RUnlock()
	return box.values[index]
}

func (box *testBox) SetValueRW(index int, value float64) {
	box.valueRWMutex.Lock()
	box.values[index] = value
	box.valueRWMutex.Unlock()
}
