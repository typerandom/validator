package validator

import (
	"sync"
)

var globalLock sync.Mutex
var globalInitialized bool
var globalDefaultValidator *validator

func getGlobalValidator() *validator {
	if !globalInitialized {
		globalLock.Lock()
		defer globalLock.Unlock()
		if !globalInitialized {
			globalInitialized = true
			globalDefaultValidator = newValidator()
		}
	}
	return globalDefaultValidator
}
