package go_concurrency

import (
	"sync"
)

type (
	// SafeNumber is a concurrent-safe number
	SafeNumber struct {
		mu           sync.Mutex
		initialValue int
		value        int
	}
)

// NewSafeNumber creates a new SafeNumber
//
// Parameters:
//
//   - initialValue: the initial value of the number
//
// Returns:
//
//   - *SafeNumber: the created SafeNumber
func NewSafeNumber(initialValue int) *SafeNumber {
	return &SafeNumber{value: initialValue, initialValue: initialValue}
}

// Increment increments the number by 1
func (sn *SafeNumber) Increment() {
	if sn == nil {
		return
	}
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value++
}

// Decrement decrements the number by 1
func (sn *SafeNumber) Decrement() {
	if sn == nil {
		return
	}
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value--
}

// GetValue returns the current value of the number
//
// Returns:
//
//   - int: the current value of the number
func (sn *SafeNumber) GetValue() int {
	if sn == nil {
		return 0
	}
	sn.mu.Lock()
	defer sn.mu.Unlock()
	return sn.value
}

// IncrementAndGetValue increments the number by 1 and returns the new value
//
// Returns:
//
//   - int: the new value of the number
func (sn *SafeNumber) IncrementAndGetValue() int {
	if sn == nil {
		return 0
	}
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value++
	return sn.value
}

// DecrementAndGetValue decrements the number by 1 and returns the new value
//
// Returns:
//
//   - int: the new value of the number
func (sn *SafeNumber) DecrementAndGetValue() int {
	if sn == nil {
		return 0
	}
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value--
	return sn.value
}

// OperateValue performs an operation on the value of the number
//
// Parameters:
//
//   - operation: a function that takes the current value and returns the new value
func (sn *SafeNumber) OperateValue(operation func(int) int) {
	if sn == nil {
		return
	}
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value = operation(sn.value)
}

// OperateAndGetValue performs an operation on the value of the number and returns the result
//
// Parameters:
//
//   - operation: a function that takes the current value and returns the new value
//
// Returns:
//
//   - int: the new value of the number
func (sn *SafeNumber) OperateAndGetValue(operation func(int) int) int {
	if sn == nil {
		return 0
	}
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value = operation(sn.value)
	return sn.value
}

// ResetValue resets the value of the number to the initial value
func (sn *SafeNumber) ResetValue() {
	if sn == nil {
		return
	}
	sn.mu.Lock()
	defer sn.mu.Unlock()
	sn.value = sn.initialValue
}
