// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
)

type FakeBlockInterpreter struct {
	BlocksStub        func() map[basil.ID]basil.BlockDescriptor
	blocksMutex       sync.RWMutex
	blocksArgsForCall []struct {
	}
	blocksReturns struct {
		result1 map[basil.ID]basil.BlockDescriptor
	}
	blocksReturnsOnCall map[int]struct {
		result1 map[basil.ID]basil.BlockDescriptor
	}
	CreateBlockStub        func(basil.ID) basil.Block
	createBlockMutex       sync.RWMutex
	createBlockArgsForCall []struct {
		arg1 basil.ID
	}
	createBlockReturns struct {
		result1 basil.Block
	}
	createBlockReturnsOnCall map[int]struct {
		result1 basil.Block
	}
	HasForeignIDStub        func() bool
	hasForeignIDMutex       sync.RWMutex
	hasForeignIDArgsForCall []struct {
	}
	hasForeignIDReturns struct {
		result1 bool
	}
	hasForeignIDReturnsOnCall map[int]struct {
		result1 bool
	}
	ParamStub        func(basil.Block, basil.ID) interface{}
	paramMutex       sync.RWMutex
	paramArgsForCall []struct {
		arg1 basil.Block
		arg2 basil.ID
	}
	paramReturns struct {
		result1 interface{}
	}
	paramReturnsOnCall map[int]struct {
		result1 interface{}
	}
	ParamsStub        func() map[basil.ID]basil.ParameterDescriptor
	paramsMutex       sync.RWMutex
	paramsArgsForCall []struct {
	}
	paramsReturns struct {
		result1 map[basil.ID]basil.ParameterDescriptor
	}
	paramsReturnsOnCall map[int]struct {
		result1 map[basil.ID]basil.ParameterDescriptor
	}
	ParseContextStub        func(*basil.ParseContext) *basil.ParseContext
	parseContextMutex       sync.RWMutex
	parseContextArgsForCall []struct {
		arg1 *basil.ParseContext
	}
	parseContextReturns struct {
		result1 *basil.ParseContext
	}
	parseContextReturnsOnCall map[int]struct {
		result1 *basil.ParseContext
	}
	SetBlockStub        func(basil.Block, basil.ID, interface{}) error
	setBlockMutex       sync.RWMutex
	setBlockArgsForCall []struct {
		arg1 basil.Block
		arg2 basil.ID
		arg3 interface{}
	}
	setBlockReturns struct {
		result1 error
	}
	setBlockReturnsOnCall map[int]struct {
		result1 error
	}
	SetParamStub        func(basil.Block, basil.ID, interface{}) error
	setParamMutex       sync.RWMutex
	setParamArgsForCall []struct {
		arg1 basil.Block
		arg2 basil.ID
		arg3 interface{}
	}
	setParamReturns struct {
		result1 error
	}
	setParamReturnsOnCall map[int]struct {
		result1 error
	}
	ValueParamNameStub        func() basil.ID
	valueParamNameMutex       sync.RWMutex
	valueParamNameArgsForCall []struct {
	}
	valueParamNameReturns struct {
		result1 basil.ID
	}
	valueParamNameReturnsOnCall map[int]struct {
		result1 basil.ID
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlockInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	fake.blocksMutex.Lock()
	ret, specificReturn := fake.blocksReturnsOnCall[len(fake.blocksArgsForCall)]
	fake.blocksArgsForCall = append(fake.blocksArgsForCall, struct {
	}{})
	fake.recordInvocation("Blocks", []interface{}{})
	fake.blocksMutex.Unlock()
	if fake.BlocksStub != nil {
		return fake.BlocksStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.blocksReturns
	return fakeReturns.result1
}

func (fake *FakeBlockInterpreter) BlocksCallCount() int {
	fake.blocksMutex.RLock()
	defer fake.blocksMutex.RUnlock()
	return len(fake.blocksArgsForCall)
}

func (fake *FakeBlockInterpreter) BlocksCalls(stub func() map[basil.ID]basil.BlockDescriptor) {
	fake.blocksMutex.Lock()
	defer fake.blocksMutex.Unlock()
	fake.BlocksStub = stub
}

func (fake *FakeBlockInterpreter) BlocksReturns(result1 map[basil.ID]basil.BlockDescriptor) {
	fake.blocksMutex.Lock()
	defer fake.blocksMutex.Unlock()
	fake.BlocksStub = nil
	fake.blocksReturns = struct {
		result1 map[basil.ID]basil.BlockDescriptor
	}{result1}
}

func (fake *FakeBlockInterpreter) BlocksReturnsOnCall(i int, result1 map[basil.ID]basil.BlockDescriptor) {
	fake.blocksMutex.Lock()
	defer fake.blocksMutex.Unlock()
	fake.BlocksStub = nil
	if fake.blocksReturnsOnCall == nil {
		fake.blocksReturnsOnCall = make(map[int]struct {
			result1 map[basil.ID]basil.BlockDescriptor
		})
	}
	fake.blocksReturnsOnCall[i] = struct {
		result1 map[basil.ID]basil.BlockDescriptor
	}{result1}
}

func (fake *FakeBlockInterpreter) CreateBlock(arg1 basil.ID) basil.Block {
	fake.createBlockMutex.Lock()
	ret, specificReturn := fake.createBlockReturnsOnCall[len(fake.createBlockArgsForCall)]
	fake.createBlockArgsForCall = append(fake.createBlockArgsForCall, struct {
		arg1 basil.ID
	}{arg1})
	fake.recordInvocation("CreateBlock", []interface{}{arg1})
	fake.createBlockMutex.Unlock()
	if fake.CreateBlockStub != nil {
		return fake.CreateBlockStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createBlockReturns
	return fakeReturns.result1
}

func (fake *FakeBlockInterpreter) CreateBlockCallCount() int {
	fake.createBlockMutex.RLock()
	defer fake.createBlockMutex.RUnlock()
	return len(fake.createBlockArgsForCall)
}

func (fake *FakeBlockInterpreter) CreateBlockCalls(stub func(basil.ID) basil.Block) {
	fake.createBlockMutex.Lock()
	defer fake.createBlockMutex.Unlock()
	fake.CreateBlockStub = stub
}

func (fake *FakeBlockInterpreter) CreateBlockArgsForCall(i int) basil.ID {
	fake.createBlockMutex.RLock()
	defer fake.createBlockMutex.RUnlock()
	argsForCall := fake.createBlockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockInterpreter) CreateBlockReturns(result1 basil.Block) {
	fake.createBlockMutex.Lock()
	defer fake.createBlockMutex.Unlock()
	fake.CreateBlockStub = nil
	fake.createBlockReturns = struct {
		result1 basil.Block
	}{result1}
}

func (fake *FakeBlockInterpreter) CreateBlockReturnsOnCall(i int, result1 basil.Block) {
	fake.createBlockMutex.Lock()
	defer fake.createBlockMutex.Unlock()
	fake.CreateBlockStub = nil
	if fake.createBlockReturnsOnCall == nil {
		fake.createBlockReturnsOnCall = make(map[int]struct {
			result1 basil.Block
		})
	}
	fake.createBlockReturnsOnCall[i] = struct {
		result1 basil.Block
	}{result1}
}

func (fake *FakeBlockInterpreter) HasForeignID() bool {
	fake.hasForeignIDMutex.Lock()
	ret, specificReturn := fake.hasForeignIDReturnsOnCall[len(fake.hasForeignIDArgsForCall)]
	fake.hasForeignIDArgsForCall = append(fake.hasForeignIDArgsForCall, struct {
	}{})
	fake.recordInvocation("HasForeignID", []interface{}{})
	fake.hasForeignIDMutex.Unlock()
	if fake.HasForeignIDStub != nil {
		return fake.HasForeignIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.hasForeignIDReturns
	return fakeReturns.result1
}

func (fake *FakeBlockInterpreter) HasForeignIDCallCount() int {
	fake.hasForeignIDMutex.RLock()
	defer fake.hasForeignIDMutex.RUnlock()
	return len(fake.hasForeignIDArgsForCall)
}

func (fake *FakeBlockInterpreter) HasForeignIDCalls(stub func() bool) {
	fake.hasForeignIDMutex.Lock()
	defer fake.hasForeignIDMutex.Unlock()
	fake.HasForeignIDStub = stub
}

func (fake *FakeBlockInterpreter) HasForeignIDReturns(result1 bool) {
	fake.hasForeignIDMutex.Lock()
	defer fake.hasForeignIDMutex.Unlock()
	fake.HasForeignIDStub = nil
	fake.hasForeignIDReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBlockInterpreter) HasForeignIDReturnsOnCall(i int, result1 bool) {
	fake.hasForeignIDMutex.Lock()
	defer fake.hasForeignIDMutex.Unlock()
	fake.HasForeignIDStub = nil
	if fake.hasForeignIDReturnsOnCall == nil {
		fake.hasForeignIDReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasForeignIDReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBlockInterpreter) Param(arg1 basil.Block, arg2 basil.ID) interface{} {
	fake.paramMutex.Lock()
	ret, specificReturn := fake.paramReturnsOnCall[len(fake.paramArgsForCall)]
	fake.paramArgsForCall = append(fake.paramArgsForCall, struct {
		arg1 basil.Block
		arg2 basil.ID
	}{arg1, arg2})
	fake.recordInvocation("Param", []interface{}{arg1, arg2})
	fake.paramMutex.Unlock()
	if fake.ParamStub != nil {
		return fake.ParamStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.paramReturns
	return fakeReturns.result1
}

func (fake *FakeBlockInterpreter) ParamCallCount() int {
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	return len(fake.paramArgsForCall)
}

func (fake *FakeBlockInterpreter) ParamCalls(stub func(basil.Block, basil.ID) interface{}) {
	fake.paramMutex.Lock()
	defer fake.paramMutex.Unlock()
	fake.ParamStub = stub
}

func (fake *FakeBlockInterpreter) ParamArgsForCall(i int) (basil.Block, basil.ID) {
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	argsForCall := fake.paramArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeBlockInterpreter) ParamReturns(result1 interface{}) {
	fake.paramMutex.Lock()
	defer fake.paramMutex.Unlock()
	fake.ParamStub = nil
	fake.paramReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeBlockInterpreter) ParamReturnsOnCall(i int, result1 interface{}) {
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

func (fake *FakeBlockInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	fake.paramsMutex.Lock()
	ret, specificReturn := fake.paramsReturnsOnCall[len(fake.paramsArgsForCall)]
	fake.paramsArgsForCall = append(fake.paramsArgsForCall, struct {
	}{})
	fake.recordInvocation("Params", []interface{}{})
	fake.paramsMutex.Unlock()
	if fake.ParamsStub != nil {
		return fake.ParamsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.paramsReturns
	return fakeReturns.result1
}

func (fake *FakeBlockInterpreter) ParamsCallCount() int {
	fake.paramsMutex.RLock()
	defer fake.paramsMutex.RUnlock()
	return len(fake.paramsArgsForCall)
}

func (fake *FakeBlockInterpreter) ParamsCalls(stub func() map[basil.ID]basil.ParameterDescriptor) {
	fake.paramsMutex.Lock()
	defer fake.paramsMutex.Unlock()
	fake.ParamsStub = stub
}

func (fake *FakeBlockInterpreter) ParamsReturns(result1 map[basil.ID]basil.ParameterDescriptor) {
	fake.paramsMutex.Lock()
	defer fake.paramsMutex.Unlock()
	fake.ParamsStub = nil
	fake.paramsReturns = struct {
		result1 map[basil.ID]basil.ParameterDescriptor
	}{result1}
}

func (fake *FakeBlockInterpreter) ParamsReturnsOnCall(i int, result1 map[basil.ID]basil.ParameterDescriptor) {
	fake.paramsMutex.Lock()
	defer fake.paramsMutex.Unlock()
	fake.ParamsStub = nil
	if fake.paramsReturnsOnCall == nil {
		fake.paramsReturnsOnCall = make(map[int]struct {
			result1 map[basil.ID]basil.ParameterDescriptor
		})
	}
	fake.paramsReturnsOnCall[i] = struct {
		result1 map[basil.ID]basil.ParameterDescriptor
	}{result1}
}

func (fake *FakeBlockInterpreter) ParseContext(arg1 *basil.ParseContext) *basil.ParseContext {
	fake.parseContextMutex.Lock()
	ret, specificReturn := fake.parseContextReturnsOnCall[len(fake.parseContextArgsForCall)]
	fake.parseContextArgsForCall = append(fake.parseContextArgsForCall, struct {
		arg1 *basil.ParseContext
	}{arg1})
	fake.recordInvocation("ParseContext", []interface{}{arg1})
	fake.parseContextMutex.Unlock()
	if fake.ParseContextStub != nil {
		return fake.ParseContextStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.parseContextReturns
	return fakeReturns.result1
}

func (fake *FakeBlockInterpreter) ParseContextCallCount() int {
	fake.parseContextMutex.RLock()
	defer fake.parseContextMutex.RUnlock()
	return len(fake.parseContextArgsForCall)
}

func (fake *FakeBlockInterpreter) ParseContextCalls(stub func(*basil.ParseContext) *basil.ParseContext) {
	fake.parseContextMutex.Lock()
	defer fake.parseContextMutex.Unlock()
	fake.ParseContextStub = stub
}

func (fake *FakeBlockInterpreter) ParseContextArgsForCall(i int) *basil.ParseContext {
	fake.parseContextMutex.RLock()
	defer fake.parseContextMutex.RUnlock()
	argsForCall := fake.parseContextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockInterpreter) ParseContextReturns(result1 *basil.ParseContext) {
	fake.parseContextMutex.Lock()
	defer fake.parseContextMutex.Unlock()
	fake.ParseContextStub = nil
	fake.parseContextReturns = struct {
		result1 *basil.ParseContext
	}{result1}
}

func (fake *FakeBlockInterpreter) ParseContextReturnsOnCall(i int, result1 *basil.ParseContext) {
	fake.parseContextMutex.Lock()
	defer fake.parseContextMutex.Unlock()
	fake.ParseContextStub = nil
	if fake.parseContextReturnsOnCall == nil {
		fake.parseContextReturnsOnCall = make(map[int]struct {
			result1 *basil.ParseContext
		})
	}
	fake.parseContextReturnsOnCall[i] = struct {
		result1 *basil.ParseContext
	}{result1}
}

func (fake *FakeBlockInterpreter) SetBlock(arg1 basil.Block, arg2 basil.ID, arg3 interface{}) error {
	fake.setBlockMutex.Lock()
	ret, specificReturn := fake.setBlockReturnsOnCall[len(fake.setBlockArgsForCall)]
	fake.setBlockArgsForCall = append(fake.setBlockArgsForCall, struct {
		arg1 basil.Block
		arg2 basil.ID
		arg3 interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetBlock", []interface{}{arg1, arg2, arg3})
	fake.setBlockMutex.Unlock()
	if fake.SetBlockStub != nil {
		return fake.SetBlockStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setBlockReturns
	return fakeReturns.result1
}

func (fake *FakeBlockInterpreter) SetBlockCallCount() int {
	fake.setBlockMutex.RLock()
	defer fake.setBlockMutex.RUnlock()
	return len(fake.setBlockArgsForCall)
}

func (fake *FakeBlockInterpreter) SetBlockCalls(stub func(basil.Block, basil.ID, interface{}) error) {
	fake.setBlockMutex.Lock()
	defer fake.setBlockMutex.Unlock()
	fake.SetBlockStub = stub
}

func (fake *FakeBlockInterpreter) SetBlockArgsForCall(i int) (basil.Block, basil.ID, interface{}) {
	fake.setBlockMutex.RLock()
	defer fake.setBlockMutex.RUnlock()
	argsForCall := fake.setBlockArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBlockInterpreter) SetBlockReturns(result1 error) {
	fake.setBlockMutex.Lock()
	defer fake.setBlockMutex.Unlock()
	fake.SetBlockStub = nil
	fake.setBlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlockInterpreter) SetBlockReturnsOnCall(i int, result1 error) {
	fake.setBlockMutex.Lock()
	defer fake.setBlockMutex.Unlock()
	fake.SetBlockStub = nil
	if fake.setBlockReturnsOnCall == nil {
		fake.setBlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setBlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlockInterpreter) SetParam(arg1 basil.Block, arg2 basil.ID, arg3 interface{}) error {
	fake.setParamMutex.Lock()
	ret, specificReturn := fake.setParamReturnsOnCall[len(fake.setParamArgsForCall)]
	fake.setParamArgsForCall = append(fake.setParamArgsForCall, struct {
		arg1 basil.Block
		arg2 basil.ID
		arg3 interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("SetParam", []interface{}{arg1, arg2, arg3})
	fake.setParamMutex.Unlock()
	if fake.SetParamStub != nil {
		return fake.SetParamStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setParamReturns
	return fakeReturns.result1
}

func (fake *FakeBlockInterpreter) SetParamCallCount() int {
	fake.setParamMutex.RLock()
	defer fake.setParamMutex.RUnlock()
	return len(fake.setParamArgsForCall)
}

func (fake *FakeBlockInterpreter) SetParamCalls(stub func(basil.Block, basil.ID, interface{}) error) {
	fake.setParamMutex.Lock()
	defer fake.setParamMutex.Unlock()
	fake.SetParamStub = stub
}

func (fake *FakeBlockInterpreter) SetParamArgsForCall(i int) (basil.Block, basil.ID, interface{}) {
	fake.setParamMutex.RLock()
	defer fake.setParamMutex.RUnlock()
	argsForCall := fake.setParamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBlockInterpreter) SetParamReturns(result1 error) {
	fake.setParamMutex.Lock()
	defer fake.setParamMutex.Unlock()
	fake.SetParamStub = nil
	fake.setParamReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlockInterpreter) SetParamReturnsOnCall(i int, result1 error) {
	fake.setParamMutex.Lock()
	defer fake.setParamMutex.Unlock()
	fake.SetParamStub = nil
	if fake.setParamReturnsOnCall == nil {
		fake.setParamReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setParamReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlockInterpreter) ValueParamName() basil.ID {
	fake.valueParamNameMutex.Lock()
	ret, specificReturn := fake.valueParamNameReturnsOnCall[len(fake.valueParamNameArgsForCall)]
	fake.valueParamNameArgsForCall = append(fake.valueParamNameArgsForCall, struct {
	}{})
	fake.recordInvocation("ValueParamName", []interface{}{})
	fake.valueParamNameMutex.Unlock()
	if fake.ValueParamNameStub != nil {
		return fake.ValueParamNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.valueParamNameReturns
	return fakeReturns.result1
}

func (fake *FakeBlockInterpreter) ValueParamNameCallCount() int {
	fake.valueParamNameMutex.RLock()
	defer fake.valueParamNameMutex.RUnlock()
	return len(fake.valueParamNameArgsForCall)
}

func (fake *FakeBlockInterpreter) ValueParamNameCalls(stub func() basil.ID) {
	fake.valueParamNameMutex.Lock()
	defer fake.valueParamNameMutex.Unlock()
	fake.ValueParamNameStub = stub
}

func (fake *FakeBlockInterpreter) ValueParamNameReturns(result1 basil.ID) {
	fake.valueParamNameMutex.Lock()
	defer fake.valueParamNameMutex.Unlock()
	fake.ValueParamNameStub = nil
	fake.valueParamNameReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeBlockInterpreter) ValueParamNameReturnsOnCall(i int, result1 basil.ID) {
	fake.valueParamNameMutex.Lock()
	defer fake.valueParamNameMutex.Unlock()
	fake.ValueParamNameStub = nil
	if fake.valueParamNameReturnsOnCall == nil {
		fake.valueParamNameReturnsOnCall = make(map[int]struct {
			result1 basil.ID
		})
	}
	fake.valueParamNameReturnsOnCall[i] = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeBlockInterpreter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blocksMutex.RLock()
	defer fake.blocksMutex.RUnlock()
	fake.createBlockMutex.RLock()
	defer fake.createBlockMutex.RUnlock()
	fake.hasForeignIDMutex.RLock()
	defer fake.hasForeignIDMutex.RUnlock()
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	fake.paramsMutex.RLock()
	defer fake.paramsMutex.RUnlock()
	fake.parseContextMutex.RLock()
	defer fake.parseContextMutex.RUnlock()
	fake.setBlockMutex.RLock()
	defer fake.setBlockMutex.RUnlock()
	fake.setParamMutex.RLock()
	defer fake.setParamMutex.RUnlock()
	fake.valueParamNameMutex.RLock()
	defer fake.valueParamNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlockInterpreter) recordInvocation(key string, args []interface{}) {
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

var _ basil.BlockInterpreter = new(FakeBlockInterpreter)
