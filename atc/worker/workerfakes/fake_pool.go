// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/db/lock"
	"github.com/concourse/concourse/atc/worker"
)

type FakePool struct {
	AcquireContainerCreatingLockStub        func(lager.Logger) (lock.Lock, bool, error)
	acquireContainerCreatingLockMutex       sync.RWMutex
	acquireContainerCreatingLockArgsForCall []struct {
		arg1 lager.Logger
	}
	acquireContainerCreatingLockReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	acquireContainerCreatingLockReturnsOnCall map[int]struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	FindOrChooseWorkerStub        func(lager.Logger, worker.WorkerSpec) (worker.Worker, error)
	findOrChooseWorkerMutex       sync.RWMutex
	findOrChooseWorkerArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.WorkerSpec
	}
	findOrChooseWorkerReturns struct {
		result1 worker.Worker
		result2 error
	}
	findOrChooseWorkerReturnsOnCall map[int]struct {
		result1 worker.Worker
		result2 error
	}
	FindOrChooseWorkerForContainerStub        func(context.Context, lager.Logger, db.ContainerOwner, worker.ContainerSpec, db.ContainerMetadata, worker.WorkerSpec, worker.ContainerPlacementStrategy) (worker.Worker, error)
	findOrChooseWorkerForContainerMutex       sync.RWMutex
	findOrChooseWorkerForContainerArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 db.ContainerOwner
		arg4 worker.ContainerSpec
		arg5 db.ContainerMetadata
		arg6 worker.WorkerSpec
		arg7 worker.ContainerPlacementStrategy
	}
	findOrChooseWorkerForContainerReturns struct {
		result1 worker.Worker
		result2 error
	}
	findOrChooseWorkerForContainerReturnsOnCall map[int]struct {
		result1 worker.Worker
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePool) AcquireContainerCreatingLock(arg1 lager.Logger) (lock.Lock, bool, error) {
	fake.acquireContainerCreatingLockMutex.Lock()
	ret, specificReturn := fake.acquireContainerCreatingLockReturnsOnCall[len(fake.acquireContainerCreatingLockArgsForCall)]
	fake.acquireContainerCreatingLockArgsForCall = append(fake.acquireContainerCreatingLockArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("AcquireContainerCreatingLock", []interface{}{arg1})
	fake.acquireContainerCreatingLockMutex.Unlock()
	if fake.AcquireContainerCreatingLockStub != nil {
		return fake.AcquireContainerCreatingLockStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.acquireContainerCreatingLockReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakePool) AcquireContainerCreatingLockCallCount() int {
	fake.acquireContainerCreatingLockMutex.RLock()
	defer fake.acquireContainerCreatingLockMutex.RUnlock()
	return len(fake.acquireContainerCreatingLockArgsForCall)
}

func (fake *FakePool) AcquireContainerCreatingLockCalls(stub func(lager.Logger) (lock.Lock, bool, error)) {
	fake.acquireContainerCreatingLockMutex.Lock()
	defer fake.acquireContainerCreatingLockMutex.Unlock()
	fake.AcquireContainerCreatingLockStub = stub
}

func (fake *FakePool) AcquireContainerCreatingLockArgsForCall(i int) lager.Logger {
	fake.acquireContainerCreatingLockMutex.RLock()
	defer fake.acquireContainerCreatingLockMutex.RUnlock()
	argsForCall := fake.acquireContainerCreatingLockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePool) AcquireContainerCreatingLockReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.acquireContainerCreatingLockMutex.Lock()
	defer fake.acquireContainerCreatingLockMutex.Unlock()
	fake.AcquireContainerCreatingLockStub = nil
	fake.acquireContainerCreatingLockReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePool) AcquireContainerCreatingLockReturnsOnCall(i int, result1 lock.Lock, result2 bool, result3 error) {
	fake.acquireContainerCreatingLockMutex.Lock()
	defer fake.acquireContainerCreatingLockMutex.Unlock()
	fake.AcquireContainerCreatingLockStub = nil
	if fake.acquireContainerCreatingLockReturnsOnCall == nil {
		fake.acquireContainerCreatingLockReturnsOnCall = make(map[int]struct {
			result1 lock.Lock
			result2 bool
			result3 error
		})
	}
	fake.acquireContainerCreatingLockReturnsOnCall[i] = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakePool) FindOrChooseWorker(arg1 lager.Logger, arg2 worker.WorkerSpec) (worker.Worker, error) {
	fake.findOrChooseWorkerMutex.Lock()
	ret, specificReturn := fake.findOrChooseWorkerReturnsOnCall[len(fake.findOrChooseWorkerArgsForCall)]
	fake.findOrChooseWorkerArgsForCall = append(fake.findOrChooseWorkerArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.WorkerSpec
	}{arg1, arg2})
	fake.recordInvocation("FindOrChooseWorker", []interface{}{arg1, arg2})
	fake.findOrChooseWorkerMutex.Unlock()
	if fake.FindOrChooseWorkerStub != nil {
		return fake.FindOrChooseWorkerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findOrChooseWorkerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePool) FindOrChooseWorkerCallCount() int {
	fake.findOrChooseWorkerMutex.RLock()
	defer fake.findOrChooseWorkerMutex.RUnlock()
	return len(fake.findOrChooseWorkerArgsForCall)
}

func (fake *FakePool) FindOrChooseWorkerCalls(stub func(lager.Logger, worker.WorkerSpec) (worker.Worker, error)) {
	fake.findOrChooseWorkerMutex.Lock()
	defer fake.findOrChooseWorkerMutex.Unlock()
	fake.FindOrChooseWorkerStub = stub
}

func (fake *FakePool) FindOrChooseWorkerArgsForCall(i int) (lager.Logger, worker.WorkerSpec) {
	fake.findOrChooseWorkerMutex.RLock()
	defer fake.findOrChooseWorkerMutex.RUnlock()
	argsForCall := fake.findOrChooseWorkerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePool) FindOrChooseWorkerReturns(result1 worker.Worker, result2 error) {
	fake.findOrChooseWorkerMutex.Lock()
	defer fake.findOrChooseWorkerMutex.Unlock()
	fake.FindOrChooseWorkerStub = nil
	fake.findOrChooseWorkerReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakePool) FindOrChooseWorkerReturnsOnCall(i int, result1 worker.Worker, result2 error) {
	fake.findOrChooseWorkerMutex.Lock()
	defer fake.findOrChooseWorkerMutex.Unlock()
	fake.FindOrChooseWorkerStub = nil
	if fake.findOrChooseWorkerReturnsOnCall == nil {
		fake.findOrChooseWorkerReturnsOnCall = make(map[int]struct {
			result1 worker.Worker
			result2 error
		})
	}
	fake.findOrChooseWorkerReturnsOnCall[i] = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakePool) FindOrChooseWorkerForContainer(arg1 context.Context, arg2 lager.Logger, arg3 db.ContainerOwner, arg4 worker.ContainerSpec, arg5 db.ContainerMetadata, arg6 worker.WorkerSpec, arg7 worker.ContainerPlacementStrategy) (worker.Worker, error) {
	fake.findOrChooseWorkerForContainerMutex.Lock()
	ret, specificReturn := fake.findOrChooseWorkerForContainerReturnsOnCall[len(fake.findOrChooseWorkerForContainerArgsForCall)]
	fake.findOrChooseWorkerForContainerArgsForCall = append(fake.findOrChooseWorkerForContainerArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 db.ContainerOwner
		arg4 worker.ContainerSpec
		arg5 db.ContainerMetadata
		arg6 worker.WorkerSpec
		arg7 worker.ContainerPlacementStrategy
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("FindOrChooseWorkerForContainer", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.findOrChooseWorkerForContainerMutex.Unlock()
	if fake.FindOrChooseWorkerForContainerStub != nil {
		return fake.FindOrChooseWorkerForContainerStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findOrChooseWorkerForContainerReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePool) FindOrChooseWorkerForContainerCallCount() int {
	fake.findOrChooseWorkerForContainerMutex.RLock()
	defer fake.findOrChooseWorkerForContainerMutex.RUnlock()
	return len(fake.findOrChooseWorkerForContainerArgsForCall)
}

func (fake *FakePool) FindOrChooseWorkerForContainerCalls(stub func(context.Context, lager.Logger, db.ContainerOwner, worker.ContainerSpec, db.ContainerMetadata, worker.WorkerSpec, worker.ContainerPlacementStrategy) (worker.Worker, error)) {
	fake.findOrChooseWorkerForContainerMutex.Lock()
	defer fake.findOrChooseWorkerForContainerMutex.Unlock()
	fake.FindOrChooseWorkerForContainerStub = stub
}

func (fake *FakePool) FindOrChooseWorkerForContainerArgsForCall(i int) (context.Context, lager.Logger, db.ContainerOwner, worker.ContainerSpec, db.ContainerMetadata, worker.WorkerSpec, worker.ContainerPlacementStrategy) {
	fake.findOrChooseWorkerForContainerMutex.RLock()
	defer fake.findOrChooseWorkerForContainerMutex.RUnlock()
	argsForCall := fake.findOrChooseWorkerForContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakePool) FindOrChooseWorkerForContainerReturns(result1 worker.Worker, result2 error) {
	fake.findOrChooseWorkerForContainerMutex.Lock()
	defer fake.findOrChooseWorkerForContainerMutex.Unlock()
	fake.FindOrChooseWorkerForContainerStub = nil
	fake.findOrChooseWorkerForContainerReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakePool) FindOrChooseWorkerForContainerReturnsOnCall(i int, result1 worker.Worker, result2 error) {
	fake.findOrChooseWorkerForContainerMutex.Lock()
	defer fake.findOrChooseWorkerForContainerMutex.Unlock()
	fake.FindOrChooseWorkerForContainerStub = nil
	if fake.findOrChooseWorkerForContainerReturnsOnCall == nil {
		fake.findOrChooseWorkerForContainerReturnsOnCall = make(map[int]struct {
			result1 worker.Worker
			result2 error
		})
	}
	fake.findOrChooseWorkerForContainerReturnsOnCall[i] = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakePool) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acquireContainerCreatingLockMutex.RLock()
	defer fake.acquireContainerCreatingLockMutex.RUnlock()
	fake.findOrChooseWorkerMutex.RLock()
	defer fake.findOrChooseWorkerMutex.RUnlock()
	fake.findOrChooseWorkerForContainerMutex.RLock()
	defer fake.findOrChooseWorkerForContainerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePool) recordInvocation(key string, args []interface{}) {
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

var _ worker.Pool = new(FakePool)
