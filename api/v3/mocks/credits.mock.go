// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/RTradeLtd/database/models"
)

type FakeCreditsManager struct {
	AddCreditsStub        func(string, float64) (*models.User, error)
	addCreditsMutex       sync.RWMutex
	addCreditsArgsForCall []struct {
		arg1 string
		arg2 float64
	}
	addCreditsReturns struct {
		result1 *models.User
		result2 error
	}
	addCreditsReturnsOnCall map[int]struct {
		result1 *models.User
		result2 error
	}
	RemoveCreditsStub        func(string, float64) (*models.User, error)
	removeCreditsMutex       sync.RWMutex
	removeCreditsArgsForCall []struct {
		arg1 string
		arg2 float64
	}
	removeCreditsReturns struct {
		result1 *models.User
		result2 error
	}
	removeCreditsReturnsOnCall map[int]struct {
		result1 *models.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreditsManager) AddCredits(arg1 string, arg2 float64) (*models.User, error) {
	fake.addCreditsMutex.Lock()
	ret, specificReturn := fake.addCreditsReturnsOnCall[len(fake.addCreditsArgsForCall)]
	fake.addCreditsArgsForCall = append(fake.addCreditsArgsForCall, struct {
		arg1 string
		arg2 float64
	}{arg1, arg2})
	fake.recordInvocation("AddCredits", []interface{}{arg1, arg2})
	fake.addCreditsMutex.Unlock()
	if fake.AddCreditsStub != nil {
		return fake.AddCreditsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.addCreditsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreditsManager) AddCreditsCallCount() int {
	fake.addCreditsMutex.RLock()
	defer fake.addCreditsMutex.RUnlock()
	return len(fake.addCreditsArgsForCall)
}

func (fake *FakeCreditsManager) AddCreditsCalls(stub func(string, float64) (*models.User, error)) {
	fake.addCreditsMutex.Lock()
	defer fake.addCreditsMutex.Unlock()
	fake.AddCreditsStub = stub
}

func (fake *FakeCreditsManager) AddCreditsArgsForCall(i int) (string, float64) {
	fake.addCreditsMutex.RLock()
	defer fake.addCreditsMutex.RUnlock()
	argsForCall := fake.addCreditsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCreditsManager) AddCreditsReturns(result1 *models.User, result2 error) {
	fake.addCreditsMutex.Lock()
	defer fake.addCreditsMutex.Unlock()
	fake.AddCreditsStub = nil
	fake.addCreditsReturns = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeCreditsManager) AddCreditsReturnsOnCall(i int, result1 *models.User, result2 error) {
	fake.addCreditsMutex.Lock()
	defer fake.addCreditsMutex.Unlock()
	fake.AddCreditsStub = nil
	if fake.addCreditsReturnsOnCall == nil {
		fake.addCreditsReturnsOnCall = make(map[int]struct {
			result1 *models.User
			result2 error
		})
	}
	fake.addCreditsReturnsOnCall[i] = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeCreditsManager) RemoveCredits(arg1 string, arg2 float64) (*models.User, error) {
	fake.removeCreditsMutex.Lock()
	ret, specificReturn := fake.removeCreditsReturnsOnCall[len(fake.removeCreditsArgsForCall)]
	fake.removeCreditsArgsForCall = append(fake.removeCreditsArgsForCall, struct {
		arg1 string
		arg2 float64
	}{arg1, arg2})
	fake.recordInvocation("RemoveCredits", []interface{}{arg1, arg2})
	fake.removeCreditsMutex.Unlock()
	if fake.RemoveCreditsStub != nil {
		return fake.RemoveCreditsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.removeCreditsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreditsManager) RemoveCreditsCallCount() int {
	fake.removeCreditsMutex.RLock()
	defer fake.removeCreditsMutex.RUnlock()
	return len(fake.removeCreditsArgsForCall)
}

func (fake *FakeCreditsManager) RemoveCreditsCalls(stub func(string, float64) (*models.User, error)) {
	fake.removeCreditsMutex.Lock()
	defer fake.removeCreditsMutex.Unlock()
	fake.RemoveCreditsStub = stub
}

func (fake *FakeCreditsManager) RemoveCreditsArgsForCall(i int) (string, float64) {
	fake.removeCreditsMutex.RLock()
	defer fake.removeCreditsMutex.RUnlock()
	argsForCall := fake.removeCreditsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCreditsManager) RemoveCreditsReturns(result1 *models.User, result2 error) {
	fake.removeCreditsMutex.Lock()
	defer fake.removeCreditsMutex.Unlock()
	fake.RemoveCreditsStub = nil
	fake.removeCreditsReturns = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeCreditsManager) RemoveCreditsReturnsOnCall(i int, result1 *models.User, result2 error) {
	fake.removeCreditsMutex.Lock()
	defer fake.removeCreditsMutex.Unlock()
	fake.RemoveCreditsStub = nil
	if fake.removeCreditsReturnsOnCall == nil {
		fake.removeCreditsReturnsOnCall = make(map[int]struct {
			result1 *models.User
			result2 error
		})
	}
	fake.removeCreditsReturnsOnCall[i] = struct {
		result1 *models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeCreditsManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addCreditsMutex.RLock()
	defer fake.addCreditsMutex.RUnlock()
	fake.removeCreditsMutex.RLock()
	defer fake.removeCreditsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreditsManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
