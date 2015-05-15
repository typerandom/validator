package validator

import (
	"sync"
)

var globalLock sync.Mutex
var globalInitialized bool
var globalDefaultValidator *Validator

func getGlobalValidator() *Validator {
	if !globalInitialized {
		globalLock.Lock()
		defer globalLock.Unlock()
		if !globalInitialized {
			globalInitialized = true
			globalDefaultValidator = New()
		}
	}
	return globalDefaultValidator
}
