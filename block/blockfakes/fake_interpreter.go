// Code generated by counterfeiter. DO NOT EDIT.
package blockfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/block"
	"github.com/opsidian/parsley/parsley"
)

type FakeInterpreter struct {
	CreateStub        func(*basil.EvalContext, basil.BlockNode) basil.Block
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 *basil.EvalContext
		arg2 basil.BlockNode
	}
	createReturns struct {
		result1 basil.Block
	}
	createReturnsOnCall map[int]struct {
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
	ParamsStub        func() map[basil.ID]block.ParameterDescriptor
	paramsMutex       sync.RWMutex
	paramsArgsForCall []struct {
	}
	paramsReturns struct {
		result1 map[basil.ID]block.ParameterDescriptor
	}
	paramsReturnsOnCall map[int]struct {
		result1 map[basil.ID]block.ParameterDescriptor
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
	SetBlockStub        func(*basil.EvalContext, basil.Block, basil.ID, interface{}) parsley.Error
	setBlockMutex       sync.RWMutex
	setBlockArgsForCall []struct {
		arg1 *basil.EvalContext
		arg2 basil.Block
		arg3 basil.ID
		arg4 interface{}
	}
	setBlockReturns struct {
		result1 parsley.Error
	}
	setBlockReturnsOnCall map[int]struct {
		result1 parsley.Error
	}
	SetParamStub        func(*basil.EvalContext, basil.Block, basil.ID, basil.BlockParamNode) parsley.Error
	setParamMutex       sync.RWMutex
	setParamArgsForCall []struct {
		arg1 *basil.EvalContext
		arg2 basil.Block
		arg3 basil.ID
		arg4 basil.BlockParamNode
	}
	setParamReturns struct {
		result1 parsley.Error
	}
	setParamReturnsOnCall map[int]struct {
		result1 parsley.Error
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

func (fake *FakeInterpreter) Create(arg1 *basil.EvalContext, arg2 basil.BlockNode) basil.Block {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 *basil.EvalContext
		arg2 basil.BlockNode
	}{arg1, arg2})
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakeInterpreter) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeInterpreter) CreateCalls(stub func(*basil.EvalContext, basil.BlockNode) basil.Block) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeInterpreter) CreateArgsForCall(i int) (*basil.EvalContext, basil.BlockNode) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInterpreter) CreateReturns(result1 basil.Block) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 basil.Block
	}{result1}
}

func (fake *FakeInterpreter) CreateReturnsOnCall(i int, result1 basil.Block) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 basil.Block
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 basil.Block
	}{result1}
}

func (fake *FakeInterpreter) HasForeignID() bool {
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

func (fake *FakeInterpreter) HasForeignIDCallCount() int {
	fake.hasForeignIDMutex.RLock()
	defer fake.hasForeignIDMutex.RUnlock()
	return len(fake.hasForeignIDArgsForCall)
}

func (fake *FakeInterpreter) HasForeignIDCalls(stub func() bool) {
	fake.hasForeignIDMutex.Lock()
	defer fake.hasForeignIDMutex.Unlock()
	fake.HasForeignIDStub = stub
}

func (fake *FakeInterpreter) HasForeignIDReturns(result1 bool) {
	fake.hasForeignIDMutex.Lock()
	defer fake.hasForeignIDMutex.Unlock()
	fake.HasForeignIDStub = nil
	fake.hasForeignIDReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeInterpreter) HasForeignIDReturnsOnCall(i int, result1 bool) {
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

func (fake *FakeInterpreter) Param(arg1 basil.Block, arg2 basil.ID) interface{} {
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

func (fake *FakeInterpreter) ParamCallCount() int {
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	return len(fake.paramArgsForCall)
}

func (fake *FakeInterpreter) ParamCalls(stub func(basil.Block, basil.ID) interface{}) {
	fake.paramMutex.Lock()
	defer fake.paramMutex.Unlock()
	fake.ParamStub = stub
}

func (fake *FakeInterpreter) ParamArgsForCall(i int) (basil.Block, basil.ID) {
	fake.paramMutex.RLock()
	defer fake.paramMutex.RUnlock()
	argsForCall := fake.paramArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInterpreter) ParamReturns(result1 interface{}) {
	fake.paramMutex.Lock()
	defer fake.paramMutex.Unlock()
	fake.ParamStub = nil
	fake.paramReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeInterpreter) ParamReturnsOnCall(i int, result1 interface{}) {
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

func (fake *FakeInterpreter) Params() map[basil.ID]block.ParameterDescriptor {
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

func (fake *FakeInterpreter) ParamsCallCount() int {
	fake.paramsMutex.RLock()
	defer fake.paramsMutex.RUnlock()
	return len(fake.paramsArgsForCall)
}

func (fake *FakeInterpreter) ParamsCalls(stub func() map[basil.ID]block.ParameterDescriptor) {
	fake.paramsMutex.Lock()
	defer fake.paramsMutex.Unlock()
	fake.ParamsStub = stub
}

func (fake *FakeInterpreter) ParamsReturns(result1 map[basil.ID]block.ParameterDescriptor) {
	fake.paramsMutex.Lock()
	defer fake.paramsMutex.Unlock()
	fake.ParamsStub = nil
	fake.paramsReturns = struct {
		result1 map[basil.ID]block.ParameterDescriptor
	}{result1}
}

func (fake *FakeInterpreter) ParamsReturnsOnCall(i int, result1 map[basil.ID]block.ParameterDescriptor) {
	fake.paramsMutex.Lock()
	defer fake.paramsMutex.Unlock()
	fake.ParamsStub = nil
	if fake.paramsReturnsOnCall == nil {
		fake.paramsReturnsOnCall = make(map[int]struct {
			result1 map[basil.ID]block.ParameterDescriptor
		})
	}
	fake.paramsReturnsOnCall[i] = struct {
		result1 map[basil.ID]block.ParameterDescriptor
	}{result1}
}

func (fake *FakeInterpreter) ParseContext(arg1 *basil.ParseContext) *basil.ParseContext {
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

func (fake *FakeInterpreter) ParseContextCallCount() int {
	fake.parseContextMutex.RLock()
	defer fake.parseContextMutex.RUnlock()
	return len(fake.parseContextArgsForCall)
}

func (fake *FakeInterpreter) ParseContextCalls(stub func(*basil.ParseContext) *basil.ParseContext) {
	fake.parseContextMutex.Lock()
	defer fake.parseContextMutex.Unlock()
	fake.ParseContextStub = stub
}

func (fake *FakeInterpreter) ParseContextArgsForCall(i int) *basil.ParseContext {
	fake.parseContextMutex.RLock()
	defer fake.parseContextMutex.RUnlock()
	argsForCall := fake.parseContextArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInterpreter) ParseContextReturns(result1 *basil.ParseContext) {
	fake.parseContextMutex.Lock()
	defer fake.parseContextMutex.Unlock()
	fake.ParseContextStub = nil
	fake.parseContextReturns = struct {
		result1 *basil.ParseContext
	}{result1}
}

func (fake *FakeInterpreter) ParseContextReturnsOnCall(i int, result1 *basil.ParseContext) {
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

func (fake *FakeInterpreter) SetBlock(arg1 *basil.EvalContext, arg2 basil.Block, arg3 basil.ID, arg4 interface{}) parsley.Error {
	fake.setBlockMutex.Lock()
	ret, specificReturn := fake.setBlockReturnsOnCall[len(fake.setBlockArgsForCall)]
	fake.setBlockArgsForCall = append(fake.setBlockArgsForCall, struct {
		arg1 *basil.EvalContext
		arg2 basil.Block
		arg3 basil.ID
		arg4 interface{}
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetBlock", []interface{}{arg1, arg2, arg3, arg4})
	fake.setBlockMutex.Unlock()
	if fake.SetBlockStub != nil {
		return fake.SetBlockStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setBlockReturns
	return fakeReturns.result1
}

func (fake *FakeInterpreter) SetBlockCallCount() int {
	fake.setBlockMutex.RLock()
	defer fake.setBlockMutex.RUnlock()
	return len(fake.setBlockArgsForCall)
}

func (fake *FakeInterpreter) SetBlockCalls(stub func(*basil.EvalContext, basil.Block, basil.ID, interface{}) parsley.Error) {
	fake.setBlockMutex.Lock()
	defer fake.setBlockMutex.Unlock()
	fake.SetBlockStub = stub
}

func (fake *FakeInterpreter) SetBlockArgsForCall(i int) (*basil.EvalContext, basil.Block, basil.ID, interface{}) {
	fake.setBlockMutex.RLock()
	defer fake.setBlockMutex.RUnlock()
	argsForCall := fake.setBlockArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeInterpreter) SetBlockReturns(result1 parsley.Error) {
	fake.setBlockMutex.Lock()
	defer fake.setBlockMutex.Unlock()
	fake.SetBlockStub = nil
	fake.setBlockReturns = struct {
		result1 parsley.Error
	}{result1}
}

func (fake *FakeInterpreter) SetBlockReturnsOnCall(i int, result1 parsley.Error) {
	fake.setBlockMutex.Lock()
	defer fake.setBlockMutex.Unlock()
	fake.SetBlockStub = nil
	if fake.setBlockReturnsOnCall == nil {
		fake.setBlockReturnsOnCall = make(map[int]struct {
			result1 parsley.Error
		})
	}
	fake.setBlockReturnsOnCall[i] = struct {
		result1 parsley.Error
	}{result1}
}

func (fake *FakeInterpreter) SetParam(arg1 *basil.EvalContext, arg2 basil.Block, arg3 basil.ID, arg4 basil.BlockParamNode) parsley.Error {
	fake.setParamMutex.Lock()
	ret, specificReturn := fake.setParamReturnsOnCall[len(fake.setParamArgsForCall)]
	fake.setParamArgsForCall = append(fake.setParamArgsForCall, struct {
		arg1 *basil.EvalContext
		arg2 basil.Block
		arg3 basil.ID
		arg4 basil.BlockParamNode
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetParam", []interface{}{arg1, arg2, arg3, arg4})
	fake.setParamMutex.Unlock()
	if fake.SetParamStub != nil {
		return fake.SetParamStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setParamReturns
	return fakeReturns.result1
}

func (fake *FakeInterpreter) SetParamCallCount() int {
	fake.setParamMutex.RLock()
	defer fake.setParamMutex.RUnlock()
	return len(fake.setParamArgsForCall)
}

func (fake *FakeInterpreter) SetParamCalls(stub func(*basil.EvalContext, basil.Block, basil.ID, basil.BlockParamNode) parsley.Error) {
	fake.setParamMutex.Lock()
	defer fake.setParamMutex.Unlock()
	fake.SetParamStub = stub
}

func (fake *FakeInterpreter) SetParamArgsForCall(i int) (*basil.EvalContext, basil.Block, basil.ID, basil.BlockParamNode) {
	fake.setParamMutex.RLock()
	defer fake.setParamMutex.RUnlock()
	argsForCall := fake.setParamArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeInterpreter) SetParamReturns(result1 parsley.Error) {
	fake.setParamMutex.Lock()
	defer fake.setParamMutex.Unlock()
	fake.SetParamStub = nil
	fake.setParamReturns = struct {
		result1 parsley.Error
	}{result1}
}

func (fake *FakeInterpreter) SetParamReturnsOnCall(i int, result1 parsley.Error) {
	fake.setParamMutex.Lock()
	defer fake.setParamMutex.Unlock()
	fake.SetParamStub = nil
	if fake.setParamReturnsOnCall == nil {
		fake.setParamReturnsOnCall = make(map[int]struct {
			result1 parsley.Error
		})
	}
	fake.setParamReturnsOnCall[i] = struct {
		result1 parsley.Error
	}{result1}
}

func (fake *FakeInterpreter) ValueParamName() basil.ID {
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

func (fake *FakeInterpreter) ValueParamNameCallCount() int {
	fake.valueParamNameMutex.RLock()
	defer fake.valueParamNameMutex.RUnlock()
	return len(fake.valueParamNameArgsForCall)
}

func (fake *FakeInterpreter) ValueParamNameCalls(stub func() basil.ID) {
	fake.valueParamNameMutex.Lock()
	defer fake.valueParamNameMutex.Unlock()
	fake.ValueParamNameStub = stub
}

func (fake *FakeInterpreter) ValueParamNameReturns(result1 basil.ID) {
	fake.valueParamNameMutex.Lock()
	defer fake.valueParamNameMutex.Unlock()
	fake.ValueParamNameStub = nil
	fake.valueParamNameReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeInterpreter) ValueParamNameReturnsOnCall(i int, result1 basil.ID) {
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

func (fake *FakeInterpreter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
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

func (fake *FakeInterpreter) recordInvocation(key string, args []interface{}) {
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

var _ block.Interpreter = new(FakeInterpreter)
