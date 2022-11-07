package functions

import "sync"

// Creates a function that is restricted to invoking func once. Repeat calls
// to the function return the value of the first invocation.
//
// Complexity: O(1)
func Once(function func() []interface{}) func() []interface{} {
	lock := sync.Mutex{}
	done := false
	cache := []interface{}{}
	return func() []interface{} {
		lock.Lock()
		if done {
			lock.Unlock()
		} else {
			done = true
			cache = function()
			lock.Unlock()
		}
		return cache
	}
}
