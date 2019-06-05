// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/parsley"
)

type FakeBlockContainer struct {
	BlockStub        func() basil.Block
	blockMutex       sync.RWMutex
	blockArgsForCall []struct {
	}
	blockReturns struct {
		result1 basil.Block
	}
	blockReturnsOnCall map[int]struct {
		result1 basil.Block
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	IDStub        func() basil.ID
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 basil.ID
	}
	iDReturnsOnCall map[int]struct {
		result1 basil.ID
	}
	NodeStub        func() basil.Node
	nodeMutex       sync.RWMutex
	nodeArgsForCall []struct {
	}
	nodeReturns struct {
		result1 basil.Node
	}
	nodeReturnsOnCall map[int]struct {
		result1 basil.Node
	}
	ParamStub        func(basil.ID) interface{}
	paramMutex       sync.RWMutex
	paramArgsForCall []struct {
		arg1 basil.ID
	}
	paramReturns struct {
		result1 interface{}
	}
	paramReturnsOnCall map[int]struct {
		result1 interface{}
	}
	PublishBlockStub        func(basil.ID, basil.BlockMessage)
	publishBlockMutex       sync.RWMutex
	publishBlockArgsForCall []struct {
		arg1 basil.ID
		arg2 basil.BlockMessage
	}
	RunStub        func()
	runMutex       sync.RWMutex
	runArgsForCall []struct {
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
	WaitGroupsStub        func() []*sync.WaitGroup
	waitGroupsMutex       sync.RWMutex
	waitGroupsArgsForCall []struct {
	}
	waitGroupsReturns struct {
		result1 []*sync.WaitGroup
	}
	waitGroupsReturnsOnCall map[int]struct {
		result1 []*sync.WaitGroup
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlockContainer) Block() basil.Block {
	fake.blockMutex.Lock()
	ret, specificReturn := fake.blockReturnsOnCall[len(fake.blockArgsForCall)]
	fake.blockArgsForCall = append(fake.blockArgsForCall, struct {
	}{})
	fake.recordInvocation("Block", []interface{}{})
	fake.blockMutex.Unlock()
	if fake.BlockStub != nil {
		return fake.BlockStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.blockReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContainer) BlockCallCount() int {
	fake.blockMutex.RLock()
	defer fake.blockMutex.RUnlock()
	return len(fake.blockArgsForCall)
}

func (fake *FakeBlockContainer) BlockCalls(stub func() basil.Block) {
	fake.blockMutex.Lock()
	defer fake.blockMutex.Unlock()
	fake.BlockStub = stub
}

func (fake *FakeBlockContainer) BlockReturns(result1 basil.Block) {
	fake.blockMutex.Lock()
	defer fake.blockMutex.Unlock()
	fake.BlockStub = nil
	fake.blockReturns = struct {
		result1 basil.Block
	}{result1}
}

func (fake *FakeBlockContainer) BlockReturnsOnCall(i int, result1 basil.Block) {
	fake.blockMutex.Lock()
	defer fake.blockMutex.Unlock()
	fake.BlockStub = nil
	if fake.blockReturnsOnCall == nil {
		fake.blockReturnsOnCall = make(map[int]struct {
			result1 basil.Block
		})
	}
	fake.blockReturnsOnCall[i] = struct {
		result1 basil.Block
	}{result1}
}

func (fake *FakeBlockContainer) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeBlockContainer) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeBlockContainer) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeBlockContainer) ID() basil.ID {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContainer) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeBlockContainer) IDCalls(stub func() basil.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeBlockContainer) IDReturns(result1 basil.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeBlockContainer) IDReturnsOnCall(i int, result1 basil.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 basil.ID
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeBlockContainer) Node() basil.Node {
	fake.nodeMutex.Lock()
	ret, specificReturn := fake.nodeReturnsOnCall[len(fake.nodeArgsForCall)]
	fake.nodeArgsForCall = append(fake.nodeArgsForCall, struct {
	}{})
	fake.recordInvocation("Node", []interface{}{})
	fake.nodeMutex.Unlock()
	if fake.NodeStub != nil {
		return fake.NodeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nodeReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContainer) NodeCallCount() int {
	fake.nodeMutex.RLock()
	defer fake.nodeMutex.RUnlock()
	return len(fake.nodeArgsForCall)
}

func (fake *FakeBlockContainer) NodeCalls(stub func() basil.Node) {
	fake.nodeMutex.Lock()
	defer fake.nodeMutex.Unlock()
	fake.NodeStub = stub
}

func (fake *FakeBlockContainer) NodeReturns(result1 basil.Node) {
	fake.nodeMutex.Lock()
	defer fake.nodeMutex.Unlock()
	fake.NodeStub = nil
	fake.nodeReturns = struct {
		result1 basil.Node
	}{result1}
}

func (fake *FakeBlockContainer) NodeReturnsOnCall(i int, result1 basil.Node) {
	fake.nodeMutex.Lock()
	defer fake.nodeMutex.Unlock()
	fake.NodeStub = nil
	if fake.nodeReturnsOnCall == nil {
		fake.nodeReturnsOnCall = make(map[int]struct {
			result1 basil.Node
		})
	}
	fake.nodeReturnsOnCall[i] = struct {
		result1 basil.Node
	}{result1}
}

func (fake *FakeBlockContainer) Param(arg1 basil.ID) interface{} {
	fake.paramMutex.Lock()
	ret, specificReturn := fake.paramReturnsOnCall[len(fake.paramArgsForCall)]
	fake.paramArgsForCall = append(fake.paramArgsForCall, struct {
		arg1 basil.ID
	}{arg1})
	fake.recordInvocation("Param", []interface{}{arg1})
	fake.paramMutex.Unlock()
	if fake.ParamStub != nil {
		return fake.ParamStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.paramReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContainer) ParamCallCount() int {
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	return len(fake.paramArgsForCall)
}

func (fake *FakeBlockContainer) ParamCalls(stub func(basil.ID) interface{}) {
	fake.paramMutex.Lock()
	defer fake.paramMutex.Unlock()
	fake.ParamStub = stub
}

func (fake *FakeBlockContainer) ParamArgsForCall(i int) basil.ID {
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	argsForCall := fake.paramArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockContainer) ParamReturns(result1 interface{}) {
	fake.paramMutex.Lock()
	defer fake.paramMutex.Unlock()
	fake.ParamStub = nil
	fake.paramReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeBlockContainer) ParamReturnsOnCall(i int, result1 interface{}) {
	fake.paramMutex.Lock()
	defer fake.paramMutex.Unlock()
	fake.ParamStub = nil
	if fake.paramReturnsOnCall == nil {
		fake.paramReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.paramReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeBlockContainer) PublishBlock(arg1 basil.ID, arg2 basil.BlockMessage) {
	fake.publishBlockMutex.Lock()
	fake.publishBlockArgsForCall = append(fake.publishBlockArgsForCall, struct {
		arg1 basil.ID
		arg2 basil.BlockMessage
	}{arg1, arg2})
	fake.recordInvocation("PublishBlock", []interface{}{arg1, arg2})
	fake.publishBlockMutex.Unlock()
	if fake.PublishBlockStub != nil {
		fake.PublishBlockStub(arg1, arg2)
	}
}

func (fake *FakeBlockContainer) PublishBlockCallCount() int {
	fake.publishBlockMutex.RLock()
	defer fake.publishBlockMutex.RUnlock()
	return len(fake.publishBlockArgsForCall)
}

func (fake *FakeBlockContainer) PublishBlockCalls(stub func(basil.ID, basil.BlockMessage)) {
	fake.publishBlockMutex.Lock()
	defer fake.publishBlockMutex.Unlock()
	fake.PublishBlockStub = stub
}

func (fake *FakeBlockContainer) PublishBlockArgsForCall(i int) (basil.ID, basil.BlockMessage) {
	fake.publishBlockMutex.RLock()
	defer fake.publishBlockMutex.RUnlock()
	argsForCall := fake.publishBlockArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBlockContainer) Run() {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
	}{})
	fake.recordInvocation("Run", []interface{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		fake.RunStub()
	}
}

func (fake *FakeBlockContainer) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeBlockContainer) RunCalls(stub func()) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeBlockContainer) Value() (interface{}, parsley.Error) {
	fake.valueMutex.Lock()
	ret, specificReturn := fake.valueReturnsOnCall[len(fake.valueArgsForCall)]
	fake.valueArgsForCall = append(fake.valueArgsForCall, struct {
	}{})
	fake.recordInvocation("Value", []interface{}{})
	fake.valueMutex.Unlock()
	if fake.ValueStub != nil {
		return fake.ValueStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.valueReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlockContainer) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeBlockContainer) ValueCalls(stub func() (interface{}, parsley.Error)) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = stub
}

func (fake *FakeBlockContainer) ValueReturns(result1 interface{}, result2 parsley.Error) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeBlockContainer) ValueReturnsOnCall(i int, result1 interface{}, result2 parsley.Error) {
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

func (fake *FakeBlockContainer) WaitGroups() []*sync.WaitGroup {
	fake.waitGroupsMutex.Lock()
	ret, specificReturn := fake.waitGroupsReturnsOnCall[len(fake.waitGroupsArgsForCall)]
	fake.waitGroupsArgsForCall = append(fake.waitGroupsArgsForCall, struct {
	}{})
	fake.recordInvocation("WaitGroups", []interface{}{})
	fake.waitGroupsMutex.Unlock()
	if fake.WaitGroupsStub != nil {
		return fake.WaitGroupsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.waitGroupsReturns
	return fakeReturns.result1
}

func (fake *FakeBlockContainer) WaitGroupsCallCount() int {
	fake.waitGroupsMutex.RLock()
	defer fake.waitGroupsMutex.RUnlock()
	return len(fake.waitGroupsArgsForCall)
}

func (fake *FakeBlockContainer) WaitGroupsCalls(stub func() []*sync.WaitGroup) {
	fake.waitGroupsMutex.Lock()
	defer fake.waitGroupsMutex.Unlock()
	fake.WaitGroupsStub = stub
}

func (fake *FakeBlockContainer) WaitGroupsReturns(result1 []*sync.WaitGroup) {
	fake.waitGroupsMutex.Lock()
	defer fake.waitGroupsMutex.Unlock()
	fake.WaitGroupsStub = nil
	fake.waitGroupsReturns = struct {
		result1 []*sync.WaitGroup
	}{result1}
}

func (fake *FakeBlockContainer) WaitGroupsReturnsOnCall(i int, result1 []*sync.WaitGroup) {
	fake.waitGroupsMutex.Lock()
	defer fake.waitGroupsMutex.Unlock()
	fake.WaitGroupsStub = nil
	if fake.waitGroupsReturnsOnCall == nil {
		fake.waitGroupsReturnsOnCall = make(map[int]struct {
			result1 []*sync.WaitGroup
		})
	}
	fake.waitGroupsReturnsOnCall[i] = struct {
		result1 []*sync.WaitGroup
	}{result1}
}

func (fake *FakeBlockContainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blockMutex.RLock()
	defer fake.blockMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.nodeMutex.RLock()
	defer fake.nodeMutex.RUnlock()
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	fake.publishBlockMutex.RLock()
	defer fake.publishBlockMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
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

func (fake *FakeBlockContainer) recordInvocation(key string, args []interface{}) {
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

var _ basil.BlockContainer = new(FakeBlockContainer)
