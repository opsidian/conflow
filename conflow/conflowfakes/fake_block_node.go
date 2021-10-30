// Code generated by counterfeiter. DO NOT EDIT.
package conflowfakes

import (
	"sync"

	"github.com/conflowio/conflow/conflow"
	"github.com/conflowio/conflow/conflow/schema"
	"github.com/conflowio/parsley/parsley"
)

type FakeBlockNode struct {
	BlockTypeStub        func() conflow.ID
	blockTypeMutex       sync.RWMutex
	blockTypeArgsForCall []struct {
	}
	blockTypeReturns struct {
		result1 conflow.ID
	}
	blockTypeReturnsOnCall map[int]struct {
		result1 conflow.ID
	}
	ChildrenStub        func() []conflow.Node
	childrenMutex       sync.RWMutex
	childrenArgsForCall []struct {
	}
	childrenReturns struct {
		result1 []conflow.Node
	}
	childrenReturnsOnCall map[int]struct {
		result1 []conflow.Node
	}
	CreateContainerStub        func(*conflow.EvalContext, conflow.RuntimeConfig, conflow.BlockContainer, interface{}, []conflow.WaitGroup, bool) conflow.JobContainer
	createContainerMutex       sync.RWMutex
	createContainerArgsForCall []struct {
		arg1 *conflow.EvalContext
		arg2 conflow.RuntimeConfig
		arg3 conflow.BlockContainer
		arg4 interface{}
		arg5 []conflow.WaitGroup
		arg6 bool
	}
	createContainerReturns struct {
		result1 conflow.JobContainer
	}
	createContainerReturnsOnCall map[int]struct {
		result1 conflow.JobContainer
	}
	DependenciesStub        func() conflow.Dependencies
	dependenciesMutex       sync.RWMutex
	dependenciesArgsForCall []struct {
	}
	dependenciesReturns struct {
		result1 conflow.Dependencies
	}
	dependenciesReturnsOnCall map[int]struct {
		result1 conflow.Dependencies
	}
	DirectivesStub        func() []conflow.BlockNode
	directivesMutex       sync.RWMutex
	directivesArgsForCall []struct {
	}
	directivesReturns struct {
		result1 []conflow.BlockNode
	}
	directivesReturnsOnCall map[int]struct {
		result1 []conflow.BlockNode
	}
	EvalStageStub        func() conflow.EvalStage
	evalStageMutex       sync.RWMutex
	evalStageArgsForCall []struct {
	}
	evalStageReturns struct {
		result1 conflow.EvalStage
	}
	evalStageReturnsOnCall map[int]struct {
		result1 conflow.EvalStage
	}
	GeneratesStub        func() []conflow.ID
	generatesMutex       sync.RWMutex
	generatesArgsForCall []struct {
	}
	generatesReturns struct {
		result1 []conflow.ID
	}
	generatesReturnsOnCall map[int]struct {
		result1 []conflow.ID
	}
	GetPropertySchemaStub        func(conflow.ID) (schema.Schema, bool)
	getPropertySchemaMutex       sync.RWMutex
	getPropertySchemaArgsForCall []struct {
		arg1 conflow.ID
	}
	getPropertySchemaReturns struct {
		result1 schema.Schema
		result2 bool
	}
	getPropertySchemaReturnsOnCall map[int]struct {
		result1 schema.Schema
		result2 bool
	}
	IDStub        func() conflow.ID
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 conflow.ID
	}
	iDReturnsOnCall map[int]struct {
		result1 conflow.ID
	}
	InterpreterStub        func() conflow.BlockInterpreter
	interpreterMutex       sync.RWMutex
	interpreterArgsForCall []struct {
	}
	interpreterReturns struct {
		result1 conflow.BlockInterpreter
	}
	interpreterReturnsOnCall map[int]struct {
		result1 conflow.BlockInterpreter
	}
	ParameterNameStub        func() conflow.ID
	parameterNameMutex       sync.RWMutex
	parameterNameArgsForCall []struct {
	}
	parameterNameReturns struct {
		result1 conflow.ID
	}
	parameterNameReturnsOnCall map[int]struct {
		result1 conflow.ID
	}
	PosStub        func() parsley.Pos
	posMutex       sync.RWMutex
	posArgsForCall []struct {
	}
	posReturns struct {
		result1 parsley.Pos
	}
	posReturnsOnCall map[int]struct {
		result1 parsley.Pos
	}
	ProvidesStub        func() []conflow.ID
	providesMutex       sync.RWMutex
	providesArgsForCall []struct {
	}
	providesReturns struct {
		result1 []conflow.ID
	}
	providesReturnsOnCall map[int]struct {
		result1 []conflow.ID
	}
	ReaderPosStub        func() parsley.Pos
	readerPosMutex       sync.RWMutex
	readerPosArgsForCall []struct {
	}
	readerPosReturns struct {
		result1 parsley.Pos
	}
	readerPosReturnsOnCall map[int]struct {
		result1 parsley.Pos
	}
	SchemaStub        func() interface{}
	schemaMutex       sync.RWMutex
	schemaArgsForCall []struct {
	}
	schemaReturns struct {
		result1 interface{}
	}
	schemaReturnsOnCall map[int]struct {
		result1 interface{}
	}
	SetSchemaStub        func(schema.Schema)
	setSchemaMutex       sync.RWMutex
	setSchemaArgsForCall []struct {
		arg1 schema.Schema
	}
	TokenStub        func() string
	tokenMutex       sync.RWMutex
	tokenArgsForCall []struct {
	}
	tokenReturns struct {
		result1 string
	}
	tokenReturnsOnCall map[int]struct {
		result1 string
	}
	ValueStub        func(interface{}) (interface{}, parsley.Error)
	valueMutex       sync.RWMutex
	valueArgsForCall []struct {
		arg1 interface{}
	}
	valueReturns struct {
		result1 interface{}
		result2 parsley.Error
	}
	valueReturnsOnCall map[int]struct {
		result1 interface{}
		result2 parsley.Error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlockNode) BlockType() conflow.ID {
	fake.blockTypeMutex.Lock()
	ret, specificReturn := fake.blockTypeReturnsOnCall[len(fake.blockTypeArgsForCall)]
	fake.blockTypeArgsForCall = append(fake.blockTypeArgsForCall, struct {
	}{})
	stub := fake.BlockTypeStub
	fakeReturns := fake.blockTypeReturns
	fake.recordInvocation("BlockType", []interface{}{})
	fake.blockTypeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) BlockTypeCallCount() int {
	fake.blockTypeMutex.RLock()
	defer fake.blockTypeMutex.RUnlock()
	return len(fake.blockTypeArgsForCall)
}

func (fake *FakeBlockNode) BlockTypeCalls(stub func() conflow.ID) {
	fake.blockTypeMutex.Lock()
	defer fake.blockTypeMutex.Unlock()
	fake.BlockTypeStub = stub
}

func (fake *FakeBlockNode) BlockTypeReturns(result1 conflow.ID) {
	fake.blockTypeMutex.Lock()
	defer fake.blockTypeMutex.Unlock()
	fake.BlockTypeStub = nil
	fake.blockTypeReturns = struct {
		result1 conflow.ID
	}{result1}
}

func (fake *FakeBlockNode) BlockTypeReturnsOnCall(i int, result1 conflow.ID) {
	fake.blockTypeMutex.Lock()
	defer fake.blockTypeMutex.Unlock()
	fake.BlockTypeStub = nil
	if fake.blockTypeReturnsOnCall == nil {
		fake.blockTypeReturnsOnCall = make(map[int]struct {
			result1 conflow.ID
		})
	}
	fake.blockTypeReturnsOnCall[i] = struct {
		result1 conflow.ID
	}{result1}
}

func (fake *FakeBlockNode) Children() []conflow.Node {
	fake.childrenMutex.Lock()
	ret, specificReturn := fake.childrenReturnsOnCall[len(fake.childrenArgsForCall)]
	fake.childrenArgsForCall = append(fake.childrenArgsForCall, struct {
	}{})
	stub := fake.ChildrenStub
	fakeReturns := fake.childrenReturns
	fake.recordInvocation("Children", []interface{}{})
	fake.childrenMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) ChildrenCallCount() int {
	fake.childrenMutex.RLock()
	defer fake.childrenMutex.RUnlock()
	return len(fake.childrenArgsForCall)
}

func (fake *FakeBlockNode) ChildrenCalls(stub func() []conflow.Node) {
	fake.childrenMutex.Lock()
	defer fake.childrenMutex.Unlock()
	fake.ChildrenStub = stub
}

func (fake *FakeBlockNode) ChildrenReturns(result1 []conflow.Node) {
	fake.childrenMutex.Lock()
	defer fake.childrenMutex.Unlock()
	fake.ChildrenStub = nil
	fake.childrenReturns = struct {
		result1 []conflow.Node
	}{result1}
}

func (fake *FakeBlockNode) ChildrenReturnsOnCall(i int, result1 []conflow.Node) {
	fake.childrenMutex.Lock()
	defer fake.childrenMutex.Unlock()
	fake.ChildrenStub = nil
	if fake.childrenReturnsOnCall == nil {
		fake.childrenReturnsOnCall = make(map[int]struct {
			result1 []conflow.Node
		})
	}
	fake.childrenReturnsOnCall[i] = struct {
		result1 []conflow.Node
	}{result1}
}

func (fake *FakeBlockNode) CreateContainer(arg1 *conflow.EvalContext, arg2 conflow.RuntimeConfig, arg3 conflow.BlockContainer, arg4 interface{}, arg5 []conflow.WaitGroup, arg6 bool) conflow.JobContainer {
	var arg5Copy []conflow.WaitGroup
	if arg5 != nil {
		arg5Copy = make([]conflow.WaitGroup, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.createContainerMutex.Lock()
	ret, specificReturn := fake.createContainerReturnsOnCall[len(fake.createContainerArgsForCall)]
	fake.createContainerArgsForCall = append(fake.createContainerArgsForCall, struct {
		arg1 *conflow.EvalContext
		arg2 conflow.RuntimeConfig
		arg3 conflow.BlockContainer
		arg4 interface{}
		arg5 []conflow.WaitGroup
		arg6 bool
	}{arg1, arg2, arg3, arg4, arg5Copy, arg6})
	stub := fake.CreateContainerStub
	fakeReturns := fake.createContainerReturns
	fake.recordInvocation("CreateContainer", []interface{}{arg1, arg2, arg3, arg4, arg5Copy, arg6})
	fake.createContainerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) CreateContainerCallCount() int {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return len(fake.createContainerArgsForCall)
}

func (fake *FakeBlockNode) CreateContainerCalls(stub func(*conflow.EvalContext, conflow.RuntimeConfig, conflow.BlockContainer, interface{}, []conflow.WaitGroup, bool) conflow.JobContainer) {
	fake.createContainerMutex.Lock()
	defer fake.createContainerMutex.Unlock()
	fake.CreateContainerStub = stub
}

func (fake *FakeBlockNode) CreateContainerArgsForCall(i int) (*conflow.EvalContext, conflow.RuntimeConfig, conflow.BlockContainer, interface{}, []conflow.WaitGroup, bool) {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	argsForCall := fake.createContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeBlockNode) CreateContainerReturns(result1 conflow.JobContainer) {
	fake.createContainerMutex.Lock()
	defer fake.createContainerMutex.Unlock()
	fake.CreateContainerStub = nil
	fake.createContainerReturns = struct {
		result1 conflow.JobContainer
	}{result1}
}

func (fake *FakeBlockNode) CreateContainerReturnsOnCall(i int, result1 conflow.JobContainer) {
	fake.createContainerMutex.Lock()
	defer fake.createContainerMutex.Unlock()
	fake.CreateContainerStub = nil
	if fake.createContainerReturnsOnCall == nil {
		fake.createContainerReturnsOnCall = make(map[int]struct {
			result1 conflow.JobContainer
		})
	}
	fake.createContainerReturnsOnCall[i] = struct {
		result1 conflow.JobContainer
	}{result1}
}

func (fake *FakeBlockNode) Dependencies() conflow.Dependencies {
	fake.dependenciesMutex.Lock()
	ret, specificReturn := fake.dependenciesReturnsOnCall[len(fake.dependenciesArgsForCall)]
	fake.dependenciesArgsForCall = append(fake.dependenciesArgsForCall, struct {
	}{})
	stub := fake.DependenciesStub
	fakeReturns := fake.dependenciesReturns
	fake.recordInvocation("Dependencies", []interface{}{})
	fake.dependenciesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) DependenciesCallCount() int {
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	return len(fake.dependenciesArgsForCall)
}

func (fake *FakeBlockNode) DependenciesCalls(stub func() conflow.Dependencies) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = stub
}

func (fake *FakeBlockNode) DependenciesReturns(result1 conflow.Dependencies) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = nil
	fake.dependenciesReturns = struct {
		result1 conflow.Dependencies
	}{result1}
}

func (fake *FakeBlockNode) DependenciesReturnsOnCall(i int, result1 conflow.Dependencies) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = nil
	if fake.dependenciesReturnsOnCall == nil {
		fake.dependenciesReturnsOnCall = make(map[int]struct {
			result1 conflow.Dependencies
		})
	}
	fake.dependenciesReturnsOnCall[i] = struct {
		result1 conflow.Dependencies
	}{result1}
}

func (fake *FakeBlockNode) Directives() []conflow.BlockNode {
	fake.directivesMutex.Lock()
	ret, specificReturn := fake.directivesReturnsOnCall[len(fake.directivesArgsForCall)]
	fake.directivesArgsForCall = append(fake.directivesArgsForCall, struct {
	}{})
	stub := fake.DirectivesStub
	fakeReturns := fake.directivesReturns
	fake.recordInvocation("Directives", []interface{}{})
	fake.directivesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) DirectivesCallCount() int {
	fake.directivesMutex.RLock()
	defer fake.directivesMutex.RUnlock()
	return len(fake.directivesArgsForCall)
}

func (fake *FakeBlockNode) DirectivesCalls(stub func() []conflow.BlockNode) {
	fake.directivesMutex.Lock()
	defer fake.directivesMutex.Unlock()
	fake.DirectivesStub = stub
}

func (fake *FakeBlockNode) DirectivesReturns(result1 []conflow.BlockNode) {
	fake.directivesMutex.Lock()
	defer fake.directivesMutex.Unlock()
	fake.DirectivesStub = nil
	fake.directivesReturns = struct {
		result1 []conflow.BlockNode
	}{result1}
}

func (fake *FakeBlockNode) DirectivesReturnsOnCall(i int, result1 []conflow.BlockNode) {
	fake.directivesMutex.Lock()
	defer fake.directivesMutex.Unlock()
	fake.DirectivesStub = nil
	if fake.directivesReturnsOnCall == nil {
		fake.directivesReturnsOnCall = make(map[int]struct {
			result1 []conflow.BlockNode
		})
	}
	fake.directivesReturnsOnCall[i] = struct {
		result1 []conflow.BlockNode
	}{result1}
}

func (fake *FakeBlockNode) EvalStage() conflow.EvalStage {
	fake.evalStageMutex.Lock()
	ret, specificReturn := fake.evalStageReturnsOnCall[len(fake.evalStageArgsForCall)]
	fake.evalStageArgsForCall = append(fake.evalStageArgsForCall, struct {
	}{})
	stub := fake.EvalStageStub
	fakeReturns := fake.evalStageReturns
	fake.recordInvocation("EvalStage", []interface{}{})
	fake.evalStageMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) EvalStageCallCount() int {
	fake.evalStageMutex.RLock()
	defer fake.evalStageMutex.RUnlock()
	return len(fake.evalStageArgsForCall)
}

func (fake *FakeBlockNode) EvalStageCalls(stub func() conflow.EvalStage) {
	fake.evalStageMutex.Lock()
	defer fake.evalStageMutex.Unlock()
	fake.EvalStageStub = stub
}

func (fake *FakeBlockNode) EvalStageReturns(result1 conflow.EvalStage) {
	fake.evalStageMutex.Lock()
	defer fake.evalStageMutex.Unlock()
	fake.EvalStageStub = nil
	fake.evalStageReturns = struct {
		result1 conflow.EvalStage
	}{result1}
}

func (fake *FakeBlockNode) EvalStageReturnsOnCall(i int, result1 conflow.EvalStage) {
	fake.evalStageMutex.Lock()
	defer fake.evalStageMutex.Unlock()
	fake.EvalStageStub = nil
	if fake.evalStageReturnsOnCall == nil {
		fake.evalStageReturnsOnCall = make(map[int]struct {
			result1 conflow.EvalStage
		})
	}
	fake.evalStageReturnsOnCall[i] = struct {
		result1 conflow.EvalStage
	}{result1}
}

func (fake *FakeBlockNode) Generates() []conflow.ID {
	fake.generatesMutex.Lock()
	ret, specificReturn := fake.generatesReturnsOnCall[len(fake.generatesArgsForCall)]
	fake.generatesArgsForCall = append(fake.generatesArgsForCall, struct {
	}{})
	stub := fake.GeneratesStub
	fakeReturns := fake.generatesReturns
	fake.recordInvocation("Generates", []interface{}{})
	fake.generatesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) GeneratesCallCount() int {
	fake.generatesMutex.RLock()
	defer fake.generatesMutex.RUnlock()
	return len(fake.generatesArgsForCall)
}

func (fake *FakeBlockNode) GeneratesCalls(stub func() []conflow.ID) {
	fake.generatesMutex.Lock()
	defer fake.generatesMutex.Unlock()
	fake.GeneratesStub = stub
}

func (fake *FakeBlockNode) GeneratesReturns(result1 []conflow.ID) {
	fake.generatesMutex.Lock()
	defer fake.generatesMutex.Unlock()
	fake.GeneratesStub = nil
	fake.generatesReturns = struct {
		result1 []conflow.ID
	}{result1}
}

func (fake *FakeBlockNode) GeneratesReturnsOnCall(i int, result1 []conflow.ID) {
	fake.generatesMutex.Lock()
	defer fake.generatesMutex.Unlock()
	fake.GeneratesStub = nil
	if fake.generatesReturnsOnCall == nil {
		fake.generatesReturnsOnCall = make(map[int]struct {
			result1 []conflow.ID
		})
	}
	fake.generatesReturnsOnCall[i] = struct {
		result1 []conflow.ID
	}{result1}
}

func (fake *FakeBlockNode) GetPropertySchema(arg1 conflow.ID) (schema.Schema, bool) {
	fake.getPropertySchemaMutex.Lock()
	ret, specificReturn := fake.getPropertySchemaReturnsOnCall[len(fake.getPropertySchemaArgsForCall)]
	fake.getPropertySchemaArgsForCall = append(fake.getPropertySchemaArgsForCall, struct {
		arg1 conflow.ID
	}{arg1})
	stub := fake.GetPropertySchemaStub
	fakeReturns := fake.getPropertySchemaReturns
	fake.recordInvocation("GetPropertySchema", []interface{}{arg1})
	fake.getPropertySchemaMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlockNode) GetPropertySchemaCallCount() int {
	fake.getPropertySchemaMutex.RLock()
	defer fake.getPropertySchemaMutex.RUnlock()
	return len(fake.getPropertySchemaArgsForCall)
}

func (fake *FakeBlockNode) GetPropertySchemaCalls(stub func(conflow.ID) (schema.Schema, bool)) {
	fake.getPropertySchemaMutex.Lock()
	defer fake.getPropertySchemaMutex.Unlock()
	fake.GetPropertySchemaStub = stub
}

func (fake *FakeBlockNode) GetPropertySchemaArgsForCall(i int) conflow.ID {
	fake.getPropertySchemaMutex.RLock()
	defer fake.getPropertySchemaMutex.RUnlock()
	argsForCall := fake.getPropertySchemaArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockNode) GetPropertySchemaReturns(result1 schema.Schema, result2 bool) {
	fake.getPropertySchemaMutex.Lock()
	defer fake.getPropertySchemaMutex.Unlock()
	fake.GetPropertySchemaStub = nil
	fake.getPropertySchemaReturns = struct {
		result1 schema.Schema
		result2 bool
	}{result1, result2}
}

func (fake *FakeBlockNode) GetPropertySchemaReturnsOnCall(i int, result1 schema.Schema, result2 bool) {
	fake.getPropertySchemaMutex.Lock()
	defer fake.getPropertySchemaMutex.Unlock()
	fake.GetPropertySchemaStub = nil
	if fake.getPropertySchemaReturnsOnCall == nil {
		fake.getPropertySchemaReturnsOnCall = make(map[int]struct {
			result1 schema.Schema
			result2 bool
		})
	}
	fake.getPropertySchemaReturnsOnCall[i] = struct {
		result1 schema.Schema
		result2 bool
	}{result1, result2}
}

func (fake *FakeBlockNode) ID() conflow.ID {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	stub := fake.IDStub
	fakeReturns := fake.iDReturns
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeBlockNode) IDCalls(stub func() conflow.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeBlockNode) IDReturns(result1 conflow.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 conflow.ID
	}{result1}
}

func (fake *FakeBlockNode) IDReturnsOnCall(i int, result1 conflow.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 conflow.ID
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 conflow.ID
	}{result1}
}

func (fake *FakeBlockNode) Interpreter() conflow.BlockInterpreter {
	fake.interpreterMutex.Lock()
	ret, specificReturn := fake.interpreterReturnsOnCall[len(fake.interpreterArgsForCall)]
	fake.interpreterArgsForCall = append(fake.interpreterArgsForCall, struct {
	}{})
	stub := fake.InterpreterStub
	fakeReturns := fake.interpreterReturns
	fake.recordInvocation("Interpreter", []interface{}{})
	fake.interpreterMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) InterpreterCallCount() int {
	fake.interpreterMutex.RLock()
	defer fake.interpreterMutex.RUnlock()
	return len(fake.interpreterArgsForCall)
}

func (fake *FakeBlockNode) InterpreterCalls(stub func() conflow.BlockInterpreter) {
	fake.interpreterMutex.Lock()
	defer fake.interpreterMutex.Unlock()
	fake.InterpreterStub = stub
}

func (fake *FakeBlockNode) InterpreterReturns(result1 conflow.BlockInterpreter) {
	fake.interpreterMutex.Lock()
	defer fake.interpreterMutex.Unlock()
	fake.InterpreterStub = nil
	fake.interpreterReturns = struct {
		result1 conflow.BlockInterpreter
	}{result1}
}

func (fake *FakeBlockNode) InterpreterReturnsOnCall(i int, result1 conflow.BlockInterpreter) {
	fake.interpreterMutex.Lock()
	defer fake.interpreterMutex.Unlock()
	fake.InterpreterStub = nil
	if fake.interpreterReturnsOnCall == nil {
		fake.interpreterReturnsOnCall = make(map[int]struct {
			result1 conflow.BlockInterpreter
		})
	}
	fake.interpreterReturnsOnCall[i] = struct {
		result1 conflow.BlockInterpreter
	}{result1}
}

func (fake *FakeBlockNode) ParameterName() conflow.ID {
	fake.parameterNameMutex.Lock()
	ret, specificReturn := fake.parameterNameReturnsOnCall[len(fake.parameterNameArgsForCall)]
	fake.parameterNameArgsForCall = append(fake.parameterNameArgsForCall, struct {
	}{})
	stub := fake.ParameterNameStub
	fakeReturns := fake.parameterNameReturns
	fake.recordInvocation("ParameterName", []interface{}{})
	fake.parameterNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) ParameterNameCallCount() int {
	fake.parameterNameMutex.RLock()
	defer fake.parameterNameMutex.RUnlock()
	return len(fake.parameterNameArgsForCall)
}

func (fake *FakeBlockNode) ParameterNameCalls(stub func() conflow.ID) {
	fake.parameterNameMutex.Lock()
	defer fake.parameterNameMutex.Unlock()
	fake.ParameterNameStub = stub
}

func (fake *FakeBlockNode) ParameterNameReturns(result1 conflow.ID) {
	fake.parameterNameMutex.Lock()
	defer fake.parameterNameMutex.Unlock()
	fake.ParameterNameStub = nil
	fake.parameterNameReturns = struct {
		result1 conflow.ID
	}{result1}
}

func (fake *FakeBlockNode) ParameterNameReturnsOnCall(i int, result1 conflow.ID) {
	fake.parameterNameMutex.Lock()
	defer fake.parameterNameMutex.Unlock()
	fake.ParameterNameStub = nil
	if fake.parameterNameReturnsOnCall == nil {
		fake.parameterNameReturnsOnCall = make(map[int]struct {
			result1 conflow.ID
		})
	}
	fake.parameterNameReturnsOnCall[i] = struct {
		result1 conflow.ID
	}{result1}
}

func (fake *FakeBlockNode) Pos() parsley.Pos {
	fake.posMutex.Lock()
	ret, specificReturn := fake.posReturnsOnCall[len(fake.posArgsForCall)]
	fake.posArgsForCall = append(fake.posArgsForCall, struct {
	}{})
	stub := fake.PosStub
	fakeReturns := fake.posReturns
	fake.recordInvocation("Pos", []interface{}{})
	fake.posMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) PosCallCount() int {
	fake.posMutex.RLock()
	defer fake.posMutex.RUnlock()
	return len(fake.posArgsForCall)
}

func (fake *FakeBlockNode) PosCalls(stub func() parsley.Pos) {
	fake.posMutex.Lock()
	defer fake.posMutex.Unlock()
	fake.PosStub = stub
}

func (fake *FakeBlockNode) PosReturns(result1 parsley.Pos) {
	fake.posMutex.Lock()
	defer fake.posMutex.Unlock()
	fake.PosStub = nil
	fake.posReturns = struct {
		result1 parsley.Pos
	}{result1}
}

func (fake *FakeBlockNode) PosReturnsOnCall(i int, result1 parsley.Pos) {
	fake.posMutex.Lock()
	defer fake.posMutex.Unlock()
	fake.PosStub = nil
	if fake.posReturnsOnCall == nil {
		fake.posReturnsOnCall = make(map[int]struct {
			result1 parsley.Pos
		})
	}
	fake.posReturnsOnCall[i] = struct {
		result1 parsley.Pos
	}{result1}
}

func (fake *FakeBlockNode) Provides() []conflow.ID {
	fake.providesMutex.Lock()
	ret, specificReturn := fake.providesReturnsOnCall[len(fake.providesArgsForCall)]
	fake.providesArgsForCall = append(fake.providesArgsForCall, struct {
	}{})
	stub := fake.ProvidesStub
	fakeReturns := fake.providesReturns
	fake.recordInvocation("Provides", []interface{}{})
	fake.providesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) ProvidesCallCount() int {
	fake.providesMutex.RLock()
	defer fake.providesMutex.RUnlock()
	return len(fake.providesArgsForCall)
}

func (fake *FakeBlockNode) ProvidesCalls(stub func() []conflow.ID) {
	fake.providesMutex.Lock()
	defer fake.providesMutex.Unlock()
	fake.ProvidesStub = stub
}

func (fake *FakeBlockNode) ProvidesReturns(result1 []conflow.ID) {
	fake.providesMutex.Lock()
	defer fake.providesMutex.Unlock()
	fake.ProvidesStub = nil
	fake.providesReturns = struct {
		result1 []conflow.ID
	}{result1}
}

func (fake *FakeBlockNode) ProvidesReturnsOnCall(i int, result1 []conflow.ID) {
	fake.providesMutex.Lock()
	defer fake.providesMutex.Unlock()
	fake.ProvidesStub = nil
	if fake.providesReturnsOnCall == nil {
		fake.providesReturnsOnCall = make(map[int]struct {
			result1 []conflow.ID
		})
	}
	fake.providesReturnsOnCall[i] = struct {
		result1 []conflow.ID
	}{result1}
}

func (fake *FakeBlockNode) ReaderPos() parsley.Pos {
	fake.readerPosMutex.Lock()
	ret, specificReturn := fake.readerPosReturnsOnCall[len(fake.readerPosArgsForCall)]
	fake.readerPosArgsForCall = append(fake.readerPosArgsForCall, struct {
	}{})
	stub := fake.ReaderPosStub
	fakeReturns := fake.readerPosReturns
	fake.recordInvocation("ReaderPos", []interface{}{})
	fake.readerPosMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) ReaderPosCallCount() int {
	fake.readerPosMutex.RLock()
	defer fake.readerPosMutex.RUnlock()
	return len(fake.readerPosArgsForCall)
}

func (fake *FakeBlockNode) ReaderPosCalls(stub func() parsley.Pos) {
	fake.readerPosMutex.Lock()
	defer fake.readerPosMutex.Unlock()
	fake.ReaderPosStub = stub
}

func (fake *FakeBlockNode) ReaderPosReturns(result1 parsley.Pos) {
	fake.readerPosMutex.Lock()
	defer fake.readerPosMutex.Unlock()
	fake.ReaderPosStub = nil
	fake.readerPosReturns = struct {
		result1 parsley.Pos
	}{result1}
}

func (fake *FakeBlockNode) ReaderPosReturnsOnCall(i int, result1 parsley.Pos) {
	fake.readerPosMutex.Lock()
	defer fake.readerPosMutex.Unlock()
	fake.ReaderPosStub = nil
	if fake.readerPosReturnsOnCall == nil {
		fake.readerPosReturnsOnCall = make(map[int]struct {
			result1 parsley.Pos
		})
	}
	fake.readerPosReturnsOnCall[i] = struct {
		result1 parsley.Pos
	}{result1}
}

func (fake *FakeBlockNode) Schema() interface{} {
	fake.schemaMutex.Lock()
	ret, specificReturn := fake.schemaReturnsOnCall[len(fake.schemaArgsForCall)]
	fake.schemaArgsForCall = append(fake.schemaArgsForCall, struct {
	}{})
	stub := fake.SchemaStub
	fakeReturns := fake.schemaReturns
	fake.recordInvocation("Schema", []interface{}{})
	fake.schemaMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) SchemaCallCount() int {
	fake.schemaMutex.RLock()
	defer fake.schemaMutex.RUnlock()
	return len(fake.schemaArgsForCall)
}

func (fake *FakeBlockNode) SchemaCalls(stub func() interface{}) {
	fake.schemaMutex.Lock()
	defer fake.schemaMutex.Unlock()
	fake.SchemaStub = stub
}

func (fake *FakeBlockNode) SchemaReturns(result1 interface{}) {
	fake.schemaMutex.Lock()
	defer fake.schemaMutex.Unlock()
	fake.SchemaStub = nil
	fake.schemaReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeBlockNode) SchemaReturnsOnCall(i int, result1 interface{}) {
	fake.schemaMutex.Lock()
	defer fake.schemaMutex.Unlock()
	fake.SchemaStub = nil
	if fake.schemaReturnsOnCall == nil {
		fake.schemaReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.schemaReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeBlockNode) SetSchema(arg1 schema.Schema) {
	fake.setSchemaMutex.Lock()
	fake.setSchemaArgsForCall = append(fake.setSchemaArgsForCall, struct {
		arg1 schema.Schema
	}{arg1})
	stub := fake.SetSchemaStub
	fake.recordInvocation("SetSchema", []interface{}{arg1})
	fake.setSchemaMutex.Unlock()
	if stub != nil {
		fake.SetSchemaStub(arg1)
	}
}

func (fake *FakeBlockNode) SetSchemaCallCount() int {
	fake.setSchemaMutex.RLock()
	defer fake.setSchemaMutex.RUnlock()
	return len(fake.setSchemaArgsForCall)
}

func (fake *FakeBlockNode) SetSchemaCalls(stub func(schema.Schema)) {
	fake.setSchemaMutex.Lock()
	defer fake.setSchemaMutex.Unlock()
	fake.SetSchemaStub = stub
}

func (fake *FakeBlockNode) SetSchemaArgsForCall(i int) schema.Schema {
	fake.setSchemaMutex.RLock()
	defer fake.setSchemaMutex.RUnlock()
	argsForCall := fake.setSchemaArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockNode) Token() string {
	fake.tokenMutex.Lock()
	ret, specificReturn := fake.tokenReturnsOnCall[len(fake.tokenArgsForCall)]
	fake.tokenArgsForCall = append(fake.tokenArgsForCall, struct {
	}{})
	stub := fake.TokenStub
	fakeReturns := fake.tokenReturns
	fake.recordInvocation("Token", []interface{}{})
	fake.tokenMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBlockNode) TokenCallCount() int {
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	return len(fake.tokenArgsForCall)
}

func (fake *FakeBlockNode) TokenCalls(stub func() string) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = stub
}

func (fake *FakeBlockNode) TokenReturns(result1 string) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = nil
	fake.tokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBlockNode) TokenReturnsOnCall(i int, result1 string) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = nil
	if fake.tokenReturnsOnCall == nil {
		fake.tokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.tokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBlockNode) Value(arg1 interface{}) (interface{}, parsley.Error) {
	fake.valueMutex.Lock()
	ret, specificReturn := fake.valueReturnsOnCall[len(fake.valueArgsForCall)]
	fake.valueArgsForCall = append(fake.valueArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	stub := fake.ValueStub
	fakeReturns := fake.valueReturns
	fake.recordInvocation("Value", []interface{}{arg1})
	fake.valueMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlockNode) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeBlockNode) ValueCalls(stub func(interface{}) (interface{}, parsley.Error)) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = stub
}

func (fake *FakeBlockNode) ValueArgsForCall(i int) interface{} {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	argsForCall := fake.valueArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockNode) ValueReturns(result1 interface{}, result2 parsley.Error) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeBlockNode) ValueReturnsOnCall(i int, result1 interface{}, result2 parsley.Error) {
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

func (fake *FakeBlockNode) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blockTypeMutex.RLock()
	defer fake.blockTypeMutex.RUnlock()
	fake.childrenMutex.RLock()
	defer fake.childrenMutex.RUnlock()
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	fake.directivesMutex.RLock()
	defer fake.directivesMutex.RUnlock()
	fake.evalStageMutex.RLock()
	defer fake.evalStageMutex.RUnlock()
	fake.generatesMutex.RLock()
	defer fake.generatesMutex.RUnlock()
	fake.getPropertySchemaMutex.RLock()
	defer fake.getPropertySchemaMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.interpreterMutex.RLock()
	defer fake.interpreterMutex.RUnlock()
	fake.parameterNameMutex.RLock()
	defer fake.parameterNameMutex.RUnlock()
	fake.posMutex.RLock()
	defer fake.posMutex.RUnlock()
	fake.providesMutex.RLock()
	defer fake.providesMutex.RUnlock()
	fake.readerPosMutex.RLock()
	defer fake.readerPosMutex.RUnlock()
	fake.schemaMutex.RLock()
	defer fake.schemaMutex.RUnlock()
	fake.setSchemaMutex.RLock()
	defer fake.setSchemaMutex.RUnlock()
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlockNode) recordInvocation(key string, args []interface{}) {
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

var _ conflow.BlockNode = new(FakeBlockNode)
