// Code generated by counterfeiter. DO NOT EDIT.
package algorithmfakes

import (
	sync "sync"

	db "github.com/concourse/concourse/atc/db"
	algorithm "github.com/concourse/concourse/atc/scheduler/algorithm"
)

type FakeInputMapper struct {
	MapInputsStub        func(*db.VersionsDB, db.Job, db.Resources) (db.InputMapping, bool, error)
	mapInputsMutex       sync.RWMutex
	mapInputsArgsForCall []struct {
		arg1 *db.VersionsDB
		arg2 db.Job
		arg3 db.Resources
	}
	mapInputsReturns struct {
		result1 db.InputMapping
		result2 bool
		result3 error
	}
	mapInputsReturnsOnCall map[int]struct {
		result1 db.InputMapping
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInputMapper) MapInputs(arg1 *db.VersionsDB, arg2 db.Job, arg3 db.Resources) (db.InputMapping, bool, error) {
	fake.mapInputsMutex.Lock()
	ret, specificReturn := fake.mapInputsReturnsOnCall[len(fake.mapInputsArgsForCall)]
	fake.mapInputsArgsForCall = append(fake.mapInputsArgsForCall, struct {
		arg1 *db.VersionsDB
		arg2 db.Job
		arg3 db.Resources
	}{arg1, arg2, arg3})
	fake.recordInvocation("MapInputs", []interface{}{arg1, arg2, arg3})
	fake.mapInputsMutex.Unlock()
	if fake.MapInputsStub != nil {
		return fake.MapInputsStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.mapInputsReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeInputMapper) MapInputsCallCount() int {
	fake.mapInputsMutex.RLock()
	defer fake.mapInputsMutex.RUnlock()
	return len(fake.mapInputsArgsForCall)
}

func (fake *FakeInputMapper) MapInputsCalls(stub func(*db.VersionsDB, db.Job, db.Resources) (db.InputMapping, bool, error)) {
	fake.mapInputsMutex.Lock()
	defer fake.mapInputsMutex.Unlock()
	fake.MapInputsStub = stub
}

func (fake *FakeInputMapper) MapInputsArgsForCall(i int) (*db.VersionsDB, db.Job, db.Resources) {
	fake.mapInputsMutex.RLock()
	defer fake.mapInputsMutex.RUnlock()
	argsForCall := fake.mapInputsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeInputMapper) MapInputsReturns(result1 db.InputMapping, result2 bool, result3 error) {
	fake.mapInputsMutex.Lock()
	defer fake.mapInputsMutex.Unlock()
	fake.MapInputsStub = nil
	fake.mapInputsReturns = struct {
		result1 db.InputMapping
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeInputMapper) MapInputsReturnsOnCall(i int, result1 db.InputMapping, result2 bool, result3 error) {
	fake.mapInputsMutex.Lock()
	defer fake.mapInputsMutex.Unlock()
	fake.MapInputsStub = nil
	if fake.mapInputsReturnsOnCall == nil {
		fake.mapInputsReturnsOnCall = make(map[int]struct {
			result1 db.InputMapping
			result2 bool
			result3 error
		})
	}
	fake.mapInputsReturnsOnCall[i] = struct {
		result1 db.InputMapping
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeInputMapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mapInputsMutex.RLock()
	defer fake.mapInputsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInputMapper) recordInvocation(key string, args []interface{}) {
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

var _ algorithm.InputMapper = new(FakeInputMapper)
