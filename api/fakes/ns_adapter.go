// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/containernetworking/plugins/pkg/ns"
)

type NSAdapter struct {
	GetNSStub        func(nspath string) (ns.NetNS, error)
	getNSMutex       sync.RWMutex
	getNSArgsForCall []struct {
		nspath string
	}
	getNSReturns struct {
		result1 ns.NetNS
		result2 error
	}
	getNSReturnsOnCall map[int]struct {
		result1 ns.NetNS
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NSAdapter) GetNS(nspath string) (ns.NetNS, error) {
	fake.getNSMutex.Lock()
	ret, specificReturn := fake.getNSReturnsOnCall[len(fake.getNSArgsForCall)]
	fake.getNSArgsForCall = append(fake.getNSArgsForCall, struct {
		nspath string
	}{nspath})
	fake.recordInvocation("GetNS", []interface{}{nspath})
	fake.getNSMutex.Unlock()
	if fake.GetNSStub != nil {
		return fake.GetNSStub(nspath)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getNSReturns.result1, fake.getNSReturns.result2
}

func (fake *NSAdapter) GetNSCallCount() int {
	fake.getNSMutex.RLock()
	defer fake.getNSMutex.RUnlock()
	return len(fake.getNSArgsForCall)
}

func (fake *NSAdapter) GetNSArgsForCall(i int) string {
	fake.getNSMutex.RLock()
	defer fake.getNSMutex.RUnlock()
	return fake.getNSArgsForCall[i].nspath
}

func (fake *NSAdapter) GetNSReturns(result1 ns.NetNS, result2 error) {
	fake.GetNSStub = nil
	fake.getNSReturns = struct {
		result1 ns.NetNS
		result2 error
	}{result1, result2}
}

func (fake *NSAdapter) GetNSReturnsOnCall(i int, result1 ns.NetNS, result2 error) {
	fake.GetNSStub = nil
	if fake.getNSReturnsOnCall == nil {
		fake.getNSReturnsOnCall = make(map[int]struct {
			result1 ns.NetNS
			result2 error
		})
	}
	fake.getNSReturnsOnCall[i] = struct {
		result1 ns.NetNS
		result2 error
	}{result1, result2}
}

func (fake *NSAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getNSMutex.RLock()
	defer fake.getNSMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *NSAdapter) recordInvocation(key string, args []interface{}) {
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
