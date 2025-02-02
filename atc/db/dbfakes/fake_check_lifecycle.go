// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	sync "sync"
	time "time"

	db "github.com/concourse/concourse/atc/db"
)

type FakeCheckLifecycle struct {
	RemoveExpiredChecksStub        func(time.Duration) error
	removeExpiredChecksMutex       sync.RWMutex
	removeExpiredChecksArgsForCall []struct {
		arg1 time.Duration
	}
	removeExpiredChecksReturns struct {
		result1 error
	}
	removeExpiredChecksReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCheckLifecycle) RemoveExpiredChecks(arg1 time.Duration) error {
	fake.removeExpiredChecksMutex.Lock()
	ret, specificReturn := fake.removeExpiredChecksReturnsOnCall[len(fake.removeExpiredChecksArgsForCall)]
	fake.removeExpiredChecksArgsForCall = append(fake.removeExpiredChecksArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	fake.recordInvocation("RemoveExpiredChecks", []interface{}{arg1})
	fake.removeExpiredChecksMutex.Unlock()
	if fake.RemoveExpiredChecksStub != nil {
		return fake.RemoveExpiredChecksStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeExpiredChecksReturns
	return fakeReturns.result1
}

func (fake *FakeCheckLifecycle) RemoveExpiredChecksCallCount() int {
	fake.removeExpiredChecksMutex.RLock()
	defer fake.removeExpiredChecksMutex.RUnlock()
	return len(fake.removeExpiredChecksArgsForCall)
}

func (fake *FakeCheckLifecycle) RemoveExpiredChecksCalls(stub func(time.Duration) error) {
	fake.removeExpiredChecksMutex.Lock()
	defer fake.removeExpiredChecksMutex.Unlock()
	fake.RemoveExpiredChecksStub = stub
}

func (fake *FakeCheckLifecycle) RemoveExpiredChecksArgsForCall(i int) time.Duration {
	fake.removeExpiredChecksMutex.RLock()
	defer fake.removeExpiredChecksMutex.RUnlock()
	argsForCall := fake.removeExpiredChecksArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCheckLifecycle) RemoveExpiredChecksReturns(result1 error) {
	fake.removeExpiredChecksMutex.Lock()
	defer fake.removeExpiredChecksMutex.Unlock()
	fake.RemoveExpiredChecksStub = nil
	fake.removeExpiredChecksReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckLifecycle) RemoveExpiredChecksReturnsOnCall(i int, result1 error) {
	fake.removeExpiredChecksMutex.Lock()
	defer fake.removeExpiredChecksMutex.Unlock()
	fake.RemoveExpiredChecksStub = nil
	if fake.removeExpiredChecksReturnsOnCall == nil {
		fake.removeExpiredChecksReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeExpiredChecksReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCheckLifecycle) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.removeExpiredChecksMutex.RLock()
	defer fake.removeExpiredChecksMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCheckLifecycle) recordInvocation(key string, args []interface{}) {
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

var _ db.CheckLifecycle = new(FakeCheckLifecycle)
