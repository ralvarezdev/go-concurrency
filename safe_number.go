package go_concurrency

import (
	"sync"
)

// SafeNumber is a concurrent-safe number
type SafeNumber struct {
	mu           sync.Mutex
	initialValue int
	value        int
}

// NewSafeNumber creates a new SafeNumber
func NewSafeNumber(initialValue int) *SafeNumber {
	return &SafeNumber{value: initialValue, initialValue: initialValue}
}

// Increment increments the number by 1
func (sn *SafeNumber) Increment() {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value++
}

// Decrement decrements the number by 1
func (sn *SafeNumber) Decrement() {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value--
}

// GetValue returns the current value of the number
func (sn *SafeNumber) GetValue() int {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	return sn.value
}

// GetValueAndIncrement returns the current value of the number and increments it by 1
func (sn *SafeNumber) GetValueAndIncrement() int {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value++
	return sn.value
}

// GetValueAndDecrement returns the current value of the number and decrements it by 1
func (sn *SafeNumber) GetValueAndDecrement() int {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value--
	return sn.value
}

// OperateValue performs an operation on the value of the number
func (sn *SafeNumber) OperateValue(operation func(int) int) {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value = operation(sn.value)
}

// GetAndOperateValue performs an operation on the value of the number and returns the result
func (sn *SafeNumber) GetAndOperateValue(operation func(int) int) int {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value = operation(sn.value)
	return sn.value
}

// ResetValue resets the value of the number to the initial value
func (sn *SafeNumber) ResetValue() {
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value = sn.initialValue
}
