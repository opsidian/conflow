// Code generated by counterfeiter. DO NOT EDIT.
package conflowfakes

import (
	"sync"

	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/parsley/parsley"
)

type FakeParameterContainer struct {
	BlockContainerStub        func() conflow.BlockContainer
	blockContainerMutex       sync.RWMutex
	blockContainerArgsForCall []struct {
	}
	blockContainerReturns struct {
		result1 conflow.BlockContainer
	}
	blockContainerReturnsOnCall map[int]struct {
		result1 conflow.BlockContainer
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	NodeStub        func() conflow.Node
	nodeMutex       sync.RWMutex
	nodeArgsForCall []struct {
	}
	nodeReturns struct {
		result1 conflow.Node
	}
	nodeReturnsOnCall map[int]struct {
		result1 conflow.Node
	}
	ValueStub        func() (interface{}, parsley.Error)
	valueMutex       sync.RWMutex
	valueArgsForCall []struct {
	}
	valueReturns struct {
		result1 interface{}
		result2 parsley.Error
	}
	valueReturnsOnCall map[int]struct {
		result1 interface{}
		result2 parsley.Error
	}
	WaitGroupsStub        func() []conflow.WaitGroup
	waitGroupsMutex       sync.RWMutex
	waitGroupsArgsForCall []struct {
	}
	waitGroupsReturns struct {
		result1 []conflow.WaitGroup
	}
	waitGroupsReturnsOnCall map[int]struct {
		result1 []conflow.WaitGroup
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeParameterContainer) BlockContainer() conflow.BlockContainer {
	fake.blockContainerMutex.Lock()
	ret, specificReturn := fake.blockContainerReturnsOnCall[len(fake.blockContainerArgsForCall)]
	fake.blockContainerArgsForCall = append(fake.blockContainerArgsForCall, struct {
	}{})
	stub := fake.BlockContainerStub
	fakeReturns := fake.blockContainerReturns
	fake.recordInvocation("BlockContainer", []interface{}{})
	fake.blockContainerMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParameterContainer) BlockContainerCallCount() int {
	fake.blockContainerMutex.RLock()
	defer fake.blockContainerMutex.RUnlock()
	return len(fake.blockContainerArgsForCall)
}

func (fake *FakeParameterContainer) BlockContainerCalls(stub func() conflow.BlockContainer) {
	fake.blockContainerMutex.Lock()
	defer fake.blockContainerMutex.Unlock()
	fake.BlockContainerStub = stub
}

func (fake *FakeParameterContainer) BlockContainerReturns(result1 conflow.BlockContainer) {
	fake.blockContainerMutex.Lock()
	defer fake.blockContainerMutex.Unlock()
	fake.BlockContainerStub = nil
	fake.blockContainerReturns = struct {
		result1 conflow.BlockContainer
	}{result1}
}

func (fake *FakeParameterContainer) BlockContainerReturnsOnCall(i int, result1 conflow.BlockContainer) {
	fake.blockContainerMutex.Lock()
	defer fake.blockContainerMutex.Unlock()
	fake.BlockContainerStub = nil
	if fake.blockContainerReturnsOnCall == nil {
		fake.blockContainerReturnsOnCall = make(map[int]struct {
			result1 conflow.BlockContainer
		})
	}
	fake.blockContainerReturnsOnCall[i] = struct {
		result1 conflow.BlockContainer
	}{result1}
}

func (fake *FakeParameterContainer) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeParameterContainer) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeParameterContainer) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeParameterContainer) Node() conflow.Node {
	fake.nodeMutex.Lock()
	ret, specificReturn := fake.nodeReturnsOnCall[len(fake.nodeArgsForCall)]
	fake.nodeArgsForCall = append(fake.nodeArgsForCall, struct {
	}{})
	stub := fake.NodeStub
	fakeReturns := fake.nodeReturns
	fake.recordInvocation("Node", []interface{}{})
	fake.nodeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParameterContainer) NodeCallCount() int {
	fake.nodeMutex.RLock()
	defer fake.nodeMutex.RUnlock()
	return len(fake.nodeArgsForCall)
}

func (fake *FakeParameterContainer) NodeCalls(stub func() conflow.Node) {
	fake.nodeMutex.Lock()
	defer fake.nodeMutex.Unlock()
	fake.NodeStub = stub
}

func (fake *FakeParameterContainer) NodeReturns(result1 conflow.Node) {
	fake.nodeMutex.Lock()
	defer fake.nodeMutex.Unlock()
	fake.NodeStub = nil
	fake.nodeReturns = struct {
		result1 conflow.Node
	}{result1}
}

func (fake *FakeParameterContainer) NodeReturnsOnCall(i int, result1 conflow.Node) {
	fake.nodeMutex.Lock()
	defer fake.nodeMutex.Unlock()
	fake.NodeStub = nil
	if fake.nodeReturnsOnCall == nil {
		fake.nodeReturnsOnCall = make(map[int]struct {
			result1 conflow.Node
		})
	}
	fake.nodeReturnsOnCall[i] = struct {
		result1 conflow.Node
	}{result1}
}

func (fake *FakeParameterContainer) Value() (interface{}, parsley.Error) {
	fake.valueMutex.Lock()
	ret, specificReturn := fake.valueReturnsOnCall[len(fake.valueArgsForCall)]
	fake.valueArgsForCall = append(fake.valueArgsForCall, struct {
	}{})
	stub := fake.ValueStub
	fakeReturns := fake.valueReturns
	fake.recordInvocation("Value", []interface{}{})
	fake.valueMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeParameterContainer) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeParameterContainer) ValueCalls(stub func() (interface{}, parsley.Error)) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = stub
}

func (fake *FakeParameterContainer) ValueReturns(result1 interface{}, result2 parsley.Error) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeParameterContainer) ValueReturnsOnCall(i int, result1 interface{}, result2 parsley.Error) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	if fake.valueReturnsOnCall == nil {
		fake.valueReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 parsley.Error
		})
	}
	fake.valueReturnsOnCall[i] = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeParameterContainer) WaitGroups() []conflow.WaitGroup {
	fake.waitGroupsMutex.Lock()
	ret, specificReturn := fake.waitGroupsReturnsOnCall[len(fake.waitGroupsArgsForCall)]
	fake.waitGroupsArgsForCall = append(fake.waitGroupsArgsForCall, struct {
	}{})
	stub := fake.WaitGroupsStub
	fakeReturns := fake.waitGroupsReturns
	fake.recordInvocation("WaitGroups", []interface{}{})
	fake.waitGroupsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeParameterContainer) WaitGroupsCallCount() int {
	fake.waitGroupsMutex.RLock()
	defer fake.waitGroupsMutex.RUnlock()
	return len(fake.waitGroupsArgsForCall)
}

func (fake *FakeParameterContainer) WaitGroupsCalls(stub func() []conflow.WaitGroup) {
	fake.waitGroupsMutex.Lock()
	defer fake.waitGroupsMutex.Unlock()
	fake.WaitGroupsStub = stub
}

func (fake *FakeParameterContainer) WaitGroupsReturns(result1 []conflow.WaitGroup) {
	fake.waitGroupsMutex.Lock()
	defer fake.waitGroupsMutex.Unlock()
	fake.WaitGroupsStub = nil
	fake.waitGroupsReturns = struct {
		result1 []conflow.WaitGroup
	}{result1}
}

func (fake *FakeParameterContainer) WaitGroupsReturnsOnCall(i int, result1 []conflow.WaitGroup) {
	fake.waitGroupsMutex.Lock()
	defer fake.waitGroupsMutex.Unlock()
	fake.WaitGroupsStub = nil
	if fake.waitGroupsReturnsOnCall == nil {
		fake.waitGroupsReturnsOnCall = make(map[int]struct {
			result1 []conflow.WaitGroup
		})
	}
	fake.waitGroupsReturnsOnCall[i] = struct {
		result1 []conflow.WaitGroup
	}{result1}
}

func (fake *FakeParameterContainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blockContainerMutex.RLock()
	defer fake.blockContainerMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.nodeMutex.RLock()
	defer fake.nodeMutex.RUnlock()
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	fake.waitGroupsMutex.RLock()
	defer fake.waitGroupsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeParameterContainer) recordInvocation(key string, args []interface{}) {
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

var _ conflow.ParameterContainer = new(FakeParameterContainer)
