// Code generated by counterfeiter. DO NOT EDIT.
package functionfakes

import (
	"sync"

	"github.com/opsidian/ocl/function"
	"github.com/opsidian/parsley/parsley"
)

type FakeCallable struct {
	CallFunctionStub        func(ctx interface{}, function parsley.Node, params []parsley.Node) (interface{}, parsley.Error)
	callFunctionMutex       sync.RWMutex
	callFunctionArgsForCall []struct {
		ctx      interface{}
		function parsley.Node
		params   []parsley.Node
	}
	callFunctionReturns struct {
		result1 interface{}
		result2 parsley.Error
	}
	callFunctionReturnsOnCall map[int]struct {
		result1 interface{}
		result2 parsley.Error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCallable) CallFunction(ctx interface{}, function parsley.Node, params []parsley.Node) (interface{}, parsley.Error) {
	var paramsCopy []parsley.Node
	if params != nil {
		paramsCopy = make([]parsley.Node, len(params))
		copy(paramsCopy, params)
	}
	fake.callFunctionMutex.Lock()
	ret, specificReturn := fake.callFunctionReturnsOnCall[len(fake.callFunctionArgsForCall)]
	fake.callFunctionArgsForCall = append(fake.callFunctionArgsForCall, struct {
		ctx      interface{}
		function parsley.Node
		params   []parsley.Node
	}{ctx, function, paramsCopy})
	fake.recordInvocation("CallFunction", []interface{}{ctx, function, paramsCopy})
	fake.callFunctionMutex.Unlock()
	if fake.CallFunctionStub != nil {
		return fake.CallFunctionStub(ctx, function, params)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.callFunctionReturns.result1, fake.callFunctionReturns.result2
}

func (fake *FakeCallable) CallFunctionCallCount() int {
	fake.callFunctionMutex.RLock()
	defer fake.callFunctionMutex.RUnlock()
	return len(fake.callFunctionArgsForCall)
}

func (fake *FakeCallable) CallFunctionArgsForCall(i int) (interface{}, parsley.Node, []parsley.Node) {
	fake.callFunctionMutex.RLock()
	defer fake.callFunctionMutex.RUnlock()
	return fake.callFunctionArgsForCall[i].ctx, fake.callFunctionArgsForCall[i].function, fake.callFunctionArgsForCall[i].params
}

func (fake *FakeCallable) CallFunctionReturns(result1 interface{}, result2 parsley.Error) {
	fake.CallFunctionStub = nil
	fake.callFunctionReturns = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeCallable) CallFunctionReturnsOnCall(i int, result1 interface{}, result2 parsley.Error) {
	fake.CallFunctionStub = nil
	if fake.callFunctionReturnsOnCall == nil {
		fake.callFunctionReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 parsley.Error
		})
	}
	fake.callFunctionReturnsOnCall[i] = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeCallable) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.callFunctionMutex.RLock()
	defer fake.callFunctionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCallable) recordInvocation(key string, args []interface{}) {
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

var _ function.Callable = new(FakeCallable)
