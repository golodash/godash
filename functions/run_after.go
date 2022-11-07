package functions

import (
	"errors"
	"sync"
	"time"

	"github.com/golodash/godash/internal"
)

type runCancel interface {
	Run() []interface{}
	Cancel() (bool, error)
}

type runCancelStruct struct {
	function   func() []interface{}
	wait       time.Duration
	parallel   bool
	lock       *sync.Mutex
	canceled   bool
	inProgress bool
	done       bool
}

// Returns an interface with two methods. one called Run which executes the
// function and other one called Cancel which cancels executing of current
// executed Run function.
//
// You can use WrapFunc to generate `function`.
//
// Complexity: O(1)
func RunAfter(function func() []interface{}, wait time.Duration, parallel bool) runCancel {
	if ok := internal.IsFunc(function); !ok {
		panic("`function` input is not a function")
	}

	r := runCancelStruct{
		function:   function,
		wait:       wait,
		parallel:   parallel,
		lock:       &sync.Mutex{},
		canceled:   false,
		inProgress: false,
		done:       false,
	}

	return &r
}

func (r *runCancelStruct) Run() []interface{} {
	r.lock.Lock()
	r.reset()
	r.lock.Unlock()
	if r.parallel {
		ch := make(chan bool)
		go func() {
			value, ok := <-ch
			if ok && value {
				r.function()
				r.lock.Lock()
				defer r.lock.Unlock()
				r.done = true
			} else {
				return
			}
		}()
		go func() {
			time.Sleep(r.wait)
			r.lock.Lock()
			defer r.lock.Unlock()
			if !r.canceled {
				r.inProgress = true
				ch <- true
			} else {
				ch <- false
			}
		}()
	} else {
		time.Sleep(r.wait)
		return r.function()
	}

	return nil
}

func (r *runCancelStruct) Cancel() (bool, error) {
	if r.parallel {
		r.lock.Lock()
		defer r.lock.Unlock()
		if r.inProgress && !r.done {
			return false, errors.New("function in progress")
		} else if r.done {
			return false, errors.New("function done")
		}

		r.canceled = true
		return true, nil
	} else {
		return false, errors.New("parallel is off")
	}
}

func (r *runCancelStruct) reset() {
	r.canceled = false
	r.inProgress = false
	r.done = false
}
